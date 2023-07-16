package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/scoring"
	"github.com/moniesto/moniesto-be/util/systemError"
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
		fmt.Println("ERROR while marshall", err)
		return err
	}
	fmt.Println("RESPONSE", string(b))

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

func (service *Service) GetAllPendingPayouts() ([]db.GetAllPendingPayoutsRow, error) {

	ctx := context.Background()

	pendingPayouts, err := service.Store.GetAllPendingPayouts(ctx)
	if err != nil {
		return nil, err
	}

	return pendingPayouts, nil
}

func (service *Service) PayoutToMoniest(payoutData db.GetAllPendingPayoutsRow) error {

	// STEP: if there is specific percentage for this payout, otherwise take default one
	operationFeePercentage := service.config.OperationFeePercentage

	if payoutData.OperationFeePercentage.Valid {
		operationFeePercentage = payoutData.OperationFeePercentage.Float64
	}

	ctx := context.Background()

	// STEP: make payout to moniest
	requestID, _, err := binance.CreatePayout(service.config, payoutData.Amount, operationFeePercentage, string(payoutData.MoniestPayoutType), payoutData.MoniestPayoutValue)
	if err != nil {
		err = service.Store.UpdatePayoutHistory(ctx, db.UpdatePayoutHistoryParams{
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

		if err != nil {
			return fmt.Errorf("error while updating payout history failure for payoutID: %s. %s", payoutData.ID, err.Error())
		}

		return fmt.Errorf("error while creating payout history for payoutID: %s. %s", payoutData.ID, err.Error())
	}

	err = service.Store.UpdatePayoutHistory(ctx, db.UpdatePayoutHistoryParams{
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
	systemError.Log("Successfull payout for payoutID", payoutData.ID)

	// TODO: send email to moniest

	return nil
}
