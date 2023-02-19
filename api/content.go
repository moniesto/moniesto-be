package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
)

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

	// STEP: get content posts
	posts, err := server.service.GetContentPosts(ctx, user_id, req.Subscribed, req.Active, req.Limit, req.Offset)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
