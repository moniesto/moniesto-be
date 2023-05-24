package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/util/payment"
)

func (server *Server) addConnectedAccount(ctx *gin.Context) {

	payment.CreateConnectedAccount(server.config.StripeSecretKey, "parvvazov@gmail.com", "Parvin", "Eyvazov", "tok_1N3468A5XzcV9fbo1Y0V1jGC")
}

func (server *Server) createAccountLink(ctx *gin.Context) {

	payment.GetAccountLink(server.config.StripeSecretKey)
}

func (server *Server) deleteConnectedAccount(ctx *gin.Context) {

	// STEP: get username from param
	acc_id := ctx.Param("acc_id")

	payment.DeleteConnectedAccount(server.config.StripeSecretKey, acc_id)
}
