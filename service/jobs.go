package service

import (
	"context"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
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

	req := model.CalculateScoreRequest{
		Parity:               activePost.Currency,
		StartPrice:           activePost.StartPrice,
		StartDate:            activePost.CreatedAt.Unix(),
		EndDate:              activePost.Duration.Unix(),
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

	// STEP: post is finished case
	if response.Finished {
		/*
			UPDATE
				success case
				score
		*/
	} else { // STEP: post is not finished
		/*
			UPDATE
				lastTargetHit
				lastCronJobTimeStamp
		*/
	}

	return nil
}
