package services

import (
	"os"

	"github.com/resendlabs/resend-go"
)

var resendClient *resend.Client

func InitEmail() {
	resendClient = resend.NewClient(os.Getenv("RESEND_API_KEY"))
}

func SendEmail(to, subject, html string) error {
	params := &resend.SendEmailRequest{
		From:    "Directory Surf <noreply@directory.surf>",
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	_, err := resendClient.Emails.Send(params)
	return err
}
