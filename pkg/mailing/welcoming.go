package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

func SendWelcomingEmail(to string, config config.Config, name string, language db.UserLanguage) error {
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
		system.LogError("server error on sending welcoming email", err.Error())
		return err
	}

	return nil
}
