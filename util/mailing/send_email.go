package mailing

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

func send(toEmails []string, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject string, data any) error {
	// STEP: fill template
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}

	from := mail.Address{Name: "Moniesto", Address: fromEmail}
	to := mail.Address{Name: "", Address: toEmails[0]}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0;"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\";"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + tpl.String()

	// STEP: Authentication. Connect to the SMTP Server
	servername := smtpHost + ":" + smtpPort

	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(fromEmail); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(toEmails[0]); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	if err := c.Quit(); err != nil {
		log.Panic(err)
	}

	return nil
}
