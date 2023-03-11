package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// PRIMARY TODO: update user db requests with moniest db requests (dont have idea why)
// @Summary Get User by Username
// @Description get user info with username
// @Security bearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} model.User "'email' field will be visible if user request for own account"
// @Failure 404 {object} clientError.ErrorResponse "not any user with this username"
// @Failure 406 {object} clientError.ErrorResponse "invalid username"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /users/:username [get]
func (server *Server) getUserByUsername(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: get own username from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	own_username := authPayload.User.Username
	own_user_id := authPayload.User.ID

	// STEP: get user by username [if own user, additional +email field]
	var response interface{}

	if username == own_username {
		user, err := server.service.GetOwnUserByUsername(ctx, username)

		if err != nil {
			ctx.JSON(clientError.ParseError(err))
			return
		}

		// TODO: find a better solution for this problem (update token when username changes)
		// if user changed username, but it is not updated on TOKEN
		if user.ID != own_user_id {
			ctx.JSON(http.StatusUnauthorized, clientError.GetError(clientError.Account_Authorization_InvalidToken))
			return
		}

		response = model.NewGetOwnUserResponseByUsername(user)
	} else {
		user, err := server.service.GetUserByUsername(ctx, username)

		if err != nil {
			ctx.JSON(clientError.ParseError(err))
			return
		}

		response = model.NewGetUserResponse(user)
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Update User Profile
// @Description Update user profile [name, surname, location, profile photo, background photo]
// @Security bearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param UpdateUserBody body model.UpdateUserProfileRequest true "all fields are optional"
// @Success 200 {object} model.OwnUser
// @Failure 404 {object} clientError.ErrorResponse "user not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /users/profile [patch]
func (server *Server) updateUserProfile(ctx *gin.Context) {
	var req model.UpdateUserProfileRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_UpdateUserProfile_InvalidBody))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: update profile
	err := server.service.UpdateUserProfile(ctx, user_id, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: update profile photo
	err = server.service.UpdateProfilePhoto(ctx, user_id, req.ProfilePhoto)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: update background photo
	err = server.service.UpdateBackgroundPhoto(ctx, user_id, req.BackgroundPhoto)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get latest form of own user, and send it as response
	user, err := server.service.GetOwnUserByID(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	response := model.NewGetOwnUserResponseByID(user)

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Subscriptions
// @Description Get subscribed Users
// @Security bearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.User
// @Failure 404 {object} clientError.ErrorResponse "user not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /users/:username/subscriptions [get]
func (server *Server) getSubscriptions(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	var req model.PaginationRequest = model.PaginationRequest{
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.User_GetSubscriptions_InvalidParam))
		return
	}

	// STEP: make limit & offset safe [arrange min-max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)

	// STEP: get user data of username [+check if user exists]
	user, err := server.service.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get subscriptions
	users, err := server.service.GetSubscriptions(ctx, user.ID, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary Get Stats
// @Description Get User & Moniest Stats
// @Security bearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} model.MoniestStatResponse "when user is not moniest, it will NOT return `subscriber_count` and `post_count`"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /users/:username/summary-stats [get]
func (server *Server) getUserStats(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check user is moniest
	userIsMoniest, err := server.service.CheckUserIsMoniestByUsername(ctx, username)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get moniest stats
	if userIsMoniest {
		stats, err := server.service.GetMoniestStats(ctx, username)
		if err != nil {
			ctx.JSON(clientError.ParseError(err))
			return
		}

		ctx.JSON(http.StatusOK, stats)

	} else { // STEP: get user stats
		stats, err := server.service.GetUserStats(ctx, username)
		if err != nil {
			ctx.JSON(clientError.ParseError(err))
			return
		}

		ctx.JSON(http.StatusOK, stats)
	}
}
