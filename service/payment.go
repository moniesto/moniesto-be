package service

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/message"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) CreateBinancePaymentTransaction(ctx *gin.Context, req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow, userID string) (db.BinancePaymentTransaction, error) {

	// STEP: create order in binance and get payment links
	product_name, err := service.getProductName(ctx, req, moniest, userID)
	if err != nil {
		system.LogError("Server error on getting product name", err.Error())
		return db.BinancePaymentTransaction{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorGetProductName)
	}

	// amount := core.GetTotalAmount(req.NumberOfMonths, moniest.Fee)
	amount := 0.00000001 // TODO: update to real amount
	transactionID := core.CreatePlainID()
	webhookURL := createWebhookURL(ctx, transactionID)
	req.ReturnURL, req.CancelURL = updateNavigateURLs(transactionID, req.ReturnURL, req.CancelURL) // add transactionID to urls

	orderData, err := binance.CreateOrder(ctx, service.config, transactionID, amount, product_name, req.ReturnURL, req.CancelURL, webhookURL)
	if err != nil {
		system.LogError("create order error", err.Error())
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
		system.LogError("create order on db error", err.Error())
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
		system.Log("PING - Payment Transaction Checker")

		latestTransactionStatus, err := service.CheckPaymentTransactionStatus(ctx, transactionID)

		// if any error occurs, skip
		if err != nil {
			_, msg := clientError.ParseError(err)
			system.LogError("error while checking payment transaction status", msg)
		} else {
			if latestTransactionStatus != string(db.BinancePaymentStatusPending) {
				system.Log("PING - Payment Transaction Checker - COMPLETED")

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

		system.LogError("server error on get binance payment transaction", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorGetTransaction)
	}

	// STEP: check it is still pending
	if binancePaymentTransaction.Status != db.BinancePaymentStatusPending {
		return string(binancePaymentTransaction.Status), nil
	}

	orderData, err := binance.QueryOrder(ctx, service.config, transactionID)
	if err != nil {
		system.LogError("server error on query order", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorQueryTransaction)
	}

	// STEP: still pending state
	if orderData.Status == binance.QUERY_ORDER_STATUS_INITIAL {
		return string(binancePaymentTransaction.Status), nil
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
			system.LogError("server error on update binance payment transaction status", err.Error())
			return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorUpdateStatusSuccess)
		}

		subscriptionStartDate := util.Now()
		subscriptionEndDate := subscriptionStartDate.AddDate(0, int(updatedBinancePaymentTransaction.DateValue), 0)

		// STEP: subscribe to moniest
		err = service.SubscribeMoniest(ctx,
			updatedBinancePaymentTransaction.MoniestID,
			updatedBinancePaymentTransaction.UserID,
			updatedBinancePaymentTransaction.ID,
			subscriptionStartDate,
			subscriptionEndDate,
		)
		if err != nil {
			return "", err
		}

		// STEP: create payout history
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

		// STEP: send emails to user and moniest
		service.sendSubscribedEmails(
			ctx, updatedBinancePaymentTransaction.UserID, updatedBinancePaymentTransaction.MoniestID,
			subscriptionStartDate, subscriptionEndDate, updatedBinancePaymentTransaction.MoniestFee,
			updatedBinancePaymentTransaction.Amount, int(updatedBinancePaymentTransaction.DateValue),
		)

		return string(updatedBinancePaymentTransaction.Status), nil
	}

	// STEP: payment failed cases [rest of them]
	param := db.UpdateBinancePaymentTransactionStatusParams{
		ID:     transactionID,
		Status: db.BinancePaymentStatusFail,
	}
	updatedBinancePaymentTransaction, err := service.Store.UpdateBinancePaymentTransactionStatus(ctx, param)
	if err != nil {
		system.LogError("server error on update binance payment transaction status", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Payment_CheckBinanceTransaction_ServerErrorUpdateStatusFail)
	}

	return string(updatedBinancePaymentTransaction.Status), nil
}

func (service *Service) sendSubscribedEmails(ctx *gin.Context, userID, moniestID string, subscriptionStartDate, subscriptionEndDate time.Time, subscriptionFee, amount float64, subscriptionMonth int) {
	// STEP: fetch user and moniest infos
	moniest, err := service.GetMoniestByMoniestID(ctx, moniestID)
	if err != nil {
		system.LogError("sending subscribe email - getting moniest error", err.Error())
	}

	user, err := service.GetOwnUserByID(ctx, userID)
	if err != nil {
		system.LogError("sending subscribe email - getting user error", err.Error())
	}

	// STEP: moniest & user is valid
	if err == nil {
		// STEP: send subscribed email to moniest
		go mailing.SendSubscribedEmailMoniest(moniest.Email, service.config, moniest.Fullname, user.Fullname, user.Username, subscriptionStartDate, subscriptionEndDate, subscriptionFee, subscriptionMonth, moniest.Language)

		// STEP: send subscribed email to user
		go mailing.SendSubscribedEmailUser(user.Email, service.config, user.Fullname, moniest.Fullname, moniest.Username, subscriptionStartDate, subscriptionEndDate, subscriptionFee, amount, subscriptionMonth, user.Language)
	}
}

func (service *Service) CheckPendingPaymentTransaction(ctx *gin.Context, moniestUsername, user_id string) (bool, *time.Time, *db.CheckPendingBinancePaymentTransactionByMoniestUsernameRow, error) {

	// STEP: check payment transaction is in pending status
	param := db.CheckPendingBinancePaymentTransactionByMoniestUsernameParams{
		UserID:   user_id,
		Username: moniestUsername,
	}

	pendingTransaction, err := service.Store.CheckPendingBinancePaymentTransactionByMoniestUsername(ctx, param)
	if err != nil {
		system.LogError("server error on check pending binance payment transaction by moniest username", err.Error())
		return false, nil, nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_SubscribeCheck_ServerErrorCheck)
	}

	if len(pendingTransaction) == 0 {
		return false, nil, nil, nil
	}

	latestTransaction := pendingTransaction[0]

	timeout := (&latestTransaction.CreatedAt).Add(binance.ORDER_EXPIRE_TIME)
	timeoutPntr := &timeout

	return latestTransaction.Pending, timeoutPntr, &latestTransaction, nil
}

// HELPER functions

// getProductName returns the product name based on the language of the user
func (service *Service) getProductName(ctx *gin.Context, req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow, userID string) (string, error) {
	user, err := service.GetOwnUserByID(ctx, userID)
	if err != nil {
		system.LogError("getting product name - getting user error", err.Error())
	}

	msg, err := message.GetMessage(user.Language, message.ProductName, moniest.Fullname, util.RoundAmountDown(moniest.Fee), req.NumberOfMonths)
	if err != err {
		return "", err
	}

	return msg, nil
}

func updateNavigateURLs(transactionID, returnURL, cancelURL string) (string, string) {

	// remove suffix /
	returnURL = strings.TrimSuffix(returnURL, "/")
	cancelURL = strings.TrimSuffix(cancelURL, "/")

	return returnURL + "?transactionID=" + transactionID, // return url with transactionID
		cancelURL + "?transactionID=" + transactionID // cancel url with transactionID
}

func createWebhookURL(ctx *gin.Context, transactionID string) string {

	return "https://moniesto-test-be-1.onrender.com" + "/webhooks/binance/transactions" // + transactionID
	// TODO: make it host
	// return ctx.Request.Host + "/webhooks/binance/transactions/" + transactionID
}
