package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func SendEmailVerificationEmail(to string, config config.Config, name, token string) error {
	templatePath := "util/mailing/templates/email_verification.html"
	subject := "Moniesto: Email Verification"

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      name,
		ActionUrl: createEmailVerificationURL(config.ClientURL, token),
	}

	err := send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, templatePath, subject, data)
	if err != nil {
		systemError.Log(systemError.InternalMessages["SendEmailVerificationEmail"](err))
		return err
	}

	return nil
}

func createEmailVerificationURL(url, token string) string {
	return url + "/" + verifyEmailURL + "?token=" + token
}
