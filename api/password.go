package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/systemError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (server *Server) ChangePassword(ctx *gin.Context) {

	validAuth := ctx.MustGet(authorizationPayloadValidityKey).(bool)

	if validAuth {
		authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

		server.changeLoggedUserPassword(ctx, authPayload.User.ID)
	} else {
		sendForgetPasswordEmail()
		verifyToken()
	}
}

func sendForgetPasswordEmail() {
	fmt.Println("sendForgetPasswordEmail")
}

func verifyToken() {
	fmt.Println("verifyToken")
}

func (server *Server) changeLoggedUserPassword(ctx *gin.Context, user_id string) {
	var req model.ChangePasswordRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_ChangePassword"]())
		return
	}

	// STEP: check old password is in valid form
	err := validation.Password(req.OldPassword)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_ChangePassword"]("old password is not in valid form"))
		return
	}

	// STEP: check new password is in valid form
	err = validation.Password(req.NewPassword)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_ChangePassword"]("new password is not in valid form"))
		return
	}

	// STEP: check old password is correct
	err = server.service.CheckPassword(ctx, user_id, req.OldPassword)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_ChangePassword"](err.Error()))
		return
	}

	// STEP: update password with new one
	err = server.service.UpdatePassword(ctx, user_id, req.NewPassword)
	if err != nil {
		ctx.JSON(systemError.Messages["Invalid_RequestBody_ChangePassword"](err.Error()))
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
