package utils

import (
	"fmt"
	"keyi/config"
	"net/smtp"
	"strings"
)

var smtpAuth smtp.Auth
var smtpAddress string
var debug bool

func init() {
	fmt.Println("init email...")
	if config.Config.EmailHost == "" {
		debug = true
	}

	smtpAuth = smtp.PlainAuth(
		"",
		config.Config.EmailUser,
		config.Config.EmailPassword,
		config.Config.EmailHost,
	)
	smtpAddress = fmt.Sprintf(
		"%s:%d",
		config.Config.EmailHost,
		config.Config.EmailPort,
	)

}

func SendEmail(subject, content string, receivers []string) error {
	body := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		strings.Join(receivers, ","),
		subject,
		content,
	))
	if debug {
		fmt.Println(string(body))
		return nil
	}
	return smtp.SendMail(smtpAddress, smtpAuth, config.Config.EmailUser, receivers, body)
}
