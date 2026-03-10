package service

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/auth/domain/vo"
	"github.com/user_service/internal/commons"
	"github.com/user_service/internal/event"
	"github.com/user_service/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type AuthService struct {
	authRepo   repository.AuthRepository
	otpRepo    repository.OTPRepository
	blacklist  commons.TokenBlacklist
	jwtService token.TokenMaker
	dispatcher *event.Dispatcher
}

func NewAuthService(authRepo repository.AuthRepository, otpRepo repository.OTPRepository, blacklist commons.TokenBlacklist, jwtService token.TokenMaker, dispatcher *event.Dispatcher) AuthServiceInterface {
	return &AuthService{
		authRepo:   authRepo,
		otpRepo:    otpRepo,
		blacklist:  blacklist,
		jwtService: jwtService,
		dispatcher: dispatcher,
	}
}

type AuthServiceInterface interface {
	RegisterService(ctx context.Context, username string, password string, email string) error
	LoginServiceWithUsername(ctx context.Context, username string, password string) (string, string, string, error)
	LoginServiceWithEmail(ctx context.Context, email string, password string) (string, string, string, error)
	LogoutService(ctx context.Context, sessionID string, ttl time.Duration) error
	RefreshService(ctx context.Context, refreshToken string) (string, string, error)
}

func (s *AuthService) RegisterService(ctx context.Context, username string, password string, email string) error {
	// 1. Create Domain Value Objects to validate integrity early
	// Email domain validation is handled by the presentation layer (gin binding)

	//Check duplicate email and username
	_, err := s.authRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return errors.New("Email already exists")
	}

	passVo, err := vo.NewPassword(password)
	if err != nil {
		return err
	}

	// ASSUME OTP AND VERIFICATION COMPLETED HERE

	s.dispatcher.Dispatch(ctx, event.Event{
		Type: event.RegisterSuccessEvent,
		Payload: event.RegisterSuccessPayload{
			Username: username,
			Email:    email,
			Password: passVo.String(),
		},
	})
	return nil
}
func (s *AuthService) LoginServiceWithEmail(ctx context.Context, email string, password string) (string, string, string, error) {

	user, err := s.authRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return "", "", "", errors.New("Invalid credentials, USER NOT FOUND")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", "", "", errors.New("Invalid credentials, PASSWORD NOT MATCH")
	}
	g, ctx := errgroup.WithContext(ctx)

	var accessToken, refreshToken string
	sessionID := uuid.NewString()

	// Encode UserID to hex string for tokens/client
	userIDHexStr := "0x" + hex.EncodeToString(user.UserID)

	g.Go(func() error {
		var err error
		accessToken, err = s.jwtService.GenerateAccessToken(userIDHexStr, sessionID)
		return err
	})

	g.Go(func() error {
		var err error
		refreshToken, err = s.jwtService.GenerateRefreshToken(userIDHexStr, sessionID)
		return err
	})

	if err := g.Wait(); err != nil {
		return "", "", "", err
	}

	s.dispatcher.Dispatch(ctx, event.Event{
		Type: event.LoginEvent,
		Payload: event.LoginPayload{
			UserID: user.UserID, // Bytes for internal event
		},
	})

	return accessToken, refreshToken, userIDHexStr, nil
}
func (s *AuthService) LoginServiceWithUsername(ctx context.Context, username string, password string) (string, string, string, error) {
	return "", "", "", errors.New("not implemented")
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
