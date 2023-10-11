package core

import (
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/util"
)

var TEST_AMOUNT = 0.00000001

func GetTotalAmount(numberOfMonths int, fee float64, config *config.Config) float64 {
	if config.IsProd() {
		return util.RoundAmountUp(float64(numberOfMonths) * fee)
	}

	return TEST_AMOUNT
}

func GetAmountAfterCommission(amount float64, percentage float64, config *config.Config) float64 {
	if config.IsProd() {
		return util.RoundAmountDown(amount - ((amount * percentage) / 100))
	}

	return TEST_AMOUNT
}

func GetAmountOfCommission(amount float64, percentage float64) float64 {
	return util.RoundAmountDown((amount * percentage) / 100)
}
