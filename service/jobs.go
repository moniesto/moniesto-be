package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/scoring"
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

func (service *Service) UpdatePostStatus(activePost db.GetAllActivePostsRow) error {
	ctx := context.Background()

	req := model.CalculateScoreRequest{
		Parity:               activePost.Currency,
		StartPrice:           activePost.StartPrice,
		StartDate:            util.DateToTimestamp(activePost.CreatedAt),
		EndDate:              util.DateToTimestamp(activePost.Duration),
		Target1:              activePost.Target1,
		Target2:              activePost.Target2,
		Target3:              activePost.Target3,
		Stop:                 activePost.Stop,
		Direction:            string(activePost.Direction),
		LastTargetHit:        activePost.LastTargetHit,
		LastCronJobTimeStamp: activePost.LastJobTimestamp,
	}

	response, err := scoring.CalculateScore(req, service.config)
	if err != nil {
		return err
	}

	b, err := json.Marshal(response)
	if err != nil {
		system.LogError("Update Post status - error while marshall", err)
		return err
	}
	system.Log("scoring response", string(b))

	// STEP: post is finished case
	if response.Finished {
		// STEP: update post status
		status := ""

		if response.Success {
			status = string(db.PostCryptoStatusSuccess)
		} else {
			status = string(db.PostCryptoStatusFail)
		}

		params := db.UpdateFinishedPostStatusParams{
			ID:     activePost.ID,
			Status: db.PostCryptoStatus(status),
			Score:  response.Score,
		}
		err := service.Store.UpdateFinishedPostStatus(ctx, params)
		if err != nil {
			return err
		}

		// STEP: update moniest score
		moniestParams := db.UpdateMoniestScoreParams{
			ID:    activePost.MoniestID,
			Score: response.Score,
		}

		err = service.Store.UpdateMoniestScore(ctx, moniestParams)
		if err != nil {
			return err
		}

	} else { // STEP: post is not finished
		params := db.UpdateUnfinishedPostStatusParams{
			ID:               activePost.ID,
			LastTargetHit:    response.LastTargetHit,
			LastJobTimestamp: response.LastCronJobTimeStamp,
		}
		err := service.Store.UpdateUnfinishedPostStatus(ctx, params)
		if err != nil {
			return err
		}
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
