package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CreateSubsriptionInfo(ctx *gin.Context, moniest_id string, req model.CreateMoniestRequest) (db.SubscriptionInfo, error) {

	subscriptionInfoParams := db.CreateSubscriptionInfoParams{
		ID:        core.CreateID(),
		MoniestID: moniest_id,
	}

	// STEP: fee is valid
	if err := validation.Fee(req.Fee); err != nil {
		return db.SubscriptionInfo{}, clientError.CreateError(http.StatusBadRequest, clientError.Moniest_CreateSubscriptionInfo_InvalidFee)
	}
	subscriptionInfoParams.Fee = req.Fee

	// STEP: message is valid
	if req.Message != "" {
		if err := validation.SubscriptionMessage(req.Message); err != nil {
			return db.SubscriptionInfo{}, clientError.CreateError(http.StatusBadRequest, clientError.Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage)
		}

		subscriptionInfoParams.Message = sql.NullString{String: req.Message, Valid: true}
	}

	// STEP: create subscription info
	subscriptionInfo, err := service.Store.CreateSubscriptionInfo(ctx, subscriptionInfoParams)
	if err != nil {
		// TODO: add server error
		return db.SubscriptionInfo{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreateSubscriptionInfo_ServerErrorOnCreate)
	}

	return subscriptionInfo, nil
}
