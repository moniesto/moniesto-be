package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type CreateMoniestRequest struct {
	Bio         string  `json:"bio" binding:"required"`
	Description string  `json:"description"` // optional
	Fee         float64 `json:"fee" binding:"required"`
	Message     string  `json:"message"` // optional
	BinanceID   string  `json:"binance_id" binding:"required"`
}

type SubscribeMoniestRequest struct {
	NumberOfMonths int    `json:"number_of_months" binding:"required,min=1"`
	ReturnURL      string `json:"returnURL" binding:"required"`
	CancelURL      string `json:"cancelURL" binding:"required"`
}

type SubscribeMoniestResponse struct {
	QrcodeLink    string `json:"qrcode_link"`
	CheckoutLink  string `json:"checkout_link"`
	DeepLink      string `json:"deep_link"`
	UniversalLink string `json:"universal_link"`
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

type GetSubscriptionInfoResponse struct {
	Subscribed bool  `json:"subscribed"`
	Pending    *bool `json:"pending,omitempty"`

	// when pending = true
	Timeout       *int    `json:"timeout,omitempty"`
	QrcodeLink    *string `json:"qrcode_link,omitempty"`
	CheckoutLink  *string `json:"checkout_link,omitempty"`
	DeepLink      *string `json:"deep_link,omitempty"`
	UniversalLink *string `json:"universal_link,omitempty"`

	// when subscribed = true
	SubscriptionInfo *SubscriptionInfo `json:"subscription_info,omitempty"`
}

type SubscriptionInfo struct {
	SubscriptionStartDate time.Time `json:"subscription_start_date"`
	SubscriptionEndDate   time.Time `json:"subscription_end_date"`
	PayerID               string    `json:"payer_id"`
	SubscribedFee         float64   `json:"subscribed_fee"`
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

type UpdateMoniestPayoutInfo struct {
	BinanceID string `json:"binance_id" binding:"required"`
}

type GetMoniestPayoutInfos struct {
	PayoutMethods PayoutMethods `json:"payout_methods"`
}

type PayoutMethods struct {
	PayoutMethodBinance []PayoutMethodBinance `json:"binance"`
}

type PayoutMethodBinance struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// MAKER
func NewCreateMoniestResponse(moniest db.GetMoniestByMoniestIdRow) OwnUser {
	response := OwnUser{
		Id:                           moniest.ID,
		Fullname:                     moniest.Fullname,
		Username:                     moniest.Username,
		Email:                        moniest.Email,
		EmailVerified:                moniest.EmailVerified,
		Language:                     moniest.Language,
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
