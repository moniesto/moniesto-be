package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	// amount := core.GetTotalAmount(req.NumberOfMonths, moniest.Fee)
	amount := 0.00000001 // TODO: update to real amount
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

// run CheckPaymentTransactionStatus function X times in Y minute
func (service *Service) SetPaymentTransactionCheckerJob(ctx *gin.Context, minute, times int, transactionID string) {

	seconds := minute * 60

	interval := seconds / times

	count := times

	for count > 0 {

		time.Sleep(time.Duration(interval) * time.Second)
		count += -1
		fmt.Println("CHECKER RUN")

		latestTransactionStatus, err := service.CheckPaymentTransactionStatus(ctx, transactionID)

		// if any error occurs, skip
		if err != nil {
			_, msg := clientError.ParseError(err)
			systemError.Log("error while checking payment transaction status", msg)
		} else {
			if latestTransactionStatus != string(db.BinancePaymentStatusPending) {
				count = -1 // stop the loop
			}
		}
	}
}

// CheckPaymentTransactionStatus check binance transaction status
//   - on succuss:
//     -- update transaction status to success
//     -- subscribe to moniest
//     -- create payout operation
//   - on failure:
//     -- update transaction status to failure
//   - on pending:
//     -- do nothing
func (service *Service) CheckPaymentTransactionStatus(ctx *gin.Context, transactionID string) (string, error) {

	// STEP: get transaction data
	binancePaymentTransaction, err := service.Store.GetBinancePaymentTransaction(ctx, transactionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", clientError.CreateError(http.StatusNotFound, clientError.Payment_CheckBinanceTransaction_TransactionIDNotFound)
		}

		systemError.Log("server error on get binance payment transaction", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorGetTransaction)
	}

	// STEP: check it is still pending
	if binancePaymentTransaction.Status != db.BinancePaymentStatusPending {
		return string(binancePaymentTransaction.Status), nil
	}

	orderData, err := binance.QueryOrder(ctx, service.config, transactionID)
	if err != nil {
		systemError.Log("server error on query order", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorQueryTransaction)
	}

	// STEP: still pending state
	if orderData.Status == binance.QUERY_ORDER_STATUS_INITIAL {
		return string(binancePaymentTransaction.Status), nil // TODO: update return -> nothing to do
	}

	// STEP: payment successful case
	if orderData.Status == binance.QUERY_ORDER_STATUS_PAID {

		// STEP: update transaction status
		param := db.UpdateBinancePaymentTransactionStatusParams{
			ID:      transactionID,
			Status:  db.BinancePaymentStatusSuccess,
			PayerID: sql.NullString{Valid: true, String: orderData.PaymentInfo.PayerId},
		}

		updatedBinancePaymentTransaction, err := service.Store.UpdateBinancePaymentTransactionStatus(ctx, param)
		if err != nil {
			systemError.Log("server error on update binance payment transaction status", err.Error())
			return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorUpdateStatusSuccess)
		}

		// STEP: subscribe to moniest
		err = service.SubscribeMoniest(ctx,
			updatedBinancePaymentTransaction.MoniestID,
			updatedBinancePaymentTransaction.UserID,
			updatedBinancePaymentTransaction.ID,
			int(updatedBinancePaymentTransaction.DateValue),
		)
		if err != nil {
			return "", err
		}

		// TODO: create payout history (or maybe do it in subscribe function)
		service.CreateBinancePayoutHistories(ctx, db.CreateBinancePayoutHistoryParams{
			TransactionID: updatedBinancePaymentTransaction.ID,
			UserID:        updatedBinancePaymentTransaction.UserID,
			MoniestID:     updatedBinancePaymentTransaction.MoniestID,
			PayerID:       orderData.PaymentInfo.PayerId,
			TotalAmount:   updatedBinancePaymentTransaction.Amount,
			Amount:        updatedBinancePaymentTransaction.MoniestFee,
			DateType:      updatedBinancePaymentTransaction.DateType,
			DateValue:     updatedBinancePaymentTransaction.DateValue,
			Status:        db.BinancePayoutStatusPending,
		})

		return string(updatedBinancePaymentTransaction.Status), nil
	}

	// STEP: payment failed cases [rest of them]
	param := db.UpdateBinancePaymentTransactionStatusParams{
		ID:     transactionID,
		Status: db.BinancePaymentStatusFail,
	}
	updatedBinancePaymentTransaction, err := service.Store.UpdateBinancePaymentTransactionStatus(ctx, param)
	if err != nil {
		systemError.Log("server error on update binance payment transaction status", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorUpdateStatusFail)
	}

	return string(updatedBinancePaymentTransaction.Status), nil
}

func (service *Service) CheckPendingPaymentTransaction(ctx *gin.Context, moniestUsername, user_id string) (bool, error) {

	// STEP: check payment transaction is in pending status
	param := db.CheckPendingBinancePaymentTransactionByMoniestUsernameParams{
		UserID:   user_id,
		Username: moniestUsername,
	}

	transactionIsPending, err := service.Store.CheckPendingBinancePaymentTransactionByMoniestUsername(ctx, param)
	if err != nil {
		systemError.Log("server error on check pending binance payment transaction by moniest username", err.Error())
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_SubscribeCheck_ServerErrorCheck)
	}

	return transactionIsPending, nil
}

// HELPER functions
func getProductName(req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow) string {
	return fmt.Sprintf("You are subscribing to %s at a monthly fee of $%.2f for %d months.", moniest.Fullname, util.RoundAmountDown(moniest.Fee), req.NumberOfMonths)
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
