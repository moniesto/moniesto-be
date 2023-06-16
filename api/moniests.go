package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
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
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_CreateMoniest_InvalidBody))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is already moniest or not
	userIsMoniest, err := server.service.CheckUserIsMoniestByUserID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}
	if userIsMoniest {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, clientError.GetError(clientError.Moniest_CreateMoniest_UserIsAlreadyMoniest))
		return
	}

	// STEP: check the email of user is verified
	user, err := server.service.GetUserByID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}
	if !user.EmailVerified {
		ctx.AbortWithStatusJSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_CreateMoniest_UnverifiedEmail))
		return
	}

	// STEP: create moniest
	moniest, err := server.service.CreateMoniest(ctx, user_id, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create subscription info
	_, err = server.service.CreateSubsriptionInfo(ctx, moniest.ID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create payout info
	_, err = server.service.CreatePayoutInfo(ctx, moniest.ID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get created moniest data [+ user data]
	createdMoniest, err := server.service.GetMoniestByMoniestID(ctx, moniest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
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
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_UpdateMoniest_InvalidBody))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: update moniest
	moniest, err := server.service.UpdateMoniestProfile(ctx, user_id, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: update subscription info [if exist in req body check]
	_, err = server.service.UpdateSubsriptionInfo(ctx, moniest.MoniestID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get updated moniest data [+ user data]
	updatedMoniest, err := server.service.GetMoniestByMoniestID(ctx, moniest.MoniestID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
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
// @Param UpdateMoniestBody body model.SubscribeMoniestRequest true "all fields are required"
// @Success 200 {object} model.SubscribeMoniestResponse
// @Failure 400 {object} clientError.ErrorResponse "already subscribed"
// @Failure 403 {object} clientError.ErrorResponse "subscribe own"
// @Failure 404 {object} clientError.ErrorResponse "moniest is not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/subscribe [post]
func (server *Server) subscribeMoniest(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	var req model.SubscribeMoniestRequest
	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_Subscribe_InvalidBody))
		return
	}

	// STEP: valid date value
	if !validation.SubscriptionDateValue(req.NumberOfMonths) {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_Subscribe_InvalidBody))
		return
	}

	// STEP: check "username" is a real moniest
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is not subscribing own
	if moniest.ID == user_id {
		ctx.AbortWithStatusJSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_Subscribe_SubscribeOwn))
		return
	}

	// STEP: check subscription status -> prevent already subscribed
	exist, subscription, err := server.service.GetUserSubscriptionStatus(ctx, moniest.MoniestID, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	if exist {
		if subscription.Active {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, clientError.Moniest_Subscribe_AlreadySubscribed)
			return
		}
	}

	// STEP: create binance payment transaction
	binancePaymentTransaction, err := server.service.CreateBinancePaymentTransaction(ctx, req, moniest, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create job to check order status
	// TODO: create checker job

	response := model.SubscribeMoniestResponse{
		QrcodeLink:    binancePaymentTransaction.QrcodeLink,
		CheckoutLink:  binancePaymentTransaction.CheckoutLink,
		DeepLink:      binancePaymentTransaction.DeepLink,
		UniversalLink: binancePaymentTransaction.UniversalLink,
	}

	ctx.JSON(http.StatusOK, response)
}

// offline
func (server *Server) subscribeMoniest1(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check "username" is a real moniest
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is not subscribing own
	if moniest.ID == user_id {
		ctx.AbortWithStatusJSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_Subscribe_SubscribeOwn))
		return
	}

	// STEP: create subscription
	err = server.service.SubscribeMoniest(ctx, moniest.MoniestID, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
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
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is not subscribing own
	if moniest.ID == user_id {
		ctx.AbortWithStatusJSON(http.StatusForbidden, clientError.GetError(clientError.Moniest_Unsubscribe_UnsubscribeOwn))
		return
	}

	// STEP: end subscription
	err = server.service.UnsubscribeMoniest(ctx, moniest.MoniestID, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Check Subscribe status
// @Description Check Subscribe status to Moniest
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param username path string true "moniest username"
// @Success 200 {object} model.CheckSubscriptionResponse "subscribed: true | false information"
// @Failure 404 {object} clientError.ErrorResponse "moniest not found with this username"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/subscribe/check [get]
func (server *Server) subscribeMoniestCheck(ctx *gin.Context) {

	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check user is moniest
	userIsMoniest, err := server.service.CheckUserIsMoniestByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}
	if !userIsMoniest {
		ctx.AbortWithStatusJSON(http.StatusNotFound, clientError.GetError(clientError.General_MoniestNotFoundByUsername))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is subscribed to moniest
	userIsSubscribed, err := server.service.CheckUserSubscriptionByMoniestUsername(ctx, user_id, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	response := model.CheckSubscriptionResponse{
		Subscribed: userIsSubscribed,
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Subscribers
// @Description Get Subscribers of Moniest
// @Security bearerAuth
// @Tags Moniest
// @Accept json
// @Produce json
// @Param username path string true "moniest username"
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.User ""
// @Failure 404 {object} clientError.ErrorResponse "moniest not found with this username"
// @Failure 406 {object} clientError.ErrorResponse "invalid query params"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/subscribers [get]
func (server *Server) getSubscribers(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	var req model.PaginationRequest = model.PaginationRequest{
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_GetSubscriber_InvalidParam))
		return
	}

	// STEP: make limit & offset safe [arrange min-max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)

	// STEP: get moniest [+ check there is a moniest w/ this username]
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get subscribers
	subscribers, err := server.service.GetSubscribers(ctx, moniest.MoniestID, req.Limit, req.Offset)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, subscribers)
}
