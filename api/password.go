package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/validation"
)

// @Summary Change Password
// @Description Authenticated user password change
// @Security bearerAuth
// @Tags Password
// @Accept json
// @Produce json
// @Param ChangePasswordBody body model.ChangePasswordRequest true " "
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "wrong password"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/password [put]
func (server *Server) changePassword(ctx *gin.Context) {
	var req model.ChangePasswordRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: check old password is in valid form
	err := validation.Password(req.OldPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidOldPassword))
		return
	}

	// STEP: check new password is in valid form
	err = validation.Password(req.NewPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidNewPassword))
		return
	}

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: check old password is correct
	err = server.service.CheckPassword(ctx, user_id, req.OldPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: update password with new one
	err = server.service.UpdatePassword(ctx, user_id, req.NewPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Send Reset Password Email
// @Description Unauthenticated user send reset password email
// @Tags Password
// @Accept json
// @Produce json
// @Param SendEmailBody body model.SendResetPasswordEmailRequest true " "
// @Success 202
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/password/send_email [post]
func (server *Server) sendResetPasswordEmail(ctx *gin.Context) {
	var req model.SendResetPasswordEmailRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: check email is in the system -> if not don't send any email and return 202 ACCEPTED
	validEmail, err := server.service.CheckEmailExistidy(ctx, req.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusAccepted) // send success case to client in email is not exist case too (security)
		// ctx.JSON(clientError.ParseError(err)) // send exact error on the client
		return
	}

	// STEP: create password_reset_token in DB
	name, password_reset_token, err := server.service.CreatePasswordResetToken(ctx, validEmail, server.config.PasswordResetTokenDuration)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: send password reset email
	err = mailing.SendPasswordResetEmail(validEmail, server.config, name, password_reset_token.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, clientError.GetError(clientError.Account_ChangePassowrd_SendEmail))
		return
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary Verify Token & Change Password
// @Description Unauthenticated verify token & change password
// @Tags Password
// @Accept json
// @Produce json
// @Param VerifyTokenBody body model.VerifyPasswordResetRequest true "token & new fiels are required"
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "token is expired"
// @Failure 404 {object} clientError.ErrorResponse "reset token not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body & data"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/password/change_password [post]
func (server *Server) verifyTokenChangePassword(ctx *gin.Context) {
	var req model.VerifyPasswordResetRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: validating password reset token [decode + expiry check]
	password_reset_token, err := server.service.GetPasswordResetToken(ctx, req.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: check new password is in valid form
	err = validation.Password(req.NewPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidNewPassword))
		return
	}

	// STEP: update password with new one
	err = server.service.UpdatePassword(ctx, password_reset_token.UserID, req.NewPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: delete the token that used
	err = server.service.DeletePasswordResetToken(ctx, password_reset_token.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Verify Token
// @Description Unauthenticated verify token
// @Tags Password
// @Accept json
// @Produce json
// @Param VerifyTokenBody body model.VerifyTokenRequest true "token is required"
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "token is expired"
// @Failure 404 {object} clientError.ErrorResponse "reset token not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /account/password/verify_token [post]
func (server *Server) verifyToken(ctx *gin.Context) {
	var req model.VerifyTokenRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Account_ChangePassword_InvalidBody))
		return
	}

	// STEP: validating password reset token [decode + expiry check]
	_, err := server.service.GetPasswordResetToken(ctx, req.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}
