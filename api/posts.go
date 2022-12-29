package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// TODO: complete endpoint
func (server *Server) CreatePost(ctx *gin.Context) {
	var req model.CreatePostRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidBody))
		return
	}

}
