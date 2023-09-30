package mailing

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"github.com/moniesto/moniesto-be/util/system"
)

const MAX_ATTEMPTS = 10

const mainTemplatePath = "util/mailing/templates/main.html"

func send(toEmails []string, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject string, data any) error {
	var err error

	for i := 0; i < MAX_ATTEMPTS; i++ {
		err = sendEmail(toEmails, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject, data)
		if err == nil {
			return nil
		}

		if err != nil {
			system.Log(fmt.Sprintf("Sending Email: Attempt %d failed with error: %v\n", i+1, err.Error()))
			time.Sleep(time.Second) // Add a delay between attempts
		}
	}

	return err
}

func sendEmail(toEmails []string, fromEmail, fromPassword, smtpHost, smtpPort, templatePath, subject string, data any) error {
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

	// Authenticate and send the email
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	err = sendMailTLS(servername, auth, fromEmail, toEmails, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// sendMailTLS not use STARTTLS commond
func sendMailTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	tlsconfig := &tls.Config{ServerName: host}
	if err = validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err = validateLine(recp); err != nil {
			return err
		}
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("a line must not contain CR or LF")
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
