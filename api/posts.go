package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/mailing"
)

// @Summary Create Post
// @Description Create Crypto Post
// @Security bearerAuth
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostBody body model.CreatePostRequest true "`description` is optional"
// @Success 200 {object} model.CreatePostResponse
// @Failure 400 {object} clientError.ErrorResponse "user is not moniest"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/posts [post]
func (server *Server) createPost(ctx *gin.Context) {
	var req model.CreatePostRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidBody))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check user is moniest
	userIsMoniest, err := server.service.CheckUserIsMoniestByUserID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}
	if !userIsMoniest {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, clientError.GetError(clientError.General_UserNotMoniest))
		return
	}

	// STEP: user is moniest => get moniest data
	moniest, err := server.service.GetMoniestByUserID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get currency
	currency, err := server.service.GetCurrency(req.Currency, req.MarketType)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create post
	post, err := server.service.CreatePost(req, currency, moniest.MoniestID, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	var response model.CreatePostResponse

	if req.Description != "" {
		// STEP: create post description
		description, err := server.service.CreatePostDescription(post.ID, req.Description, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(clientError.ParseError(err))
			return
		}

		response = model.NewCreatePostResponse(post, description)
	} else {
		response = model.NewCreatePostResponse(post, db.PostCryptoDescription{})
	}

	ctx.JSON(http.StatusOK, response)

	// STEP: send email to all subscribers
	subscribers, _ := server.service.GetSubscribersBriefs(ctx, moniest.MoniestID)
	for _, subscriber := range subscribers {
		go mailing.SendNewPostEmail(subscriber.Email, server.config, subscriber.Fullname, moniest.Fullname, moniest.Username, post.Currency, subscriber.Language)
	}
}

// @Summary Get Moniest Posts
// @Description Get Posts of Moniest [active or all]
// @Security bearerAuth
// @Tags Post
// @Accept json
// @Produce json
// @Param username path string true "moniest username"
// @Param active query bool false "default: false, true: only live(active), false: all posts"
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.GetContentPostResponse
// @Failure 403 {object} clientError.ErrorResponse "forbidden access (when not subscribed, but asks for active posts)"
// @Failure 404 {object} clientError.ErrorResponse "no moniest with this username"
// @Failure 406 {object} clientError.ErrorResponse "invalid params"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/:username/posts [get]
func (server *Server) getMoniestPosts(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	var req model.GetMoniestPostsRequest = model.GetMoniestPostsRequest{
		Active: false,
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_GetMoniestPosts_InvalidParam))
		return
	}

	// STEP: make limit & offset safe [arrange min-max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)

	// STEP: check "username" is a real moniest
	moniest, err := server.service.GetMoniestByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: asks for own posts
	if moniest.ID == user_id {

		posts, err := server.service.GetOwnPosts(ctx, username, req.Active, req.Limit, req.Offset)
		if err != nil {
			ctx.AbortWithStatusJSON(clientError.ParseError(err))
			return
		}

		ctx.JSON(http.StatusOK, posts)

	} else { // STEP: not own posts
		// STEP: get user subscription status -> subscribed or not
		userIsSubscribed, err := server.service.CheckUserSubscriptionByMoniestUsername(ctx, user_id, username)
		if err != nil {
			ctx.AbortWithStatusJSON(clientError.ParseError(err))
			return
		}

		// STEP: get posts
		posts, err := server.service.GetMoniestPosts(ctx, username, userIsSubscribed, req.Active, req.Limit, req.Offset)
		if err != nil {
			ctx.AbortWithStatusJSON(clientError.ParseError(err))
			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}

// @Summary Calculate PNL and ROI
// @Description Calculate PNL and ROI based on the start price, target price, leverage, direction
// @Security bearerAuth
// @Tags Post
// @Accept json
// @Produce json
// @Param CalculatePnlRoiBody body model.CalculatePnlRoiRequest true "all required"
// @Success 200 {object} model.CalculatePnlRoiResponse
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/posts/calculate-pnl-roi [post]
func (server *Server) calculatePnlRoi(ctx *gin.Context) {
	var req model.CalculatePnlRoiRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidBody))
		return
	}

	response, err := server.service.CalculatePnlRoi(req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
