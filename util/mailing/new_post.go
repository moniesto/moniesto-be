package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func SendNewPostEmail(to string, config config.Config, fullname_user, fullname_moniest, username, currency string, language model.UserLanguage) error {
	template, err := GetTemplate("new_post", language)
	if err != nil {
		return err
	}

	data := struct {
		NameUser    string
		NameMoniest string
		Username    string
		Currency    string
		ActionUrl   string
	}{
		NameUser:    fullname_user,
		NameMoniest: fullname_moniest,
		Username:    username,
		Currency:    currency,
		ActionUrl:   createMoniestPageURL(config.ClientURL, username),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		systemError.Log("Server error on sending new post email", err)
		return err
	}

	return nil
}

func createMoniestPageURL(url, username string) string {
	return url + "/" + username
}
