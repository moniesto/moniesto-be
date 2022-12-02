package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// GetUserByUsername gets user data
// PRIMARY TODO: update user db requests with moniest db requests
func (server *Server) GetUserByUsername(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: get own username from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	own_username := authPayload.User.Username

	// STEP: get user by username [if own user, additional +email field]
	var user interface{}
	var err error

	if username == own_username {
		user, err = server.service.GetOwnUserByUsername(ctx, username)
	} else {
		user, err = server.service.GetUserByUsername(ctx, username)
	}
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
