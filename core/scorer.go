package core

import (
	"fmt"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/validation"
)

// CalculatePNL_ROI calculated the profit or loss (PNL) and ROI based on the [startPrice, lastPrice, leverage, direction]
func CalculatePNL_ROI(startPrice, lastPrice float64, leverage int32, direction db.Direction) (pnl float64, roi float64, err error) {
	// Calculate the profit or loss (PNL) based on the trade direction
	switch direction {
	case db.DirectionLong:
		pnl = (lastPrice - startPrice) * float64(leverage) * (validation.InvestmentAmount / startPrice)
	case db.DirectionShort:
		pnl = (startPrice - lastPrice) * float64(leverage) * (validation.InvestmentAmount / startPrice)
	default:
		return 0, 0, fmt.Errorf(validation.ERROR_DirectionNotValid)
	}

	// Calculate the return on investment (ROI)
	roi = (pnl / validation.InvestmentAmount / float64(leverage)) * 100

	roi = roi * float64(leverage)

	return util.RoundAmountDown(pnl), util.RoundAmountDown(roi), nil
}
