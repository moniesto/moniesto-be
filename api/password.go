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

}

func sendForgetPasswordEmail() {

	fmt.Println("sendForgetPasswordEmail")
}

func verifyToken() {
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
