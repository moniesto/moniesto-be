package mailing

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func send(to []string, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject string, data any) error {
	// STEP: Authentication.
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	// STEP: fill template
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	// STEP: Send email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, body.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
