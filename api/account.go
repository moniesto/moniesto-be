package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/systemError"
)

type loginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password" binding:"required,min=6"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.service.GetOwnUser(ctx, req.Identifier, req.Password)
	if err != nil {
		// TODO: update with systemError
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		// TODO: update with systemError
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginResponse{
		Token: accessToken,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) registerUser(ctx *gin.Context) {
	var req model.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_Register"]())
		return
	}

	// save user to db
	createdUser, err := server.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_Register"](err.Error()))
		return
	}

	// create token
	accessToken, err := server.tokenMaker.CreateToken(req.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := model.RegisterResponse{
		Token: accessToken,
	}

	_ = rsp

	fmt.Println("createdUser", createdUser)
	ctx.JSON(http.StatusOK, createdUser)
}
