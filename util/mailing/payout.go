package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

func SendPayoutEmail(to string, config config.Config, fullname_user, username, fullname_moniest, binanceID string, currentMonth, totalMonth int, subscribedFee, operationFeePercentage float64, language db.UserLanguage) error {

	template, err := GetTemplate("payout", language)
	if err != nil {
		return err
	}

	data := struct {
		NameMoniest            string
		NameUser               string
		Username               string
		BinanceID              string
		CurrentMonth           int
		TotalMonth             int
		SubscribedFee          float64
		OperationFeePercentage float64
		OperationFee           float64
		Amount                 float64
	}{
		NameMoniest:            fullname_moniest,
		NameUser:               fullname_user,
		Username:               username,
		BinanceID:              binanceID,
		CurrentMonth:           currentMonth,
		TotalMonth:             totalMonth,
		SubscribedFee:          subscribedFee,
		OperationFeePercentage: operationFeePercentage,
		OperationFee:           core.GetAmountOfCommission(subscribedFee, operationFeePercentage),
		Amount:                 core.GetAmountAfterCommission(subscribedFee, operationFeePercentage),
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		system.LogError("Server error on sending payout email", err)
		return err
	}

	return nil
}
