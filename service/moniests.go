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

func (service *Service) UserIsMoniest(ctx *gin.Context, user_id string) (bool, error) {

	// STEP: get moniest by user id
	moniest, err := service.Store.GetMoniestByUserId(ctx, user_id)
	if err != nil {

		// STEP: no moniest with this user id
		if err == sql.ErrNoRows {
			return false, nil
		}

		// TODO: add server error
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreateMoniest_ServerErrorUserIsMoniest)
	}

	// STEP: double check for moniest id is empty or not
	if moniest.MoniestID == "" && moniest.ID == "" {
		return false, nil
	}

	return true, nil
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
		// TODO: add server error
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

		return db.GetMoniestByMoniestIdRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniest_ServerErrorGetMoniest)
	}

	return moniest, nil
}
