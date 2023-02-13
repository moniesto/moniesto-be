package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Create Feedback
// @Description [AUTH OPTIONAL] Create Feedback as Authenticated or Anonymous
// @Tags Feedback
// @Security bearerAuth || ""
// @Accept json
// @Produce json
// @Param CreateFeedbackBody body model.CreateFeedbackRequest true "type is optional"
// @Success 200
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /feedback [post]
func (server *Server) createFeedback(ctx *gin.Context) {

	var req model.CreateFeedbackRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Feedback_CreateFeedback_InvalidBody))
		return
	}

	user_id := ""

	// STEP: get user_id if user logged in
	authValidity := ctx.MustGet(authorizationPayloadValidityKey).(bool)
	if authValidity {
		// get user id from token
		authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
		user_id = authPayload.User.ID
	}

	// STEP: create feedback
	err := server.service.CreateFeedback(ctx, user_id, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}
