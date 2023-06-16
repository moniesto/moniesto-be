package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) CreateBinancePaymentTransaction(ctx *gin.Context, req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow, userID string) (db.BinancePaymentTransaction, error) {

	// STEP: create order in binance and get payment links
	product_name := getProductName(req, moniest)
	amount := getAmount(req, moniest)
	transactionID := core.CreatePlainID()
	webhookURL := createWebhookURL(ctx, transactionID)
	req.ReturnURL, req.CancelURL = updateNavigateURLs(transactionID, req.ReturnURL, req.CancelURL) // add transactionID to urls

	orderData, err := binance.CreateOrder(ctx, service.config, transactionID, amount, product_name, req.ReturnURL, req.CancelURL, webhookURL)
	if err != nil {
		systemError.Log("create order error", err.Error())
		return db.BinancePaymentTransaction{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorCreateBinanceOrder)
	}

	// STEP: add payment transactions to db
	paymentTransactionParams := db.CreateBinancePaymentTransactionsParams{
		ID:            transactionID,
		QrcodeLink:    orderData.QrcodeLink,
		CheckoutLink:  orderData.CheckoutUrl,
		DeepLink:      orderData.Deeplink,
		UniversalLink: orderData.UniversalUrl,
		Status:        db.BinancePaymentStatusPending,
		UserID:        userID,
		MoniestID:     moniest.MoniestID,
		DateType:      db.BinancePaymentDateTypeMONTH,
		DateValue:     int32(req.NumberOfMonths),
		MoniestFee:    moniest.Fee,
		Amount:        amount,
		WebhookUrl:    webhookURL,
	}

	binancePaymentTransaction, err := service.Store.CreateBinancePaymentTransactions(ctx, paymentTransactionParams)
	if err != nil {
		systemError.Log("create order on db error", err.Error())
		return db.BinancePaymentTransaction{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorCreateBinanceTransaction)
	}

	return binancePaymentTransaction, nil
}

func getProductName(req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow) string {
	return fmt.Sprintf("You are subscribing to %s %s at a monthly fee of $%.2f for %d months.", moniest.Name, moniest.Surname, util.RoundAmountDown(moniest.Fee), req.NumberOfMonths)
}

func getAmount(req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow) float64 {
	return util.RoundAmountUp(float64(req.NumberOfMonths) * moniest.Fee)
}

func updateNavigateURLs(transactionID, returnURL, cancelURL string) (string, string) {

	// remove suffix /
	returnURL = strings.TrimSuffix(returnURL, "/")
	cancelURL = strings.TrimSuffix(cancelURL, "/")

	return returnURL + "?transactionID=" + transactionID, // return url with transactionID
		cancelURL + "?transactionID=" + transactionID // cancel url with transactionID
}

func createWebhookURL(ctx *gin.Context, transactionID string) string {

	return "https://api.moniesto.com" + "/webhooks/binance/transactions/" + transactionID
	// TODO: make it host
	// return ctx.Request.Host + "/webhooks/binance/transactions/" + transactionID
}
