package mailing

import (
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

func SendPayoutEmail(to string, config config.Config, fullname_user, username_user, fullname_moniest, binanceID string, currentMonth, totalMonth int, subscribedFee, operationFeePercentage, amount float64, language db.UserLanguage) error {

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
		Username:               username_user,
		BinanceID:              binanceID,
		CurrentMonth:           currentMonth,
		TotalMonth:             totalMonth,
		SubscribedFee:          subscribedFee,
		OperationFeePercentage: operationFeePercentage,
		OperationFee:           10, // TODO calculate
		Amount:                 amount,
	}

	err = send([]string{to}, config.NoReplyEmail, config.NoReplyPassword, config.SmtpHost, config.SmtpPort, template.Path, template.Subject, data)
	if err != nil {
		system.LogError("Server error on sending payout email", err)
		return err
	}

	return nil
}
