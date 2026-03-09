package repository

import (
	"context"
	"time"
)

// OTPRepository abstracts how OTP codes are stored for registration verification.
type OTPRepository interface {
	SaveOTP(ctx context.Context, email string, otp int, ttl time.Duration) error
	GetOTP(ctx context.Context, email string) (int, error)
	DeleteOTP(ctx context.Context, email string) error
}
