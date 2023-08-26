package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/system"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) GetMoniestByUserID(ctx *gin.Context, user_id string) (db.GetMoniestByUserIdRow, error) {

	// STEP: get moniest by user id
	moniest, err := service.Store.GetMoniestByUserId(ctx, user_id)
	if err != nil {

		// STEP: no moniest with this user id
		if err == sql.ErrNoRows {
			return db.GetMoniestByUserIdRow{}, nil
		}

		system.LogError("server error on getting moniest by user id")
		return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreateMoniest_ServerErrorUserIsMoniest)
	}

	// STEP: double check for moniest id is empty or not
	if moniest.MoniestID == "" && moniest.ID == "" {
		return db.GetMoniestByUserIdRow{}, nil
	}

	return moniest, nil
}

func (service *Service) CreateMoniest(ctx *gin.Context, user_id string, req model.CreateMoniestRequest) (db.Moniest, error) {

	// STEP: create params
	moniestParams := db.CreateMoniestParams{
		ID:     core.CreateID(),
		UserID: user_id,
	}

	// STEP: if bio is added, add to param
	if req.Bio != "" {

		if err := validation.Bio(req.Bio, service.config); err != nil {
			return db.Moniest{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateMoniest_InvalidBio)
		}

		moniestParams.Bio = sql.NullString{String: req.Bio, Valid: true}
	}

	// STEP: if description is added, add to param
	if req.Description != "" {

		if err := validation.Description(req.Description, service.config); err != nil {
			return db.Moniest{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateMoniest_InvalidDescription)
		}

		moniestParams.Description = sql.NullString{String: req.Description, Valid: true}
	}

	// STEP: check all subscription info is valid [double check before creating invalid moniest]
	// STEP: fee is valid
	if err := validation.Fee(req.Fee, service.config); err != nil {
		return db.Moniest{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateSubscriptionInfo_InvalidFee)
	}

	// STEP: message is valid
	if req.Message != "" {
		if err := validation.SubscriptionMessage(req.Message, service.config); err != nil {
			return db.Moniest{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage)
		}
	}

	// STEP: create moniest
	moniest, err := service.Store.CreateMoniest(ctx, moniestParams)
	if err != nil {
		system.LogError("server error on create moniest")
		return db.Moniest{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreateMoniest_ServerErrorCreateMoniest)
	}

	return moniest, nil
}

func (service *Service) GetMoniestByMoniestID(ctx *gin.Context, moniest_id string) (db.GetMoniestByMoniestIdRow, error) {

	// STEP: get moniest by moniest id
	moniest, err := service.Store.GetMoniestByMoniestId(ctx, moniest_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.GetMoniestByMoniestIdRow{}, clientError.CreateError(http.StatusNotFound, clientError.Moniest_GetMoniest_NoMoniest)
		}

		system.LogError("server error on get moniest by moniest ID", err.Error())
		return db.GetMoniestByMoniestIdRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniest_ServerErrorGetMoniest)
	}

	return moniest, nil
}

func (service *Service) GetMoniestByUsername(ctx *gin.Context, username string) (db.GetMoniestByUsernameRow, error) {

	// STEP: get moniest by username
	moniest, err := service.Store.GetMoniestByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.GetMoniestByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.General_MoniestNotFoundByUsername)
		}

		system.LogError("server error on get moniest by moniest username", err.Error())
		return db.GetMoniestByUsernameRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.General_ServerErrorGetMoniestByUsername)
	}

	return moniest, nil
}

func (service *Service) UpdateMoniestProfile(ctx *gin.Context, user_id string, req model.UpdateMoniestProfileRequest) (db.GetMoniestByUserIdRow, error) {
	difference := false

	// STEP: get moniest by user id
	moniest, err := service.Store.GetMoniestByUserId(ctx, user_id)
	if err != nil {
		// STEP: no moniest with this user id
		if err == sql.ErrNoRows {
			return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusForbidden, clientError.General_UserNotMoniest)
		}

		system.LogError("server error on getting moniest by user id")
		return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdateMoniest_ServerErrorGetUser)
	}

	// STEP: update the ones that provided on parameters
	var param db.UpdateMoniestParams = db.UpdateMoniestParams{
		ID:          moniest.MoniestID,
		Bio:         moniest.Bio,
		Description: moniest.Description,
	}

	// STEP: check validity and patch param
	if req.Bio != "" || req.Description != "" {
		difference = true

		if req.Bio != "" {
			if err := validation.Bio(req.Bio, service.config); err != nil {
				return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_UpdateMoniest_InvalidBio)
			}

			param.Bio = sql.NullString{
				Valid:  true,
				String: req.Bio,
			}

			moniest.Bio = param.Bio
		}

		if req.Description != "" {
			if err := validation.Description(req.Description, service.config); err != nil {
				return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Moniest_UpdateMoniest_InvalidDescription)
			}

			param.Description = sql.NullString{
				Valid:  true,
				String: req.Description,
			}

			moniest.Description = param.Description
		}
	}

	// STEP: check if there are new values or not
	if difference {
		// STEP: update moniest
		_, err = service.Store.UpdateMoniest(ctx, param)
		if err != nil {
			system.LogError("server error on update moniest", err.Error())
			return db.GetMoniestByUserIdRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdateMoniest_ServerErrorUpdateMoniest)
		}
	}

	return moniest, nil
}

func (service *Service) GetMoniestPayoutInfos(ctx *gin.Context, user_id string) (model.GetMoniestPayoutInfos, error) {

	// STEP: get moniest by user id
	moniest, err := service.Store.GetMoniestByUserId(ctx, user_id)
	if err != nil {
		// STEP: no moniest with this user id
		if err == sql.ErrNoRows {
			return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusForbidden, clientError.General_UserNotMoniest)
		}

		system.LogError("server error on getting moniest by user id")
		return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetPayoutInfo_ServerErrorGetMoniest)
	}

	payoutInfos, err := service.Store.GetMoniestPayoutInfos(ctx, moniest.MoniestID)
	if err != nil {
		return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetPayoutInfo_ServerErrorGetPayoutInfo)
	}

	if len(payoutInfos) == 0 {
		return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusNotFound, clientError.Moniest_GetPayoutInfo_PayoutInfoNotFound)
	}

	response := model.GetMoniestPayoutInfos{
		PayoutMethods: model.PayoutMethods{
			PayoutMethodBinance: []model.PayoutMethodBinance{
				{
					Type:  string(payoutInfos[0].Type),
					Value: payoutInfos[0].Value,
				},
			},
		},
	}

	return response, nil
}

func (service *Service) UpdateMoniestPayoutInfo(ctx *gin.Context, user_id string, req model.UpdateMoniestPayoutInfo) (model.GetMoniestPayoutInfos, error) {

	// STEP: get moniest by user id
	moniest, err := service.Store.GetMoniestByUserId(ctx, user_id)
	if err != nil {
		// STEP: no moniest with this user id
		if err == sql.ErrNoRows {
			return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusForbidden, clientError.General_UserNotMoniest)
		}

		system.LogError("server error on getting moniest by user id")
		return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdatePayout_ServerErrorGetMoniest)
	}

	params := db.UpdateMoniestPayoutInfoParams{
		MoniestID: moniest.MoniestID,
		Source:    db.PayoutSourceBINANCE,
		Type:      db.PayoutTypeBINANCEID,
		Value:     req.BinanceID,
	}

	err = service.Store.UpdateMoniestPayoutInfo(ctx, params)
	if err != nil {
		return model.GetMoniestPayoutInfos{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_UpdatePayout_ServerErrorUpdatePayoutInfo)
	}

	return service.GetMoniestPayoutInfos(ctx, user_id)
}

func (service *Service) CheckUserIsMoniestByUsername(ctx *gin.Context, username string) (bool, error) {

	// STEP: check user is moniest
	userIsMoniest, err := service.Store.CheckUserIsMoniestByUsername(ctx, username)
	if err != nil {
		system.LogError("server error on check user is moniest by username", err.Error())
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.General_ServerErrorCheckMoniestByUsername)
	}

	return userIsMoniest, nil
}

func (service *Service) CheckUserIsMoniestByUserID(ctx *gin.Context, user_id string) (bool, error) {

	// STEP: check user is moniest
	userIsMoniest, err := service.Store.CheckUserIsMoniestByID(ctx, user_id)
	if err != nil {
		system.LogError("server error on check user is moniest by ID", err.Error())
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.General_ServerErrorCheckMoniestByUserID)
	}

	return userIsMoniest, nil
}

func (service *Service) GetMoniestStats(ctx *gin.Context, username string) (model.MoniestStatResponse, error) {
	stats, err := service.Store.GetMoniestStatsByUsername(ctx, username)
	if err != nil {
		system.LogError("server error on get moniest stats by username", err.Error())
		return model.MoniestStatResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetStats_ServerErrorGetStats)
	}

	return model.MoniestStatResponse{
		SubscriptionCount: stats.SubscriptionCount,
		SubscriberCount:   stats.SubscriberCount,
		PostCount:         stats.PostCount,
	}, nil
}
