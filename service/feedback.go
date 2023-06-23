package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) CreateFeedback(ctx *gin.Context, userID string, req model.CreateFeedbackRequest) error {

	// STEP: create params [userID & type is optional]
	params := db.CreateFeedbackParams{
		ID:      core.CreateID(),
		Message: req.Message,
	}

	if userID != "" {
		params.UserID = sql.NullString{
			Valid:  true,
			String: userID,
		}
	}

	if req.Type != "" {
		params.Type = sql.NullString{
			Valid:  true,
			String: req.Type,
		}
	}

	// STEP: create feedback in db
	err := service.Store.CreateFeedback(ctx, params)
	if err != nil {
		systemError.Log("server error on create feedback", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Feedback_CreateFeedback_ServerErrorCreateFeedback)
	}

	return nil
}
