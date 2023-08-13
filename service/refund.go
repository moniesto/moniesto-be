package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/systemError"
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
		// TODO: send refund email [about nothing to refund]
		return nil
	}

	// STEP: pop first element [do not refund current month]
	binancePayoutHistories = binancePayoutHistories[1:]

	if len(binancePayoutHistories) == 0 {
		// TODO: send refund email [about nothing to refund]
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

	// TODO: send refund email
	return nil
}

func (service *Service) UpdateBinancePayoutHistoriesRefund(ctx *gin.Context, payoutHistories []db.GetBinancePayoutHistoriesRow, status db.BinancePayoutStatus, failureMessage *string) {
	for _, payoutHistory := range payoutHistories {
		err := service.UpdateBinancePayoutHistoryRefund(ctx, payoutHistory.ID, status, failureMessage)
		if err != nil {
			systemError.Log("error while updating binance payout history [refund]", err.Error(), "| payout id:", payoutHistory.ID)
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
