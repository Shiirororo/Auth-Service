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
	"golang.org/x/sync/errgroup"
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
	updateLastLoginBestEffort(userID string)
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
func (s *AuthService) LoginService(ctx context.Context, username string, password string) (string, string, error) {
	user, err := s.authRepo.FindByUsername(ctx, username)

	if err != nil {
		return "", "", errors.New("Invalid credentials, USER NOT FOUND")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", "", errors.New("Invalid credentials, PASSWORD NOT MATCH")
	}
	g, ctx := errgroup.WithContext(ctx)

	var accessToken, refreshToken string
	sessionID := uuid.NewString()
	g.Go(func() error {
		var err error
		accessToken, err = s.jwtService.GenerateAccessToken(user.ID, sessionID)
		return err
	})

	g.Go(func() error {
		var err error
		refreshToken, err = s.jwtService.GenerateRefreshToken(user.ID, sessionID)
		return err
	})

	if err := g.Wait(); err != nil {
		return "", "", err
	}
	go s.updateLastLoginBestEffort(user.ID)
	return accessToken, refreshToken, nil
}
func (s *AuthService) updateLastLoginBestEffort(userID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_ = s.authRepo.UpdateLastLogin(ctx, userID)
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
	g, ctx := errgroup.WithContext(ctx)
	var newAccessToken, newRefreshToken string
	g.Go(func() error {
		var err error
		newAccessToken, err = s.jwtService.GenerateAccessToken(claims.UserID, claims.SessionID)
		return err
	})
	g.Go(func() error {
		var err error
		newRefreshToken, err = s.jwtService.GenerateRefreshToken(claims.UserID, claims.SessionID)
		return err
	})

	if err := g.Wait(); err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}
