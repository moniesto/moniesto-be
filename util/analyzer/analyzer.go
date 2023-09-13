package analyzer

import (
	"fmt"
	"strconv"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/crypto"
	"github.com/moniesto/moniesto-be/util/system"
)

// Analyze
func Analyze(symbol string, takeProfit, stopPrice float64, tradeStart_UTC_TS, tradeEnd_UTC_TS int64, direction db.Direction) (db.PostCryptoStatus, float64, int64, error) {

	chunkSize := 1000
	limit := 1000
	currentTime := tradeStart_UTC_TS

	count := 0
	var lastClosePrice float64
	var lastCloseDate_UTC_TS int64

	for currentTime < tradeEnd_UTC_TS {
		count += 1
		system.Log("chunk no: ", count)

		histories, err := crypto.GetHistories(symbol, string(db.PostCryptoMarketTypeSpot), crypto.INTERVAL_1second, currentTime, tradeEnd_UTC_TS, limit)
		if err != nil {
			system.LogError("error happened while getting history", err.Error())
			return db.PostCryptoStatusPending, lastClosePrice, lastCloseDate_UTC_TS, err
		}

		// STEP: no new history data found
		if len(histories) == 0 {
			system.Log("stop: 0 history")
			return db.PostCryptoStatusPending, lastClosePrice, lastCloseDate_UTC_TS, nil
		}

		for _, history := range histories {
			closePrice, closeDate_UTC_TS, err := parseHistory(history)
			if err != nil {
				system.LogError("error happened while parsing history data", err.Error())
				return db.PostCryptoStatusPending, lastClosePrice, lastCloseDate_UTC_TS, err
			}

			lastClosePrice = closePrice
			lastCloseDate_UTC_TS = closeDate_UTC_TS

			status := lookupStatus(closePrice, takeProfit, stopPrice, direction)
			// STEP: target or stop hit
			if status != nil {
				system.Log("stop: found")
				return *status, lastClosePrice, lastCloseDate_UTC_TS, nil
			}
		}

		// STEP: no hit for this chunk -> expand start(current) time
		currentTime += int64(chunkSize * 1000)
	}

	system.Log("stop: no hit")
	// STEP: no hit for whole time period
	return db.PostCryptoStatusPending, lastClosePrice, lastCloseDate_UTC_TS, nil
}

func parseHistory(history model.History) (float64, int64, error) {
	closePrice := history[4]
	closePriceStr := closePrice.(string)
	closePriceFloat, err := strconv.ParseFloat(closePriceStr, 64)

	closeDate := history[0]
	closeDateStr := fmt.Sprintf("%.0f", closeDate)
	closeDate_UTC_TS, err := strconv.ParseInt(closeDateStr, 10, 64)

	return closePriceFloat, closeDate_UTC_TS, err
}

func lookupStatus(closePrice float64, takeProfit, stopPrice float64, direction db.Direction) *db.PostCryptoStatus {
	success := db.PostCryptoStatusSuccess
	fail := db.PostCryptoStatusFail

	switch direction {

	// long case
	case db.DirectionLong:
		if closePrice >= takeProfit {
			return &success
		} else if closePrice <= stopPrice {
			return &fail
		}
		return nil

	// short case
	case db.DirectionShort:
		if closePrice <= takeProfit {
			return &success
		} else if closePrice >= stopPrice {
			return &fail
		}
		return nil
	}

	return nil
}
