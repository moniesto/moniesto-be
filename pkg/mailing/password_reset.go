package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

var resetPasswordURL string = "change-password"
var verifyEmailURL string = "verify-email"

func SendPasswordResetEmail(to string, config config.Config, fullname, token string, language db.UserLanguage) error {
	template, err := GetTemplate("password-reset", language)
	if err != nil {
		return err
	}

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      fullname,
		ActionUrl: createResetPasswordURL(config.ClientURL, token),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		system.LogError("server error on sending password reset email", err.Error())
		return err
	}

	return nil
}

func createResetPasswordURL(url, token string) string {
	return url + "/" + resetPasswordURL + "?token=" + token
}
