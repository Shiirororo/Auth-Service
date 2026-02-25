package service

import (
	"context"
	"errors"
	"time"

	"github.com/auth_service/internal/repository"
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
	LogoutService(ctx context.Context, userID string, JIT string, TTL time.Time) error
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

	// 4. Generate Token
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID) // Convert UUID to string
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	// 5. Update last_login (async or sync)
	s.authRepo.UpdateLastLogin(ctx, user.ID)

	return accessToken, refreshToken, nil
}

func (s *AuthService) LogoutService(ctx context.Context, userID string, JIT string, TTL time.Time) error {
	// 1. Revoke Refresh Token
	if err := s.blacklist.RevokeRefreshToken(ctx, userID, JIT, TTL); err != nil {
		return err
	}

	return s.blacklist.RevokeRefreshToken(ctx, userID, JIT, TTL)
}

func (s *AuthService) RefreshService(ctx context.Context, refreshToken string) (string, string, error) {
	// 1. Verify Refresh Token
	claims, err := s.jwtService.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	storedToken := "sadasdasda"

	if storedToken != refreshToken {
		// Token reuse detected!
		return "", "", errors.New("invalid refresh token (reuse detected)")
	}

	// 3. Generate New Tokens
	// TODO: Fetch user to get latest role. For now assuming "user" role.

	newAccessToken, err := s.jwtService.GenerateAccessToken(claims.UserID)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.jwtService.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return "", "", err
	}

	// 4. Update Redis (Rotate)
	err = s.blacklist.RevokeRefreshToken(ctx, claims.UserID, newRefreshToken, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
