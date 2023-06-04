package api

import (
	"fmt"

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

func (server *Server) createOrder(ctx *gin.Context) {

	payment.CreateOrder(server.config)

}

func (server *Server) webhook(ctx *gin.Context) {

	jsonData, err := ctx.GetRawData()

	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Println("WEBHOOK-DATA", string(jsonData))

}
