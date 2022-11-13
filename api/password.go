package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (server *Server) ChangePassword(ctx *gin.Context) {

	validAuth := ctx.MustGet(authorizationPayloadValidityKey).(bool)

	if validAuth {
		authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

		server.changeLoggedInUserPassword(ctx, authPayload.User.ID)
	} else {
		server.changeLoggedOutUserPassword(ctx)
	}
}

func (server *Server) changeLoggedOutUserPassword(ctx *gin.Context) {
	var req model.ChangePasswordRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: choose which kind of request is it [send email OR verify token&change password]
	if req.Token == "" && req.NewPassword == "" && req.Email != "" {
		server.sendForgetPasswordEmail(ctx, &req)
	} else if req.Token != "" && req.NewPassword != "" && req.Email == "" {
		server.verifyToken(ctx, &req)
	} else {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}
}
func (server *Server) sendForgetPasswordEmail(ctx *gin.Context, req *model.ChangePasswordRequest) {
	/*
		3-create password_reset_token object, save to db
		4-create email content, send email
		5-return 200 OK
	*/

	// STEP: check it is in the system -> if not dont send any email and return 200 OK
	validEmail, err := server.service.CheckEmailExistidy(ctx, req.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusOK) // send success case to client in email is not exist case too (security)
		// ctx.JSON(clientError.ParseError(err)) // send exact error on the client
		return
	}

	// STEP: create password_reset_token in DB
	password_reset_token, err := server.service.CreatePasswordResetToken(ctx, validEmail, server.config.PasswordResetTokenDuration)
	_ = password_reset_token
	fmt.Println("sendForgetPasswordEmail")
}

func (server *Server) verifyToken(ctx *gin.Context, req *model.ChangePasswordRequest) {
	fmt.Println("verifyToken")
}

func (server *Server) changeLoggedInUserPassword(ctx *gin.Context, user_id string) {
	var req model.ChangePasswordRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: check old password is in valid form
	err := validation.Password(req.OldPassword)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidOldPassword))
		return
	}

	// STEP: check new password is in valid form
	err = validation.Password(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidNewPassword))
		return
	}

	// STEP: check old password is correct
	err = server.service.CheckPassword(ctx, user_id, req.OldPassword)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	// STEP: update password with new one
	err = server.service.UpdatePassword(ctx, user_id, req.NewPassword)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
