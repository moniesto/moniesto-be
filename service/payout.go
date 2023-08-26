package service

import (
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) CreateBinancePayoutHistories(ctx *gin.Context, defaultParam db.CreateBinancePayoutHistoryParams) {

	params := createBinancePayoutHistoryParams(defaultParam)

	for _, param := range params {
		_, err := service.Store.CreateBinancePayoutHistory(ctx, param)
		if err != nil {
			system.LogError("error while create binance payout history on DB, moniestID:", param.MoniestID, err.Error())
		}
	}
}

// createBinancePayoutHistoryParams creates payout history params for each months
func createBinancePayoutHistoryParams(defaultParam db.CreateBinancePayoutHistoryParams) []db.CreateBinancePayoutHistoryParams {
	numberOfMonth := defaultParam.DateValue

	now := util.Now()

	params := []db.CreateBinancePayoutHistoryParams{}

	// from next month -> + numberOfMonth times months
	for i := 1; i <= int(numberOfMonth); i++ {
		temp_param := defaultParam

		payout_date := now.AddDate(0, i, 0)

		temp_param.ID = core.CreateID()
		temp_param.DateIndex = int32(i)
		temp_param.PayoutDate = payout_date
		temp_param.PayoutYear = int32(payout_date.Year())
		temp_param.PayoutMonth = int32(payout_date.Month())
		temp_param.PayoutDay = int32(payout_date.Day())

		params = append(params, temp_param)
	}

	return params
}
