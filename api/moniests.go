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

	// PAYMENT FUTURE TODO: add card payment info

	// STEP: update data form
	response := model.NewCreateMoniestResponse(createdMoniest)

	ctx.JSON(http.StatusOK, response)
}

// @Summary Update Moniest Profile
// @Description Update Moniest Profile details
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param UpdateMoniestBody body model.UpdateMoniestProfileRequest true "all fields are optional"
// @Success 200 {object} model.OwnUser
// @Failure 403 {object} clientError.ErrorResponse "user is not moniest"
// @Failure 404 {object} clientError.ErrorResponse "user is not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body | invalid bio | invalid desc | invalid fee | invalid message"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/profile [patch]
func (server *Server) updateMoniestProfile(ctx *gin.Context) {
	var req model.UpdateMoniestProfileRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_UpdateMoniest_InvalidBody))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: update moniest
	moniest, err := server.service.UpdateMoniestProfile(ctx, user_id, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: update subscription info [if exist in req body check]
	_, err = server.service.UpdateSubsriptionInfo(ctx, moniest.MoniestID, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get updated moniest data [+ user data]
	updatedMoniest, err := server.service.GetMoniestByMoniestID(ctx, moniest.MoniestID)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: update data form
	response := model.NewCreateMoniestResponse(updatedMoniest)

	ctx.JSON(http.StatusOK, response)
}

// @Summary Subscribe to Moniest
// @Description Subscribe to Moniest
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param username path string true "moniest username"
// @Success 200
// @Failure 400 {object} clientError.ErrorResponse "already subscribed"
// @Failure 403 {object} clientError.ErrorResponse "subscribe own"
// @Failure 404 {object} clientError.ErrorResponse "moniest is not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/subscribe [post]
func (server *Server) subscribeMoniest(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check "username" is a real moniest
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is not subscribing own
	if moniest.ID == user_id {
		ctx.JSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_Subscribe_SubscribeOwn))
		return
	}

	// STEP: create subscription
	err = server.service.SubscribeMoniest(ctx, moniest.MoniestID, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Unsubscribe from Moniest
// @Description Unsubscribe from Moniest
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param username path string true "moniest username"
// @Success 200
// @Failure 400 {object} clientError.ErrorResponse "user not subscribed"
// @Failure 403 {object} clientError.ErrorResponse "unsubscribe own"
// @Failure 404 {object} clientError.ErrorResponse "moniest is not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/unsubscribe [post]
func (server *Server) unsubscribeMoniest(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check "username" is a real moniest
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is not subscribing own
	if moniest.ID == user_id {
		ctx.JSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_Unsubscribe_UnsubscribeOwn))
		return
	}

	// STEP: end subscription
	err = server.service.UnsubscribeMoniest(ctx, moniest.MoniestID, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (server *Server) subscribeMoniestCheck(ctx *gin.Context) {

}
