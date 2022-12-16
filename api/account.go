package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Login
// @Description Login with [email & password] OR [username & password]
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginBody body model.LoginRequest true "identifier can be email OR username"
// @Success 200 {object} model.LoginResponse
// @Failure 403 {object} clientError.ErrorResponse "wrong password"
// @Failure 404 {object} clientError.ErrorResponse "email OR username not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req model.LoginRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_Login_InvalidBody))
		return
	}

	// STEP: get own user [+ checking password]
	user, err := server.service.GetOwnUser(ctx, req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
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
		// TODO: add server error
		ctx.JSON(http.StatusInternalServerError, clientError.GetError(clientError.Account_Login_ServerErrorToken))
		return
	}

	// STEP: update login stat
	server.service.UpdateLoginStats(ctx, user.ID)

	// STEP: create response object
	rsp := model.NewLoginResponse(accessToken, user)

	ctx.JSON(http.StatusOK, rsp)
}

// @Summary Register
// @Description Register as user
// @Tags Auth
// @Accept json
// @Produce json
// @Param "Register body" body model.RegisterRequest true " "
// @Success 200 {object} model.RegisterResponse
// @Failure 403 {object} clientError.ErrorResponse "wrong password"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/register [post]
func (server *Server) registerUser(ctx *gin.Context) {
	var req model.RegisterRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_Register_InvalidBody))
		return
	}

	// STEP: create user
	createdUser, err := server.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
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

// @Summary Check username
// @Description Check username is valid of not
// @Tags Account
// @Accept json
// @Produce json
// @Param username path int true "username"
// @Success 200 {object} model.CheckUsernameResponse
// @Failure 406 {object} clientError.ErrorResponse "invalid username"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /usernames/:username/check [get]
func (server *Server) checkUsername(ctx *gin.Context) {
	// STEP: get username from param
	username := ctx.Param("username")

	// STEP: check username validity
	validity, err := server.service.CheckUsername(ctx, username)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: create response object
	rsp := model.NewCheckUsernameResponse(validity)

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) updateProfile(ctx *gin.Context) {

	fmt.Println("updateProfile")

}
