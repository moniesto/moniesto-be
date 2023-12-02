package mailing

import (
	"time"

	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

func SendSubscriptionExpiredEmail(to string, config config.Config, fullname_user, fullname_moniest, username string, subscriptionStartDate, subscriptionEndDate time.Time, subscriptionFee float64, subscriptionMonth int, language db.UserLanguage) error {
	template, err := GetTemplate("subscription_expired", language)
	if err != nil {
		return err
	}

	data := struct {
		NameUser              string
		NameMoniest           string
		Username              string
		SubscriptionStartDate string
		SubscriptionEndDate   string
		SubscriptionFee       float64
		SubscriptionMonth     int
		ActionUrl             string
	}{
		NameUser:              fullname_user,
		NameMoniest:           fullname_moniest,
		Username:              username,
		SubscriptionStartDate: core.FormatDate(subscriptionStartDate, core.DD_MM_YYYY),
		SubscriptionEndDate:   core.FormatDate(subscriptionEndDate, core.DD_MM_YYYY),
		SubscriptionFee:       subscriptionFee,
		SubscriptionMonth:     subscriptionMonth,
		ActionUrl:             createMoniestPageURL(config.ClientURL, username),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		system.LogError("server error on sending subscription expired email", err.Error())
		return err
	}

	return nil
}
