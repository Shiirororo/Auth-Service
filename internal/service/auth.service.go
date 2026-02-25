package service

import (
	"context"
	"errors"
	"time"

	"github.com/auth_service/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo   repository.AuthRepository
	blacklist  TokenBlacklist
	jwtService JWTService
}

func NewAuthService(authRepo repository.AuthRepository, blacklist TokenBlacklist, jwtService JWTService) AuthServiceInterface {
	return &AuthService{
		authRepo:   authRepo,
		blacklist:  blacklist,
		jwtService: jwtService,
	}
}

type AuthServiceInterface interface {
	RegisterService(ctx context.Context, username string, password string, email string) error
	LoginService(ctx context.Context, username string, password string) (string, string, error)
	LogoutService(ctx context.Context, sessionID string, ttl time.Duration) error
	RefreshService(ctx context.Context, refreshToken string) (string, string, error)
}

// Should I add interface here?
func (s *AuthService) RegisterService(ctx context.Context, username string, password string, email string) error {
	err := s.authRepo.CreateNewUser(ctx, username, password, email)

	return err
}

// Login authenticates a user and returns access/refresh tokens
func (s *AuthService) LoginService(ctx context.Context, username string, password string) (string, string, error) {
	// 1. Find user
	user, err := s.authRepo.FindByUsername(ctx, username)

	if err != nil {
		return "", "", errors.New("Invalid credentials, USER NOT FOUND")
	}

	// 2. Check lock
	// Assuming LockedUntil is *time.Time
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return "", "", errors.New("Account is locked")
	}

	// 3. Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("Invalid credentials")
	}

	// 4. Generate SessionID
	sessionID := uuid.NewString()

	// 5. Generate Tokens
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, sessionID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, sessionID)
	if err != nil {
		return "", "", err
	}

	// 6. Update last_login (async or sync)
	s.authRepo.UpdateLastLogin(ctx, user.ID)

	return accessToken, refreshToken, nil
}

func (s *AuthService) LogoutService(ctx context.Context, sessionID string, ttl time.Duration) error {
	return s.blacklist.BlacklistSession(ctx, sessionID, ttl)
}

func (s *AuthService) RefreshService(ctx context.Context, refreshToken string) (string, string, error) {
	// 1. Verify Refresh Token
	claims, err := s.jwtService.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	// 2. Check Blacklist
	// Check Session
	isSessionBlocked, err := s.blacklist.IsSessionBlacklisted(ctx, claims.SessionID)
	if err != nil || isSessionBlocked {
		return "", "", errors.New("session has been revoked")
	}

	// Check JTI (Reuse detection)
	isJTIBlocked, err := s.blacklist.IsJTIBlacklisted(ctx, claims.ID)
	if err != nil || isJTIBlocked {
		return "", "", errors.New("token has been reused")
	}

	// 3. Blacklist old RT's JTI
	ttl := time.Until(claims.ExpiresAt.Time)
	if err := s.blacklist.BlacklistJTI(ctx, claims.ID, ttl); err != nil {
		return "", "", err
	}

	// 4. Generate New Tokens with SAME SessionID
	newAccessToken, err := s.jwtService.GenerateAccessToken(claims.UserID, claims.SessionID)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.jwtService.GenerateRefreshToken(claims.UserID, claims.SessionID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
