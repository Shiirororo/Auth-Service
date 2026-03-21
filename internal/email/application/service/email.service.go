package email_service

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sendgrid/sendgrid-go"
)

type EmailService struct {
}

type EmailServiceInterface interface {
	SendOTP(email, otp string) error
}

func NewEmailService() EmailServiceInterface {
	return &EmailService{}
}
func (e *EmailService) SendOTP(email, otp string) error {
	from := mail.NewEmail("Your App", "no-reply@yourapp.com")
	subject := "Your OTP Code"
	to := mail.NewEmail("", email)

	content := mail.NewContent("text/plain", "OTP: "+otp)
	msg := mail.NewV3MailInit(from, subject, to, content)

	client := sendgrid.NewSendClient("API_KEY")
	_, err := client.Send(msg)
	return err
} //MOCK
