package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func SendWelcomingEmail(to string, config config.Config, name string, language model.UserLanguage) error {
	template, err := GetTemplate("welcoming", language)
	if err != nil {
		return err
	}

	data := struct {
		Name      string
		ActionUrl string
	}{
		Name:      name,
		ActionUrl: config.ClientURL,
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		systemError.Log(systemError.InternalMessages["SendWelcomingEmail"](err))
		return err
	}

	return nil
}
