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
	CardID      string  `json:"card_id" binding:"required"`
}

type Moniest struct {
	ID               string            `json:"id,omitempty"`
	Bio              string            `json:"bio,omitempty"`
	Description      string            `json:"description,omitempty"`
	Score            float64           `json:"score"`
	SubscriptionInfo *SubscriptionInfo `json:"subscription_info,omitempty"`
}

type SubscriptionInfo struct {
	Fee       float64   `json:"fee,omitempty"`
	Message   string    `json:"message,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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
		},
	}

	return response
}
