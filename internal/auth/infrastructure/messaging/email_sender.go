package messaging

import (
	"context"
	"fmt"
	"log"

	"github.com/user_service/internal/auth/domain/sender"
	"github.com/user_service/internal/auth/domain/vo"
)

type mockEmailSender struct{}

func NewMockEmailSender() sender.EmailSender {
	return &mockEmailSender{}
}

func (s *mockEmailSender) SendOTPEmail(ctx context.Context, destination vo.Email, otp int) error {
	// In a real application, this would use an SMTP client, SendGrid, AWS SES, etc.
	// We log it here to verify the abstract interface is operating correctly within the Domain.
	log.Printf("=========================================================\n")
	log.Printf("📧 MOCK EMAIL DISPATCHED\n")
	log.Printf("TO: %s\n", destination.String())
	log.Printf("SUBJECT: Your Registration OTP\n")
	log.Printf("BODY: Your One-Time Password is: %06d. It expires in 5 minutes.\n", otp)
	log.Printf("=========================================================\n")

	// Print to stdout specifically so the user can easily see it in the terminal logs
	fmt.Printf("[EmailSender] Sent OTP %06d to %s\n", otp, destination.String())

	return nil
}
