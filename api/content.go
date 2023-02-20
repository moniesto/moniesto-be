package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Get Content Post
// @Description Get Posts for timeline & explore
// @Security bearerAuth
// @Tags Content
// @Accept json
// @Produce json
// @Param subscribed query bool true "default: true"
// @Param active query bool true "default: true"
// @Param limit query int true "default: 10 & max: 50"
// @Param offset query int true "default: 0"
// @Success 200 {object} []model.GetContentPostResponse
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /content/posts [get]
func (server *Server) getContentPosts(ctx *gin.Context) {
	var req model.GetContentPostRequest = model.GetContentPostRequest{
		Subscribed: true,
		Active:     true,
		Limit:      util.DEFAULT_POST_LIMIT,
		Offset:     0,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetPosts_InvalidParam))
		return
	}

	// get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check max limit
	if req.Limit > util.MAX_POST_LIMIT {
		req.Limit = util.MAX_POST_LIMIT
	}

	// STEP: get content posts
	posts, err := server.service.GetContentPosts(ctx, user_id, req.Subscribed, req.Active, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (server *Server) getContentMoniests(ctx *gin.Context) {
	var req model.GetContentMoniestRequest = model.GetContentMoniestRequest{
		Subscribed: true,
		Limit:      util.DEFAULT_MONIEST_LIMIT,
		Offset:     0,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetPosts_InvalidParam))
		return
	}

	// get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check max limit
	if req.Limit > util.MAX_MONIEST_LIMIT {
		req.Limit = util.MAX_MONIEST_LIMIT
	}

	// STEP: get content moniests
	/* moniests, err := */
	server.service.GetContentMoniests(ctx, user_id, req.Subscribed, req.Limit, req.Offset)

}
