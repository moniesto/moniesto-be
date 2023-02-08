package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Be Moniest
// @Description Turn into moniest
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param CreateMoniestBody body model.CreateMoniestRequest true " "
// @Success 200 {object} model.OwnUser
// @Failure 400 {object} clientError.ErrorResponse "user is already moniest"
// @Failure 403 {object} clientError.ErrorResponse "forbidden operation: email is not verified"
// @Failure 404 {object} clientError.ErrorResponse "not found user"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests [post]
func (server *Server) createMoniest(ctx *gin.Context) {
	var req model.CreateMoniestRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_CreateMoniest_InvalidBody))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is already moniest or not
	userIsMoniest, _, err := server.service.UserIsMoniest(ctx, user_id)
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

// TODO: complete endpoint
func (server *Server) updateMoniestProfile(ctx *gin.Context) {
	var req model.UpdateMoniestProfileRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_UpdateMoniest_InvalidBody))
		return
	}

	/*
		STEPS:
			get user id from token
			update moniest profile
				check user is moniest (by getting moniest info)
				check which values changed
				if Fee changed, send additional request to payment
				update in db
	*/

}
