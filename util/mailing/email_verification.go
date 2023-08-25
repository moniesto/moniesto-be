package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func SendEmailVerificationEmail(to string, config config.Config, fullname, token string, language db.UserLanguage) error {
	template, err := GetTemplate("email_verification", language)
	if err != nil {
		return err
	}

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      fullname,
		ActionUrl: createEmailVerificationURL(config.ClientURL, token),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		systemError.Log("server error on sending email verification email", err.Error())
		return err
	}

	return nil
}

func createEmailVerificationURL(url, token string) string {
	return url + "/" + verifyEmailURL + "?token=" + token
}
