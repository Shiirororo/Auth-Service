package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/auth_service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo   repository.AuthRepository
	blacklist  TokenBlacklist
	jwtService *jwtService
}

func NewAuthService(authRepo repository.AuthRepository, blacklist TokenBlacklist) *AuthService {
	return &AuthService{
		authRepo:   authRepo,
		blacklist:  blacklist,
		jwtService: NewJWTService(os.Getenv("JWT_SECRET")),
	}
}

// Login authenticates a user and returns access/refresh tokens
func (s *AuthService) LoginService(ctx context.Context, username string, password string) (string, string, error) {
	// 1. Find user
	user, err := s.authRepo.FindByUsername(ctx, username)
	if err != nil {
		// Log the actual error internally if needed, but return generic error to user
		return "", "", errors.New("Invalid credentials")
	}

	// 2. Check lock
	// Assuming LockedUntil is *time.Time
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return "", "", errors.New("Account is locked")
	}

	// 3. Compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", "", errors.New("Invalid credentials")
	}

	// 4. Generate Token
	// Assuming logic to get role from user model exists, or default to "user"
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID.String()) // Convert UUID to string
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return "", "", err
	}

	// 5. Update last_login (async or sync)
	// go s.authRepo.UpdateLastLogin(ctx, user.ID)

	// 5. Store Refresh Token
	err = s.blacklist.SetRefreshToken(ctx, user.ID.String(), refreshToken, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) LogoutService(ctx context.Context, userID string, JIT string, TTL time.Duration) error {
	// 1. Revoke Refresh Token
	if err := s.blacklist.SetRefreshToken(ctx, userID, JIT, TTL); err != nil {
		return err
	}

	return s.blacklist.SetRefreshToken(ctx, userID, JIT, TTL)
}

func (s *AuthService) RefreshTokenService(ctx context.Context, refreshToken string) (string, string, error) {
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
	err = s.blacklist.SetRefreshToken(ctx, claims.UserID, newRefreshToken, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

// JWT service management:
// - Generate access token
// - Generate refresh token
// - Parse access token
// - Parse refresh token
// - Validate access token
// - Validate refresh token
// - Blacklist access token
// - Blacklist refresh token
// - Get access token
// - Get refresh token
// - Delete access token
// - Delete refresh token
