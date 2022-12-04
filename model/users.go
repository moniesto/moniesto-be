package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

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

type User struct {
	Id                           string    `json:"id,omitempty"`
	Name                         string    `json:"name,omitempty"`
	Surname                      string    `json:"surname,omitempty"`
	Username                     string    `json:"username,omitempty"`
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
func NewGetOwnUserResponse(user db.GetOwnUserByUsernameRow) (response OwnUser) {
	response = OwnUser{
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
	}

	if user.MoniestID.String != "" {
		response.Moniest = &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Description.String,
			Score:       user.Score.Float64,
		}
	}

	return
}

func NewGetUserResponse(user db.GetUserByUsernameRow) (response User) {
	response = User{
		Id:                           user.ID,
		Name:                         user.Name,
		Surname:                      user.Surname,
		Username:                     user.Username,
		EmailVerified:                user.EmailVerified,
		Location:                     user.Location.String,
		ProfilePhotoLink:             user.ProfilePhotoLink.(string),
		ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
		BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
		BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
		CreatedAt:                    user.CreatedAt,
		UpdatedAt:                    user.UpdatedAt,
	}

	if user.MoniestID.String != "" {
		response.Moniest = &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Description.String,
			Score:       user.Score.Float64,
		}
	}

	return
}
