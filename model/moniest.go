package model

import "time"

type CreateMoniestRequest struct {
	Bio         string  `json:"bio"`         // optional
	Description string  `json:"description"` // optional
	Fee         float64 `json:"fee" binding:"required"`
	Message     string  `json:"message"` // optional
	CardID      string  `json:"card_id" binding:"required"`
}

type Moniest struct {
	ID               string            `json:"id,omitempty"`
	Bio              string            `json:"bio,omitempty"`
	Description      string            `json:"description,omitempty"`
	Score            float64           `json:"score,omitempty"`
	SubscriptionInfo *SubscriptionInfo `json:"subscription_info,omitempty"`
}

type SubscriptionInfo struct {
	Fee       float64   `json:"fee,omitempty"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
