package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/util/payment"
	"github.com/moniesto/moniesto-be/util/payment/binance"
)

func (server *Server) CheckBinancePaymentTransaction(ctx *gin.Context) {
	// STEP: get username from param
	transactionID := ctx.Param("transaction_id")

	binance.QueryOrder(ctx, server.config, transactionID)

}

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

func (server *Server) createPayout(ctx *gin.Context) {

	payment.CreatePayout(server.config)

}

func (server *Server) webhook(ctx *gin.Context) {

	jsonData, err := ctx.GetRawData()

	/*

		RESPONSE [success]:
			{"bizType":"PAY","data":"{\"merchantTradeNo\":\"755f2a5bdc42444991b08124eda15638\",
			\"productType\":\"02\",\"productName\":\"Moniest 1 - A\",\"transactTime\":1685921049737,
			\"tradeType\":\"WEB\",\"totalFee\":1.0E-7,\"currency\":\"USDT\",\"transactionId\":\"P_A1BQS87BCQ171112\",
			\"commission\":0,\"paymentInfo\":{\"payerId\":741232235,\"payMethod\":\"funding\",
			\"paymentInstructions\":[{\"currency\":\"USDT\",\"amount\":1.0E-7,\"price\":1}],
			\"channel\":\"DEFAULT\"}}","bizIdStr":"232103202548367360","bizId":232103202548367360,"bizStatus":"PAY_SUCCESS"}

	*/

	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Println("WEBHOOK-DATA", string(jsonData))

}
