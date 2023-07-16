package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/util/systemError"
)

var resetPasswordURL string = "change-password"
var verifyEmailURL string = "verify-email"

func SendPasswordResetEmail(to string, config config.Config, fullname, token string) error {
	templatePath := "util/mailing/templates/password_reset.html"
	subject := "Moniesto: Reset password"

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      fullname,
		ActionUrl: createResetPasswordURL(config.ClientURL, token),
	}

	err := send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, templatePath, subject, data)
	if err != nil {
		systemError.Log(systemError.InternalMessages["SendPasswordResetEmail"](err))
		return err
	}

	return nil
}

func createResetPasswordURL(url, token string) string {
	return url + "/" + resetPasswordURL + "?token=" + token
}
