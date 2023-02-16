package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (server *Server) getContentPosts(ctx *gin.Context) {
	var req model.GetContentPostRequest = model.GetContentPostRequest{
		Subscribed: true,
		Limit:      util.DEFAULT_POST_LIMIT,
		Offset:     0,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Content_GetPosts_InvalidParam))
		return
	}

	// STEP: get content posts
	server.service.GetContentPosts(ctx, req.Subscribed, req.Limit, req.Offset)

}
