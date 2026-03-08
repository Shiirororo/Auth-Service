package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/auth/domain/sender"
	"github.com/user_service/internal/auth/domain/vo"
	"github.com/user_service/internal/commons"
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/utils/random"
	"github.com/user_service/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type AuthService struct {
	authRepo    repository.AuthRepository
	otpRepo     repository.OTPRepository
	emailSender sender.EmailSender
	blacklist   commons.TokenBlacklist
	jwtService  token.TokenMaker
	dispatcher  *event.Dispatcher
}

func NewAuthService(authRepo repository.AuthRepository, otpRepo repository.OTPRepository, emailSender sender.EmailSender, blacklist commons.TokenBlacklist, jwtService token.TokenMaker, dispatcher *event.Dispatcher) AuthServiceInterface {
	return &AuthService{
		authRepo:    authRepo,
		otpRepo:     otpRepo,
		emailSender: emailSender,
		blacklist:   blacklist,
		jwtService:  jwtService,
		dispatcher:  dispatcher,
	}
}

type AuthServiceInterface interface {
	RegisterService(ctx context.Context, username string, password string, email string) error
	LoginServiceWithUsername(ctx context.Context, username string, password string) (string, string, error)
	LoginServiceWithEmail(ctx context.Context, email vo.Email, password string) (string, string, error)
	LogoutService(ctx context.Context, sessionID string, ttl time.Duration) error
	RefreshService(ctx context.Context, refreshToken string) (string, string, error)
}

func (s *AuthService) RegisterService(ctx context.Context, username string, password string, email string) error {
	// 1. Create Domain Value Objects to validate integrity early
	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return err
	}

	passVo, err := vo.NewPassword(password)
	if err != nil {
		return err
	}

	// 2. Generate new OTP
	otp := random.GenerateOPT6Digit()

	// 3. Save OTP in Redis (5 minutes TTL)
	err = s.otpRepo.SaveOTP(ctx, emailVo, otp, 5*time.Minute)
	if err != nil {
		return err
	}

	// 4. Send OTP to email
	err = s.emailSender.SendOTPEmail(ctx, emailVo, otp)
	if err != nil {
		return err
	}

	// 5. Hash Password explicitly for persistence
	hashPass, err := bcrypt.GenerateFromPassword([]byte(passVo.String()), 10)
	if err != nil {
		return err
	}
	hashedPassVo := vo.RestorePassword(string(hashPass))

	// 6. Generate UUID and construct valid Domain Entity
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	user := entity.NewAuth(id.String(), emailVo, hashedPassVo)

	// 7. Persist to DB
	return s.authRepo.CreateNewUser(ctx, user)
}
func (s *AuthService) LoginServiceWithEmail(ctx context.Context, email vo.Email, password string) (string, string, error) {

	//U DUMB U FORGOT TO REVOKE TOKEN IN THE SAME IP

	user, err := s.authRepo.GetUserByEmail(ctx, vo.Email(email))

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
		accessToken, err = s.jwtService.GenerateAccessToken(string(user.UserID), sessionID)
		return err
	})

	g.Go(func() error {
		var err error
		refreshToken, err = s.jwtService.GenerateRefreshToken(string(user.UserID), sessionID)
		return err
	})

	if err := g.Wait(); err != nil {
		return "", "", err
	}

	s.dispatcher.Dispatch(ctx, event.Event{
		Type:    event.LoginEvent,
		Payload: user.UserID,
	})

	return accessToken, refreshToken, nil
}
func (s *AuthService) LoginServiceWithUsername(ctx context.Context, username string, password string) (string, string, error) {
	return "", "", errors.New("not implemented")
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
