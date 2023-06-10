package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type CreateMoniestRequest struct {
	Bio         string  `json:"bio"`         // optional
	Description string  `json:"description"` // optional
	Fee         float64 `json:"fee" binding:"required"`
	Message     string  `json:"message"` // optional
	BinanceID   string  `json:"binance_id" binding:"required"`
}

type Moniest struct {
	ID                      string                   `json:"id,omitempty"`
	Bio                     string                   `json:"bio,omitempty"`
	Description             string                   `json:"description,omitempty"`
	Score                   float64                  `json:"score"`
	MoniestSubscriptionInfo *MoniestSubscriptionInfo `json:"subscription_info,omitempty"`
}

type MoniestSubscriptionInfo struct {
	Fee       float64   `json:"fee,omitempty"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateMoniestProfileRequest struct {
	Bio                     string                   `json:"bio"`
	Description             string                   `json:"description"`
	MoniestSubscriptionInfo *MoniestSubscriptionInfo `json:"subscription_info"`
}

type CheckSubscriptionResponse struct {
	Subscribed bool `json:"subscribed"`
}

type GetMoniestPostsRequest struct {
	Active bool `form:"active" json:"active"`
	Limit  int  `form:"limit" json:"limit"`
	Offset int  `form:"offset" json:"offset"`
}

type MoniestStatResponse struct {
	SubscriptionCount int64 `json:"subscription_count"`
	SubscriberCount   int64 `json:"subscriber_count"`
	PostCount         int64 `json:"post_count"`
}

// MAKER
func NewCreateMoniestResponse(moniest db.GetMoniestByMoniestIdRow) OwnUser {
	response := OwnUser{
		Id:                           moniest.ID,
		Name:                         moniest.Name,
		Surname:                      moniest.Surname,
		Username:                     moniest.Username,
		Email:                        moniest.Email,
		EmailVerified:                moniest.EmailVerified,
		Location:                     moniest.Location.String,
		ProfilePhotoLink:             moniest.ProfilePhotoLink.(string),
		ProfilePhotoThumbnailLink:    moniest.ProfilePhotoThumbnailLink.(string),
		BackgroundPhotoLink:          moniest.BackgroundPhotoLink.(string),
		BackgroundPhotoThumbnailLink: moniest.BackgroundPhotoThumbnailLink.(string),
		CreatedAt:                    moniest.CreatedAt,
		UpdatedAt:                    moniest.UpdatedAt,
		Moniest: &Moniest{
			ID:          moniest.MoniestID,
			Bio:         moniest.Bio.String,
			Description: moniest.Description.String,
			Score:       moniest.Score,
			MoniestSubscriptionInfo: &MoniestSubscriptionInfo{
				Fee:       moniest.Fee,
				Message:   moniest.Message.String,
				UpdatedAt: moniest.MoniestSubscriptionInfoUpdatedAt,
			},
		},
	}

	return response
}
