package service

import (
	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

func (service *Service) CreateBinancePayoutHistories(ctx *gin.Context, param db.CreateBinancePayoutHistoryParams) {

}

func createHistoryParams(param db.CreateBinancePayoutHistoryParams) []db.CreateBinancePayoutHistoryParams {
	// params := []db.CreateBinancePayoutHistoryParams{}

	// create payout history for each month and return

	return nil
}
