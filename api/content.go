package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Get Content Posts
// @Description Get Posts for timeline & explore
// @Security bearerAuth
// @Tags Content
// @Accept json
// @Produce json
// @Param subscribed query bool false "default: true"
// @Param active query bool false "default: true"
// @Param sortBy query string false "options: [score | created_at] default: score, only affect when subscription & active = false"
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.GetContentPostResponse
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /content/posts [get]
func (server *Server) getContentPosts(ctx *gin.Context) {
	var req model.GetContentPostRequest = model.GetContentPostRequest{
		Subscribed: true,
		Active:     true,
		SortBy:     "score",
		Limit:      util.DEFAULT_LIMIT,
		Offset:     util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetPosts_InvalidParam))
		return
	}

	// get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: make limit & offset safe [arrange min-max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)
	req.SortBy = util.SafePostSortBy(req.SortBy)

	// STEP: get content posts
	posts, err := server.service.GetContentPosts(ctx, user_id, req.Subscribed, req.Active, req.SortBy, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

// @Summary Get Content Moniests
// @Description Get All moniests
// @Security bearerAuth
// @Tags Content
// @Accept json
// @Produce json
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.User
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /content/moniests [get]
func (server *Server) getContentMoniests(ctx *gin.Context) {
	var req model.GetContentMoniestRequest = model.GetContentMoniestRequest{
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetMoniests_InvalidParam))
		return
	}

	// get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: make limit & offset safe [arrange min-max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)

	// STEP: get content moniests
	moniests, err := server.service.GetContentMoniests(ctx, user_id, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, moniests)
}

// @Summary Search Moniest
// @Description Search Moniest by their name & surname & username
// @Security bearerAuth
// @Tags Content
// @Accept json
// @Produce json
// @Param searchText query string true "length min:1 max:30"
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.User
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /content/moniests/search [get]
func (server *Server) searchMoniest(ctx *gin.Context) {
	var req model.SearchMoniestRequest = model.SearchMoniestRequest{
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetMoniests_InvalidParam))
		return
	}

	// STEP: safe limit & offset & searchtext [arrange min - max]
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)
	req.SearchText = util.SafeSearchText(req.SearchText)

	// STEP: search moniest
	moniests, err := server.service.SearchMoniest(ctx, req.SearchText, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, moniests)
}
