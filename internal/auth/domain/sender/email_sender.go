package sender

import (
	"context"
)

// EmailSender abstracts how emails are transmitted to users.
type EmailSender interface {
	SendOTPEmail(ctx context.Context, destination string, otp int) error
}
