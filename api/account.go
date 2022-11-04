package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) loginUser(ctx *gin.Context) {
	var req model.LoginRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_Login"]())
		return
	}

	// STEP: get own user [+ checking password]
	user, err := server.service.GetOwnUser(ctx, req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(systemError.Messages["Wrong_LoginCredentials"](err.Error()))
		return
	}

	// STEP: create token
	accessToken, err := server.tokenMaker.CreateToken(token.GeneralPaylod{
		UserPayload: token.UserPayload{
			Username: user.Username,
			ID:       user.ID,
		},
	}, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(systemError.Messages["Server_TokenCreate"]())
		return
	}

	// STEP: update login stat
	server.service.UpdateLoginStats(ctx, user.ID)

	// STEP: create response object
	rsp := model.NewLoginResponse(accessToken, user)

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) registerUser(ctx *gin.Context) {
	var req model.RegisterRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_Register"]())
		return
	}

	// STEP: create user
	createdUser, err := server.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_Register"](err.Error()))
		return
	}

	// STEP: login
	loginRequestBody := model.LoginRequest{
		Identifier: createdUser.Email,
		Password:   req.Password,
	}
	loginRequestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(loginRequestBodyBytes).Encode(loginRequestBody)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(loginRequestBodyBytes.Bytes()))

	server.loginUser(ctx)
}

func (server *Server) checkUsername(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check username validity
	validity, err := server.service.CheckUsername(ctx, username)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_Username"]())
		return
	}

	// STEP: create response object
	rsp := model.NewCheckUsernameResponse(validity)

	ctx.JSON(http.StatusOK, rsp)
}
