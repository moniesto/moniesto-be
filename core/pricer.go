package core

import "github.com/moniesto/moniesto-be/util"

func GetTotalAmount(numberOfMonths int, fee float64) float64 {
	return util.RoundAmountUp(float64(numberOfMonths) * fee)
}

func GetAmountAfterCommission(amount float64, percentage float64) float64 {
	return util.RoundAmountDown(amount - ((amount * percentage) / 100))
}

func GetAmountOfCommission(amount float64, percentage float64) float64 {
	return util.RoundAmountDown((amount * percentage) / 100)
}
