package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Surname  string `json:"surname" binding:"required,min=1"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string  `json:"token"`
	User  OwnUser `json:"user"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required,min=1"`
	Password   string `json:"password" binding:"required,min=6"`
}

type LoginResponse RegisterResponse

type CheckUsernameResponse struct {
	Validity bool `json:"validity"`
}

type OwnUser struct {
	Id                           string    `json:"id,omitempty"`
	Name                         string    `json:"name,omitempty"`
	Surname                      string    `json:"surname,omitempty"`
	Username                     string    `json:"username,omitempty"`
	Email                        string    `json:"email,omitempty"`
	EmailVerified                bool      `json:"email_verified"`
	Location                     string    `json:"location,omitempty"`
	ProfilePhotoLink             string    `json:"profile_photo_link,omitempty"`
	ProfilePhotoThumbnailLink    string    `json:"profile_photo_thumbnail_link,omitempty"`
	BackgroundPhotoLink          string    `json:"background_photo_link,omitempty"`
	BackgroundPhotoThumbnailLink string    `json:"background_photo_thumbnail_link,omitempty"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	Moniest                      *Moniest  `json:"moniest,omitempty"`
}

// MAKER

// NewLoginResponse creates/return LoginResponse object
func NewLoginResponse(token string, user db.LoginUserByEmailRow) (response LoginResponse) {
	// asserting RegisterResponse to LoginResponse
	return LoginResponse(NewRegisterResponse(token, user))
}

// NewCheckUsernameResponse creates/return CheckUsernameResponse object
func NewCheckUsernameResponse(validity bool) (response CheckUsernameResponse) {
	response = CheckUsernameResponse{
		Validity: validity,
	}

	return
}

// NewRegisterResponse creates/return RegisterResponse object
func NewRegisterResponse(token string, user db.LoginUserByEmailRow) (response RegisterResponse) {
	response = RegisterResponse{
		Token: token,
		User: OwnUser{
			Id:                           user.ID,
			Name:                         user.Name,
			Surname:                      user.Surname,
			Username:                     user.Username,
			Email:                        user.Email,
			EmailVerified:                user.EmailVerified,
			Location:                     user.Location.String,
			ProfilePhotoLink:             user.ProfilePhotoLink.(string),
			ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
			BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
			BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
			CreatedAt:                    user.CreatedAt,
			UpdatedAt:                    user.UpdatedAt,
		},
	}

	if user.MoniestID.String != "" {
		response.User.Moniest = &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Bio.String,
			Score:       user.Score.Float64,
		}
	}

	return
}
