package model

import (
	"fmt"
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
)

type CreatePostRequest struct {
	MarketType string    `json:"market_type" binding:"required"`
	Currency   string    `json:"currency" binding:"required"`
	Duration   time.Time `json:"duration" binding:"required"`
	TakeProfit float64   `json:"take_profit" binding:"required"`
	Stop       float64   `json:"stop" binding:"required"`

	Target1 *float64 `json:"target1"`
	Target2 *float64 `json:"target2"`
	Target3 *float64 `json:"target3"`

	Direction   string `json:"direction" binding:"required"`
	Leverage    int32  `json:"leverage"`
	Description string `json:"description"`
}

type CreatePostResponse struct {
	ID         string                  `json:"id"`
	MoniestID  string                  `json:"moniest_id"`
	MarketType db.PostCryptoMarketType `json:"market_type"`
	Currency   string                  `json:"currency"`
	StartPrice float64                 `json:"start_price"`
	Duration   time.Time               `json:"duration"`
	TakeProfit float64                 `json:"take_profit"`
	Stop       float64                 `json:"stop"`
	Target1    *float64                `json:"target1,omitempty"`
	Target2    *float64                `json:"target2,omitempty"`
	Target3    *float64                `json:"target3,omitempty"`
	Direction  db.EntryPosition        `json:"direction"`
	Leverage   int32                   `json:"leverage"`
	Pnl        float64                 `json:"pnl"`
	Roi        float64                 `json:"roi"`
	CreatedAt  time.Time               `json:"created_at"`
	UpdatedAt  time.Time               `json:"updated_at"`

	Description string `json:"description,omitempty"`
}

type PostDescriptionType struct {
	Time   int64              `json:"time"`
	Blocks []DescriptionBlock `json:"blocks"`
}

type DescriptionBlock struct {
	ID   string                 `json:"id"`
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func NewCreatePostResponse(post db.CreatePostRow, description db.PostCryptoDescription) CreatePostResponse {

	fmt.Println("post.Target1", post.Target1)
	fmt.Println("post.Target2", post.Target2)
	fmt.Println("post.Target3", post.Target3)

	return CreatePostResponse{
		ID:          post.ID,
		MoniestID:   post.MoniestID,
		MarketType:  post.MarketType,
		Currency:    post.Currency,
		StartPrice:  post.StartPrice,
		Duration:    post.Duration,
		TakeProfit:  post.TakeProfit,
		Stop:        post.Stop,
		Target1:     util.SafeSQLNullToFloat(post.Target1),
		Target2:     util.SafeSQLNullToFloat(post.Target2),
		Target3:     util.SafeSQLNullToFloat(post.Target3),
		Direction:   post.Direction,
		Leverage:    post.Leverage,
		Pnl:         post.Pnl,
		Roi:         post.Roi,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Description: description.Description,
	}
}
