package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type CreatePostRequest struct {
	Currency    string    `json:"currency" binding:"required"`
	Duration    time.Time `json:"duration" binding:"required"`
	Target1     float64   `json:"target1" binding:"required"`
	Target2     float64   `json:"target2" binding:"required"`
	Target3     float64   `json:"target3" binding:"required"`
	Stop        float64   `json:"stop" binding:"required"`
	Direction   string    `json:"direction" binding:"required"`
	Description string    `json:"description"`
}

type CreatePostResponse struct {
	ID         string           `json:"id"`
	MoniestID  string           `json:"moniest_id"`
	Currency   string           `json:"currency"`
	StartPrice float64          `json:"start_price"`
	Duration   time.Time        `json:"duration"`
	Target1    float64          `json:"target1"`
	Target2    float64          `json:"target2"`
	Target3    float64          `json:"target3"`
	Stop       float64          `json:"stop"`
	Direction  db.EntryPosition `json:"direction"`
	Score      float64          `json:"score"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`

	Description string `json:"description,omitempty"`
}

/*
{
 "id": "7ad33645-fbe5-49e0-9284-e72cd5b7e312",
 "moniest_id": "3a770db0-f737-406c-ba57-406dc13879c5",
 "currency": "BTCUSDT",
 "start_price": 16556.68,
 "duration": "2023-12-31T23:33:47.479Z",
 "target1": 17000,
 "target2": 17001,
 "target3": 17002,
 "stop": 15000,
 "direction": "long",
 "score": 0,
 "created_at": "2023-01-01T15:57:10.466862Z",
 "updated_at": "2023-01-01T15:57:10.466862Z"
}
*/

func NewCreatePostResponse(post db.CreatePostRow, description db.PostCryptoDescription) CreatePostResponse {
	return CreatePostResponse{
		ID:          post.ID,
		MoniestID:   post.MoniestID,
		Currency:    post.Currency,
		StartPrice:  post.StartPrice,
		Duration:    post.Duration,
		Target1:     post.Target1,
		Target2:     post.Target2,
		Target3:     post.Target3,
		Stop:        post.Stop,
		Direction:   post.Direction,
		Score:       post.Score,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Description: description.Description,
	}
}
