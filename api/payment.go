package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/system"
)

func (server *Server) TriggerBinanceTransactionWebhook(ctx *gin.Context) {
	var req binance.WebhookRequest

	system.Log("binance transaction webhook trigger")

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		system.LogError("webhook body bind error", err.Error())
		return
	}

	// STEP: convert webhook data from str to struct
	webhookDataStr := req.WebhookDataStr
	webhookDataStr = strings.ReplaceAll(webhookDataStr, "\\", "")

	webhookData := binance.WebhookData{}
	err := json.Unmarshal([]byte(webhookDataStr), &webhookData)
	if err != nil {
		system.LogError("webhook data bind error", err.Error())
	}

	system.Log("--webhook: transaction", webhookData.MerchantTradeNo)

	// STEP: check payment transaction status
	_, err = server.service.CheckPaymentTransactionStatus(ctx, webhookData.MerchantTradeNo)
	if err != nil {
		system.LogError("webhook check transaction status error", err.Error())

		// STEP: stop sending webhook if user is already subscribed [subscription handled by other checker]
		if clientError.ParseErrorCode(err) == clientError.Moniest_Subscribe_AlreadySubscribed {
			ctx.JSON(http.StatusOK, binance.WebhookResponseSuccess())
			return
		}

		return
	}

	system.Log("Webhook Transaction - COMPLETED")
	ctx.JSON(http.StatusOK, binance.WebhookResponseSuccess())
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
		system.LogError("check transaction error", err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
