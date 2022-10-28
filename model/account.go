package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Surname  string `json:"surname" binding:"required,min=1"`
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string  `json:"token"`
	User  OwnUser `json:"user"`
}

type OwnUser struct {
	Id                       string
	Name                     string
	Surname                  string
	Username                 string
	Email                    string
	EmailVerified            bool
	Location                 string
	ProfilePhoto             string
	ProfilePhotoThumbnail    string
	BackgroundPhoto          string
	BackgroundPhotoThumbnail string
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	Moniest                  struct {
		Bio              string  `json:"bio"`
		Description      string  `json:"description"`
		Score            float64 `json:"score"`
		SubscriptionInfo struct {
			Fee       float64   `json:""`
			Message   string    `json:"message"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"subscription_info"`
	} `json:"moniest"`
}

func NewRegisterResponse(token string, user db.LoginUserByEmailRow) RegisterResponse {
	return RegisterResponse{
		Token: token,
	}
}
