package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// CreateMoniest creates moniest if do
func (server *Server) CreateMoniest(ctx *gin.Context) {
	var req model.CreateMoniestRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_CreateMoniest_InvalidBody))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is already moniest or not
	userIsMoniest, err := server.service.UserIsMoniest(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}
	if userIsMoniest {
		ctx.JSON(http.StatusBadRequest, clientError.GetError(clientError.Moniest_CreateMoniest_UserIsAlreadyMoniest))
		return
	}

	// STEP: check the email of user is verified
	user, err := server.service.GetUserByID(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}
	if !user.EmailVerified {
		ctx.JSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_CreateMoniest_UnverifiedEmail))
		return
	}

	/*
		TODO:
			create moniest
			create subscription info
			add card payment info

		should be converted to db transaction
		[one fail, all should fail, done ones should be converted back]
	*/

	// STEP: create moniest
	moniest, err := server.service.CreateMoniest(ctx, user_id, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: create subscription info
	_, err = server.service.CreateSubsriptionInfo(ctx, moniest.ID, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get created moniest data [+ user data]
	createdMoniest, err := server.service.GetMoniestByMoniestID(ctx, moniest.ID)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// TODO: add card payment info

	// STEP: update data form
	response := model.NewCreateMoniestResponse(createdMoniest)

	ctx.JSON(http.StatusOK, response)
}
