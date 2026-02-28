package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/user_service/internal/repository"
	"github.com/user_service/internal/utils/random"
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
	// 0. Hash email
	// 1. Check email exists in DB
	// 2. new OTP
	otp := random.GenerateOPT6Digit()
	fmt.Printf("OTP: %d\n", otp)
	// 3. Save OTP in Redis
	// 4. Send OTP to email
	err := s.authRepo.CreateNewUser(ctx, username, password, email)

	return err
}

// Login authenticates a user and returns access/refresh tokens
func (s *AuthService) LoginService(ctx context.Context, username string, password string) (string, string, error) {
	user, err := s.authRepo.FindByUsername(ctx, username)

	if err != nil {
		return "", "", errors.New("Invalid credentials, USER NOT FOUND")
	}

	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return "", "", errors.New("Account is locked")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("Invalid credentials")
	}
	sessionID := uuid.NewString()
	s.authRepo.UpdateLastLogin(ctx, user.ID)
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, sessionID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, sessionID)
	if err != nil {
		return "", "", err
	}

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

	isSessionBlocked, err := s.blacklist.IsSessionBlacklisted(ctx, claims.SessionID)
	if err != nil || isSessionBlocked {
		return "", "", errors.New("session has been revoked")
	}

	isJTIBlocked, err := s.blacklist.IsJTIBlacklisted(ctx, claims.ID)
	if err != nil || isJTIBlocked {
		return "", "", errors.New("token has been reused")
	}
	ttl := time.Until(claims.ExpiresAt.Time)
	if err := s.blacklist.BlacklistJTI(ctx, claims.ID, ttl); err != nil {
		return "", "", err
	}
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
