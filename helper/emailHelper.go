package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/edulink-api/public"
	"gopkg.in/gomail.v2"
)

func SendResetTokenEmail(email string, userName string, resetTokenLink string) error {
	tmpl, err := template.New("resetPassword").Parse(public.HtmlTemplate)
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
		return fmt.Errorf("error parsing email port: %v", err)
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
