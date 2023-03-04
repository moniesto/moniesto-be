package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func SendWelcomingEmail(to string, config config.Config, name string) error {
	templatePath := "util/mailing/templates/welcoming.html"
	subject := "Thank You for Joining Moniesto!"

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      name,
		ActionUrl: config.ClientURL,
	}

	err := send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, templatePath, subject, data)
	if err != nil {
		systemError.Log(systemError.InternalMessages["SendWelcomingEmail"](err))
		return err
	}

	return nil
}
