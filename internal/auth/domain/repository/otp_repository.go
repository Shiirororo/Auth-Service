package repository

import (
	"context"
	"time"

	"github.com/user_service/internal/auth/domain/vo"
)

// OTPRepository abstracts how OTP codes are stored for registration verification.
type OTPRepository interface {
	SaveOTP(ctx context.Context, email vo.Email, otp int, ttl time.Duration) error
	GetOTP(ctx context.Context, email vo.Email) (int, error)
	DeleteOTP(ctx context.Context, email vo.Email) error
}
