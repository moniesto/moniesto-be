package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// GetUserByUsername gets user data
// PRIMARY TODO: update user db requests with moniest db requests (dont have idea why)
// @Summary Get User by Username
// @Description get user info with username
// @Security bearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path int true "username"
// @Success 200 {object} model.User "'email' field will be visible if user request for own account"
// @Failure 404 {object} clientError.ErrorResponse "not any user with this username"
// @Failure 406 {object} clientError.ErrorResponse "invalid username"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /users/:username [get]
func (server *Server) GetUserByUsername(ctx *gin.Context) {
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

		response = model.NewGetOwnUserResponse(user)
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
