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
	"github.com/moniesto/moniesto-be/util/mailing"
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
// @Param RegisterBody body model.RegisterRequest true " "
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

// @Summary Check Username
// @Description Check username is valid of not
// @Tags Account
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} model.CheckUsernameResponse
// @Failure 406 {object} clientError.ErrorResponse "invalid username"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/usernames/:username/check [get]
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

// @Summary Send Verification Email
// @Description Email verification email sender
// @Security bearerAuth
// @Tags Account
// @Accept json
// @Produce json
// @Param SendVerificationEmailBody body model.SendVerificationEmailRequest true "redirect_url is required"
// @Success 202
// @Failure 400 {object} clientError.ErrorResponse "email already verified"
// @Failure 404 {object} clientError.ErrorResponse "user not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/email/send_verification_email [post]
func (server *Server) sendVerificationEmail(ctx *gin.Context) {
	var req model.SendVerificationEmailRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_EmailVerification_InvalidBody))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: get user
	user, err := server.service.GetOwnUserByID(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: check user email is already verified
	if user.EmailVerified {
		ctx.JSON(http.StatusBadRequest, clientError.GetError(clientError.Account_EmailVerification_AlreadyVerified))
		return
	}

	// STEP: create email verification
	email_verification_token, err := server.service.CreateEmailVerificationToken(ctx, user_id, req.RedirectURL, server.config.EmailVerificationTokenDuration)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: send verification email
	err = mailing.SendEmailVerificationEmail(user.Email, server.config, user.Name, email_verification_token.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, clientError.GetError(clientError.Account_EmailVerification_SendEmail))
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary Verify Email
// @Description Verify email by token
// @Tags Account
// @Accept json
// @Produce json
// @Param VerifyEmailBody body model.VerifyEmailRequest true "token is required"
// @Success 200 {object} model.VerifyEmailResponse
// @Failure 400 {object} clientError.ErrorResponse "already verified email"
// @Failure 403 {object} clientError.ErrorResponse "expired token"
// @Failure 404 {object} clientError.ErrorResponse "token not found | user not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & token"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/email/verify_email [post]
func (server *Server) verifyEmail(ctx *gin.Context) {
	var req model.VerifyEmailRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_EmailVerification_InvalidBody))
		return
	}

	// STEP: validating email verification token [decode + expiry check]
	email_verification_token, err := server.service.GetEmailVerificationToken(ctx, req.Token)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: get user id from email verification record
	user_id := email_verification_token.UserID

	// STEP: get user
	user, err := server.service.GetOwnUserByID(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: check user email is already verified
	if user.EmailVerified {
		ctx.JSON(http.StatusBadRequest, clientError.GetError(clientError.Account_EmailVerification_AlreadyVerified))
		return
	}

	// STEP: verify user email
	err = server.service.VerifyEmail(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: delete email verification token
	err = server.service.DeleteEmailVerificationToken(ctx, email_verification_token.Token)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	response := model.VerifyEmailResponse{
		RedirectURL: email_verification_token.RedirectUrl,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) changeUsername(ctx *gin.Context) {
	var req model.ChangeUsernameRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangeUsername_InvalidBody))
		return
	}

	// STEP: check username validity
	validity, err := server.service.CheckUsername(ctx, req.NewUsername)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: username is already registered
	if !validity {
		ctx.JSON(http.StatusForbidden, clientError.GetError(clientError.Account_ChangeUsername_RegisteredUsername))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: update username of user
	err = server.service.ChangeUsername(ctx, user_id, req.NewUsername)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: create new token
	accessToken, err := server.tokenMaker.CreateToken(token.GeneralPaylod{
		UserPayload: token.UserPayload{
			Username: req.NewUsername,
			ID:       user_id,
		},
	}, server.config.AccessTokenDuration)
	if err != nil {
		// TODO: add server error
		ctx.JSON(http.StatusInternalServerError, clientError.GetError(clientError.Account_ChangeUsername_ServerErrorToken))
		return
	}

	// STEP: send response
	response := model.ChangeUsernameResponse{
		Token: accessToken,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) updateProfile(ctx *gin.Context) {

	fmt.Println("updateProfile")

}
