package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) RefundToUser(ctx *gin.Context, transactionID, moniestID, userID string) error {

	params := db.GetBinancePayoutHistoriesParams{
		TransactionID: transactionID,
		UserID:        userID,
		MoniestID:     moniestID,
	}

	binancePayoutHistories, err := service.Store.GetBinancePayoutHistories(ctx, params)
	if err != nil {
		return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Unsubscribe_ServerErrorGetPayoutHistory)
	}

	if len(binancePayoutHistories) == 0 {
		service.sendUnsubscribeEmail(ctx, transactionID, userID, moniestID, 0, 0) // 0 as remainingMonth
		return nil
	}

	// STEP: pop first element [do not refund current month]
	binancePayoutHistories = binancePayoutHistories[1:]

	if len(binancePayoutHistories) == 0 {
		service.sendUnsubscribeEmail(ctx, transactionID, userID, moniestID, 0, 0) // 0 as remainingMonth
		return nil
	}

	// STEP: calculate amount
	var amount float64 = 0
	for _, payoutHistory := range binancePayoutHistories {
		amount += payoutHistory.Amount
	}

	// STEP: refund
	payer_id := binancePayoutHistories[0].PayerID
	_, _, err = binance.CreateTransfer(
		service.config,
		amount,
		service.config.OperationFeePercentage,
		binance.BINANCE_TRANSFER_TYPE_OTHERS,
		binance.BINANCE_RECEIVE_TYPE_PAY_ID,
		payer_id,
		binance.BINANCE_TRANSFER_REMARK_REFUND,
	)
	if err != nil {
		errMsg := err.Error()
		// STEP: update status of payout histories as refund_fail
		service.UpdateBinancePayoutHistoriesRefund(ctx, binancePayoutHistories, db.BinancePayoutStatusRefundFail, &errMsg)

		return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Unsubscribe_ServerErrorRefund)
	}

	// STEP: update status of payout histories as refund
	service.UpdateBinancePayoutHistoriesRefund(ctx, binancePayoutHistories, db.BinancePayoutStatusRefund, nil)

	service.sendUnsubscribeEmail(ctx, transactionID, userID, moniestID, len(binancePayoutHistories), amount)

	return nil
}

func (service *Service) sendUnsubscribeEmail(ctx *gin.Context, transactionID, userID, moniestID string, remainingMonth int, amount float64) {
	moniest, err := service.GetMoniestByMoniestID(ctx, moniestID)
	if err != nil {
		system.LogError("sending unsubscribe/refund email - getting moniest error", err.Error())
	}

	user, err := service.GetOwnUserByID(ctx, userID)
	if err != nil {
		system.LogError("sending unsubscribe/refund email - getting user error", err.Error())
	}

	binanceTransaction, err := service.Store.GetBinancePaymentTransaction(ctx, transactionID)
	if err != nil {
		system.LogError("sending unsubscribe/refund email - getting binance transaction error", err.Error())
	}

	if err == nil {
		go mailing.SendUnsubscribeEmail(user.Email, service.config, user.Fullname, moniest.Fullname, moniest.Username, binanceTransaction.PayerID.String, util.Now(), binanceTransaction.MoniestFee, service.config.OperationFeePercentage, remainingMonth, amount, user.Language)
	}
}

func (service *Service) UpdateBinancePayoutHistoriesRefund(ctx *gin.Context, payoutHistories []db.GetBinancePayoutHistoriesRow, status db.BinancePayoutStatus, failureMessage *string) {
	for _, payoutHistory := range payoutHistories {
		err := service.UpdateBinancePayoutHistoryRefund(ctx, payoutHistory.ID, status, failureMessage)
		if err != nil {
			system.LogError("error while updating binance payout history [refund]", err.Error(), "| payout id:", payoutHistory.ID)
		}
	}
}

func (service *Service) UpdateBinancePayoutHistoryRefund(ctx *gin.Context, id string, status db.BinancePayoutStatus, failureMessage *string) error {
	params := db.UpdateBinancePayoutHistoryRefundParams{
		ID:     id,
		Status: status,
	}

	if failureMessage != nil {
		params.FailureMessage = sql.NullString{
			Valid:  true,
			String: *failureMessage,
		}
	}

	err := service.Store.UpdateBinancePayoutHistoryRefund(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
