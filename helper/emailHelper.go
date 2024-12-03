package helper

import (
	"bytes"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendResetTokenEmail(email string, userName string, resetTokenLink string) error {
	templatePath := "public/reset-password.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	data := struct {
		ResetTokenLink string
		UserName       string
	}{
		ResetTokenLink: resetTokenLink,
		UserName:       userName,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("MAIL_SENDER_NAME"))
	mailer.SetHeader("To", email)
	mailer.SetAddressHeader("Cc", "rifianfernando19@gmail.com", "Reset Password")
	mailer.SetHeader("Subject", "Password reset requested")
	mailer.SetBody("text/html", body.String())

	EmailPort, err := strconv.ParseInt(os.Getenv("MAIL_PORT"), 10, 64)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		int(EmailPort),
		os.Getenv("MAIL_EMAIL"),
		os.Getenv("MAIL_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
