package service

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/scoring"
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
