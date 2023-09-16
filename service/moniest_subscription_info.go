package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/system"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CreateSubsriptionInfo(ctx *gin.Context, moniest_id string, req model.CreateMoniestRequest) (db.MoniestSubscriptionInfo, error) {

	subscriptionInfoParams := db.CreateMoniestSubscriptionInfoParams{
		ID:        core.CreateID(),
		MoniestID: moniest_id,
	}

	// STEP: fee is valid
	if err := validation.Fee(req.Fee, service.config); err != nil {
		return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateSubscriptionInfo_InvalidFee)
	}
	subscriptionInfoParams.Fee = util.RoundAmountDown(req.Fee)

	// STEP: message is valid
	if req.Message != "" {
		if err := validation.SubscriptionMessage(req.Message, service.config); err != nil {
			return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage)
		}

		subscriptionInfoParams.Message = sql.NullString{String: req.Message, Valid: true}
	}

	// STEP: create subscription info
	subscriptionInfo, err := service.Store.CreateMoniestSubscriptionInfo(ctx, subscriptionInfoParams)
	if err != nil {
		system.LogError("create moniest subscription info error", err.Error())
		return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreateSubscriptionInfo_ServerErrorOnCreate)
	}

	return subscriptionInfo, nil
}

func (service *Service) UpdateSubsriptionInfo(ctx *gin.Context, moniest_id string, req model.UpdateMoniestProfileRequest) (db.MoniestSubscriptionInfo, error) {

	// STEP: subscription info is valid
	if req.MoniestSubscriptionInfo == nil {
		return db.MoniestSubscriptionInfo{}, nil
	}

	subscription_info, err := service.Store.GetMoniestSubscriptionInfoByMoniestId(ctx, moniest_id)
	if err != nil {
		system.LogError("get moniest subscription info by moniest ID error", err.Error())
		return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdateMoniest_ServerErrorGetSubscriptionInfo)
	}

	subscriptionUpdateInfoParams := db.UpdateMoniestSubscriptionInfoParams{
		MoniestID: moniest_id,
		Fee:       subscription_info.Fee,
		Message:   subscription_info.Message,
	}

	// STEP: fee is valid / updated
	if req.MoniestSubscriptionInfo.Fee != 0 && subscriptionUpdateInfoParams.Fee != req.MoniestSubscriptionInfo.Fee {
		if err := validation.Fee(req.MoniestSubscriptionInfo.Fee, service.config); err != nil {
			return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_UpdateMoniest_InvalidFee)
		} else {
			subscriptionUpdateInfoParams.Fee = req.MoniestSubscriptionInfo.Fee
		}
	}

	// STEP: message is valid / updated
	if req.MoniestSubscriptionInfo.Message != "" {
		if err := validation.SubscriptionMessage(req.MoniestSubscriptionInfo.Message, service.config); err != nil {
			return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_UpdateMoniest_InvalidSubscriptionMessage)
		} else {
			subscriptionUpdateInfoParams.Message = sql.NullString{
				Valid:  true,
				String: req.MoniestSubscriptionInfo.Message,
			}
		}
	}

	// STEP: update subsctription info in DB
	updated_subscription_info, err := service.Store.UpdateMoniestSubscriptionInfo(ctx, subscriptionUpdateInfoParams)
	if err != nil {
		return db.MoniestSubscriptionInfo{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdateMoniest_ServerErrorUpdateSubscriptionInfo)
	}

	return updated_subscription_info, nil
}
