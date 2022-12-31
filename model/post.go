package model

import "time"

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
