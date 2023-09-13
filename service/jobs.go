package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/analyzer"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) GetAllActivePosts() ([]db.GetAllActivePostsRow, error) {

	ctx := context.Background()

	posts, err := service.Store.GetAllActivePosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (service *Service) UpdatePostStatus(activePost db.GetAllActivePostsRow) (db.PostCryptoStatus, error) {
	ctx := context.Background()

	// STEP: get earliest date [set as last date for analyze]
	now := util.Now()
	tradeEnd_UTC := util.EarliestDate(now, activePost.Duration)
	tradeEnd_UTC_TS := util.DateToTimestamp(tradeEnd_UTC)

	// STEP: get analyze data
	status, hitPrice, hitDate, err := analyzer.Analyze(activePost.Currency, activePost.TakeProfit, activePost.Stop, activePost.LastOperatedAt, tradeEnd_UTC_TS, activePost.Direction)
	if err != nil {
		return db.PostCryptoStatusFail, fmt.Errorf("error while analyzing active post: %s", err.Error())
	}

	switch status {
	// STEP: hit to -> take profit
	case db.PostCryptoStatusSuccess:
		return db.PostCryptoStatusSuccess, service.UpdateFinishedPostStatus(&ctx, activePost, status, activePost.TakeProfit, hitPrice, hitDate)

	// STEP: hit to -> stop
	case db.PostCryptoStatusFail:
		return db.PostCryptoStatusFail, service.UpdateFinishedPostStatus(&ctx, activePost, status, activePost.Stop, hitPrice, hitDate)

	// STEP: no hit, still pending
	case db.PostCryptoStatusPending:
		// STEP: duration is over but no hit
		if activePost.Duration.Before(now) {
			system.Log("no hit & duration is over for post:", activePost.ID)
			return db.PostCryptoStatusFail, service.UpdateFinishedPostStatus(&ctx, activePost, db.PostCryptoStatusFail, hitPrice, hitPrice, hitDate)
		}

		// STEP: still active post
		return db.PostCryptoStatusPending, service.UpdateUnFinishedPostStatus(&ctx, activePost, hitDate)

	default:
		return db.PostCryptoStatusFail, fmt.Errorf("unexpected status: %s", status)
	}
}

func (service *Service) UpdateFinishedPostStatus(ctx *context.Context, activePost db.GetAllActivePostsRow, status db.PostCryptoStatus, lastPrice, hitPrice float64, hitDate int64) error {
	pnl, roi, err := core.CalculatePNL_ROI(activePost.StartPrice, lastPrice, activePost.Leverage, activePost.Direction)
	if err != nil {
		return fmt.Errorf("error while calculating pnl and roi: %s", err.Error())
	}

	params := db.UpdateFinishedPostStatusParams{
		ID:             activePost.ID,
		Status:         status,
		Pnl:            pnl,
		Roi:            roi,
		HitPrice:       sql.NullFloat64{Valid: true, Float64: hitPrice},
		LastOperatedAt: hitDate,
	}

	err = service.Store.UpdateFinishedPostStatus(*ctx, params)
	if err != nil {
		return fmt.Errorf("error while updating post status: %s", err.Error())
	}

	system.Log("post", activePost.ID, "status:", status, "pnl", pnl, "roi(%)", roi)

	return nil
}

func (service *Service) UpdateUnFinishedPostStatus(ctx *context.Context, activePost db.GetAllActivePostsRow, hitDate int64) error {
	params := db.UpdateUnfinishedPostStatusParams{
		ID:             activePost.ID,
		LastOperatedAt: hitDate,
	}

	err := service.Store.UpdateUnfinishedPostStatus(*ctx, params)
	if err != nil {
		return fmt.Errorf("error while updating post status: %s", err.Error())
	}

	return nil
}

func (service *Service) GetAllPendingPayouts(ctx *gin.Context) ([]db.GetAllPendingPayoutsRow, error) {

	pendingPayouts, err := service.Store.GetAllPendingPayouts(ctx)
	if err != nil {
		return nil, err
	}

	return pendingPayouts, nil
}

func (service *Service) PayoutToMoniest(ctx *gin.Context, payoutData db.GetAllPendingPayoutsRow) error {

	// STEP: if there is specific percentage for this payout, otherwise take default one
	operationFeePercentage := service.config.OperationFeePercentage

	if payoutData.OperationFeePercentage.Valid {
		operationFeePercentage = payoutData.OperationFeePercentage.Float64
	}

	// STEP: make payout to moniest
	requestID, _, err := binance.CreateTransfer(service.config, payoutData.Amount, operationFeePercentage, binance.BINANCE_TRANSFER_TYPE_MERCHANT_PAYMENT, string(payoutData.MoniestPayoutType), payoutData.MoniestPayoutValue, binance.BINANCE_TRANSFER_REMARK_PAYOUT)
	if err != nil {
		err1 := service.Store.UpdateBinancePayoutHistoryPayout(ctx, db.UpdateBinancePayoutHistoryPayoutParams{
			ID:     payoutData.ID,
			Status: db.BinancePayoutStatusFail,
			OperationFeePercentage: sql.NullFloat64{
				Valid:   true,
				Float64: operationFeePercentage,
			},
			FailureMessage: sql.NullString{
				Valid:  true,
				String: err.Error(),
			},
			PayoutRequestID: sql.NullString{
				Valid:  true,
				String: requestID,
			},
		})

		if err1 != nil {
			return fmt.Errorf("error while updating payout history failure for payoutID: %s. %s", payoutData.ID, err1.Error())
		}

		return fmt.Errorf("error while creating payout history for payoutID: %s. %s", payoutData.ID, err.Error())
	}

	err = service.Store.UpdateBinancePayoutHistoryPayout(ctx, db.UpdateBinancePayoutHistoryPayoutParams{
		ID:     payoutData.ID,
		Status: db.BinancePayoutStatusSuccess,
		OperationFeePercentage: sql.NullFloat64{
			Valid:   true,
			Float64: operationFeePercentage,
		},
		PayoutDoneAt: sql.NullTime{
			Valid: true,
			Time:  util.Now(),
		},
		PayoutRequestID: sql.NullString{
			Valid:  true,
			String: requestID,
		},
	})
	if err != nil {
		return fmt.Errorf("error while updating payout history success for payoutID: %s. %s", payoutData.ID, err.Error())
	}

	service.sendPayoutEmail(ctx, payoutData, operationFeePercentage)

	system.Log("Successfull payout for payoutID", payoutData.ID)

	return nil
}

func (service *Service) sendPayoutEmail(ctx *gin.Context, payoutData db.GetAllPendingPayoutsRow, operationFeePercentage float64) {

	// STEP: get moniest and user data
	moniest, err := service.GetMoniestByMoniestID(ctx, payoutData.MoniestID)
	if err != nil {
		system.LogError("sending payout email - getting moniest error", err.Error())
	}

	user, err := service.GetOwnUserByID(ctx, payoutData.UserID)
	if err != nil {
		system.LogError("sending payout email - getting user error", err.Error())
	}

	payoutInfo, err := service.GetMoniestPayoutInfos(ctx, moniest.ID)
	if err != nil {
		system.LogError("sending payout email - getting payout-info error", err.Error())
	}

	if err == nil {
		go mailing.SendPayoutEmail(
			moniest.Email, service.config,
			user.Fullname, user.Username,
			moniest.Fullname, payoutInfo.PayoutMethods.PayoutMethodBinance[0].Value,
			int(payoutData.DateIndex), int(payoutData.DateValue),
			payoutData.TotalAmount, operationFeePercentage,
			moniest.Language,
		)
	}
}

func (service *Service) GetExpiredActiveSubscriptions(ctx *gin.Context) ([]db.UserSubscription, error) {

	// STEP: get expired subscriptions and return
	expiredSubscriptions, err := service.Store.GetExpiredActiveSubscriptions(ctx)
	if err != nil {
		return []db.UserSubscription{}, fmt.Errorf("error while getting expired active subscriptions")
	}

	return expiredSubscriptions, nil
}

func (service *Service) DeactivateExpiredSubscriptions(ctx *gin.Context, expiredSubscription db.UserSubscription) error {

	// Update expired subsctriptions
	err := service.Store.UpdateExpiredActiveSubscription(ctx, expiredSubscription.ID)
	if err != nil {
		return fmt.Errorf("error while updating expired active subscription")
	}

	// Save history of subscription
	params := db.CreateUserSubscriptionHistoryParams{
		ID:                    core.CreateID(),
		UserID:                expiredSubscription.UserID,
		MoniestID:             expiredSubscription.MoniestID,
		TransactionID:         expiredSubscription.LatestTransactionID,
		SubscriptionStartDate: expiredSubscription.SubscriptionStartDate,
		SubscriptionEndDate:   expiredSubscription.SubscriptionEndDate,
	}

	_, err = service.Store.CreateUserSubscriptionHistory(ctx, params)
	if err != nil {
		return fmt.Errorf("error while creating user subscription history")
	}

	service.sendSubscriptionExpiredEmail(ctx, expiredSubscription)

	return nil
}

func (service *Service) sendSubscriptionExpiredEmail(ctx *gin.Context, expiredSubscription db.UserSubscription) {
	moniest, err := service.GetMoniestByMoniestID(ctx, expiredSubscription.MoniestID)
	if err != nil {
		system.LogError("sending subscription expired email - getting moniest error", err.Error())
	}

	user, err := service.GetOwnUserByID(ctx, expiredSubscription.UserID)
	if err != nil {
		system.LogError("sending subscription expired email - getting user error", err.Error())
	}

	oldBinanceTransaction, err := service.Store.GetBinancePaymentTransaction(ctx, expiredSubscription.LatestTransactionID.String)
	if err != nil {
		system.LogError("sending subscription expired email - getting binance transaction details error", err.Error())
	}

	if err == nil {
		go mailing.SendSubscriptionExpiredEmail(user.Email, service.config, user.Fullname, moniest.Fullname, moniest.Username, expiredSubscription.SubscriptionStartDate, expiredSubscription.SubscriptionEndDate, oldBinanceTransaction.MoniestFee, int(oldBinanceTransaction.DateValue), user.Language)
	}
}

func (service *Service) GetExpiredPendingBinanceTransactions(ctx context.Context) ([]db.BinancePaymentTransaction, error) {

	// STEP: get expired pending transactions
	expiredPendingTransactions, err := service.Store.GetExpiredPendingBinanceTransactions(ctx)
	if err != nil {
		return []db.BinancePaymentTransaction{}, err
	}

	return expiredPendingTransactions, nil
}

func (service *Service) UpdateExpiredPendingBinanceTransaction(ctx context.Context, transactionID string) error {

	// STEP: update expired pending binance transactions
	err := service.Store.UpdateExpiredPendingBinanceTransaction(ctx, transactionID)
	if err != nil {
		return err
	}

	return nil
}
