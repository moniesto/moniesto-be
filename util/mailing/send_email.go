package mailing

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

const mainTemplatePath = "util/mailing/templates/main.html"

func send(toEmails []string, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject string, data any) error {
	contentTpl, err := fillContentTemplate(templatePath, data)
	if err != nil {
		return err
	}

	mainTpl, err := fillMainTemplate(contentTpl)
	if err != nil {
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

	message += "\r\n" + mainTpl.String()

	// STEP: Authentication. Connect to the SMTP Server
	servername := smtpHost + ":" + smtpPort

	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false, // false -> prod | true -> test
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(fromEmail); err != nil {
		return err
	}

	if err = c.Rcpt(toEmails[0]); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	if err := c.Quit(); err != nil {
		return err
	}

	return nil
}

func fillContentTemplate(templatePath string, data any) (bytes.Buffer, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return bytes.Buffer{}, err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return bytes.Buffer{}, err
	}

	return tpl, nil
}

func fillMainTemplate(contentTemplate bytes.Buffer) (bytes.Buffer, error) {
	t, err := template.ParseFiles(mainTemplatePath)
	if err != nil {
		return bytes.Buffer{}, err
	}

	data := struct {
		Body      template.HTML
		ActionUrl string
	}{
		Body: template.HTML(contentTemplate.String()),
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return bytes.Buffer{}, err
	}

	return tpl, nil
}
