package core

import (
	"fmt"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/validation"
)

func CalculatePNL_ROI(startPrice, lastPrice float64, leverage int32, direction db.EntryPosition) (pnl float64, roi float64, err error) {
	// Calculate the profit or loss (PNL) based on the trade direction
	switch direction {
	case db.EntryPositionLong:
		pnl = (lastPrice - startPrice) * float64(leverage)
	case db.EntryPositionShort:
		pnl = (startPrice - lastPrice) * float64(leverage)
	default:
		return 0, 0, fmt.Errorf(validation.ERROR_DirectionNotValid)
	}

	// Calculate the return on investment (ROI)
	roi = (pnl / (startPrice * float64(leverage))) * 100

	return pnl, roi, nil
}
