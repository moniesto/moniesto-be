package mailing

import (
	"time"

	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

func SendUnsubscribeEmail(to string, config config.Config, fullname_user, fullname_moniest, username, payerID string, subscriptionCancelDate time.Time, subscriptionFee, operationFeePercentage float64, remainingMonth int, totalAmount float64, language db.UserLanguage) error {

	template, err := GetTemplate("unsubscribe", language)
	if err != nil {
		return err
	}

	data := struct {
		NameUser               string
		NameMoniest            string
		Username               string
		PayerID                string
		SubscriptionCancelDate string
		SubscriptionFee        float64
		RemainingMonth         int
		OperationFeePercentage float64
		OperationFee           float64
		Amount                 float64
	}{
		NameUser:               fullname_user,
		NameMoniest:            fullname_moniest,
		Username:               username,
		PayerID:                payerID,
		SubscriptionCancelDate: core.FormatDate(subscriptionCancelDate, core.DD_MM_YYYY),
		SubscriptionFee:        subscriptionFee,
		RemainingMonth:         remainingMonth,
		OperationFeePercentage: operationFeePercentage,
		OperationFee:           core.GetAmountOfCommission((subscriptionFee * float64(remainingMonth)), operationFeePercentage),
		Amount:                 core.GetAmountAfterCommission(totalAmount, operationFeePercentage, &config),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		system.LogError("Server error on sending ubsubscribe email", err.Error())
		return err
	}

	return nil
}
