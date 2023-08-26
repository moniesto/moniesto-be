package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) TriggerBinanceTransactionWebhook(ctx *gin.Context) {
	var req binance.WebhookRequest

	// TODO: remove log
	systemError.LogBody("binance trigger webhook", ctx)

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		systemError.Log("webhook body bind error", err.Error())
		return
	}

	// STEP: check payment transaction status
	_, err := server.service.CheckPaymentTransactionStatus(ctx, req.WebhookData.MerchantTradeNo)
	if err != nil {
		systemError.Log("webhook check transaction status error", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, `{"returnCode":"SUCCESS","returnMessage":null}`)
}

// @Summary Check Transaction ID
// @Description Check Transaction ID
// @Security bearerAuth
// @Tags Payment
// @Accept json
// @Produce json
// @Param transaction_id path string true "transaction ID"
// @Success 200
// @Failure 404 {object} clientError.ErrorResponse "transactionID not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /payment/binance/transactions/check/:transaction_id [post]
func (server *Server) CheckBinancePaymentTransaction(ctx *gin.Context) {
	// STEP: get username from param
	transactionID := ctx.Param("transaction_id")

	// STEP: check payment transaction status
	_, err := server.service.CheckPaymentTransactionStatus(ctx, transactionID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		systemError.Log("check transaction error", err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
