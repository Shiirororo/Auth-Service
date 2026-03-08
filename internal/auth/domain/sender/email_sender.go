package sender

import (
	"context"

	"github.com/user_service/internal/auth/domain/vo"
)

// EmailSender abstracts how emails are transmitted to users.
type EmailSender interface {
	SendOTPEmail(ctx context.Context, destination vo.Email, otp int) error
}
