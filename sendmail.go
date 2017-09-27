package chaos

import (
	"net/smtp"
)

type Email struct {
	UserName   string
	Host       string
	Port       string
	Password   string
	From       string
	To         []string
	EmailAlias string
	Subject    string
}

func SendEmail(email *Email, message string) error {
	b := []byte(
		"From: " + email.EmailAlias + "<" + email.From + ">\r\nSubject: " +
			email.Subject + " message" + "\r\n" + "\r\n\r\n" + message)

	auth := smtp.PlainAuth("", email.UserName, email.Password, email.Host)
	return smtp.SendMail(
		email.Host,
		auth,
		email.From,
		email.To,
		b,
	)
}
