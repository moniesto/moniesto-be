package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CreatePayoutInfo(ctx *gin.Context, moniest_id string, req model.CreateMoniestRequest) (db.MoniestPayoutInfo, error) {

	err := validation.BinanceID(req.BinanceID)
	if err != nil {
		return db.MoniestPayoutInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreatePayoutInfo_InvalidBinanceID)
	}

	params := db.CreateMoniestPayoutInfoParams{
		ID:        core.CreateID(),
		MoniestID: moniest_id,
		Source:    db.PayoutSourceBINANCE,
		Type:      db.PayoutTypeBINANCEID,
		Value:     req.BinanceID,
	}

	payout_info, err := service.Store.CreateMoniestPayoutInfo(ctx, params)
	if err != nil {
		return db.MoniestPayoutInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreatePayoutInfo_ServerErrorOnCreate)
	}

	return payout_info, nil
}
