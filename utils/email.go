package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"keyi/config"
	"net/smtp"
)

var auth smtp.Auth
var smtpAddr string

func init() {
	fmt.Println("init email config...")
	auth = smtp.PlainAuth(
		"",
		config.Config.SmtpUser,
		config.Config.SmtpPassword,
		config.Config.SmtpHost,
	)
	smtpAddr = fmt.Sprintf(
		"%s:%d",
		config.Config.SmtpHost,
		config.Config.SmtpPort,
	)
}

func SendEmail(to []string, subject, body string) error {
	mail := email.NewEmail()
	mail.From = config.Config.FromEmail
	mail.To = to
	mail.Subject = subject
	mail.HTML = []byte(body)

	if config.Config.Mode != "production" {
		bytes, err := mail.Bytes()
		fmt.Println(string(bytes))
		return err
	}

	return mail.Send(smtpAddr, auth)
}
