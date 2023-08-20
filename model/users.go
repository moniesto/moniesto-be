package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type UserDBResponse []db.GetUserByUsernameRow

type UserStatResponse struct {
	SubscriptionCount int64 `json:"subscription_count"`
}

type OwnUser struct {
	Id                           string          `json:"id,omitempty"`
	Fullname                     string          `json:"fullname,omitempty"`
	Username                     string          `json:"username,omitempty"`
	Email                        string          `json:"email,omitempty"`
	EmailVerified                bool            `json:"email_verified"`
	Language                     db.UserLanguage `json:"language"`
	Location                     string          `json:"location,omitempty"`
	ProfilePhotoLink             string          `json:"profile_photo_link,omitempty"`
	ProfilePhotoThumbnailLink    string          `json:"profile_photo_thumbnail_link,omitempty"`
	BackgroundPhotoLink          string          `json:"background_photo_link,omitempty"`
	BackgroundPhotoThumbnailLink string          `json:"background_photo_thumbnail_link,omitempty"`
	CreatedAt                    time.Time       `json:"created_at,omitempty"`
	UpdatedAt                    time.Time       `json:"updated_at,omitempty"`
	Moniest                      *Moniest        `json:"moniest,omitempty"`
}

type User struct {
	Id                           string     `json:"id,omitempty"`
	Fullname                     string     `json:"fullname,omitempty"`
	Username                     string     `json:"username,omitempty"`
	EmailVerified                bool       `json:"email_verified"`
	Location                     string     `json:"location,omitempty"`
	ProfilePhotoLink             string     `json:"profile_photo_link,omitempty"`
	ProfilePhotoThumbnailLink    string     `json:"profile_photo_thumbnail_link,omitempty"`
	BackgroundPhotoLink          string     `json:"background_photo_link,omitempty"`
	BackgroundPhotoThumbnailLink string     `json:"background_photo_thumbnail_link,omitempty"`
	CreatedAt                    *time.Time `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time `json:"updated_at,omitempty"`
	Moniest                      *Moniest   `json:"moniest,omitempty"`
}

type UpdateUserProfileRequest struct {
	Fullname        string          `json:"fullname"`
	Location        string          `json:"location"`
	ProfilePhoto    string          `json:"profile_photo"`
	BackgroundPhoto string          `json:"background_photo"`
	Language        db.UserLanguage `json:"language"`
}

// MAKER
func NewGetOwnUserResponseByUsername(user db.GetOwnUserByUsernameRow) (response OwnUser) {
	response = OwnUser{
		Id:                           user.ID,
		Fullname:                     user.Fullname,
		Username:                     user.Username,
		Email:                        user.Email,
		EmailVerified:                user.EmailVerified,
		Language:                     user.Language,
		Location:                     user.Location.String,
		ProfilePhotoLink:             user.ProfilePhotoLink.(string),
		ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
		BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
		BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
		CreatedAt:                    user.CreatedAt,
		UpdatedAt:                    user.UpdatedAt,
	}

	if user.MoniestID.String != "" {
		moniest := &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Description.String,
			Score:       user.Score.Float64,
		}

		if user.MoniestSubscriptionInfoID.Valid {
			moniest.MoniestSubscriptionInfo = &MoniestSubscriptionInfo{
				Fee:       user.Fee.Float64,
				Message:   user.Message.String,
				UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt.Time,
			}
		}

		response.Moniest = moniest
	}

	return
}

func NewGetOwnUserResponseByID(user db.GetOwnUserByIDRow) (response OwnUser) {
	response = OwnUser{
		Id:                           user.ID,
		Fullname:                     user.Fullname,
		Username:                     user.Username,
		Email:                        user.Email,
		EmailVerified:                user.EmailVerified,
		Language:                     user.Language,
		Location:                     user.Location.String,
		ProfilePhotoLink:             user.ProfilePhotoLink.(string),
		ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
		BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
		BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
		CreatedAt:                    user.CreatedAt,
		UpdatedAt:                    user.UpdatedAt,
	}

	if user.MoniestID.String != "" {
		moniest := &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Description.String,
			Score:       user.Score.Float64,
		}

		if user.MoniestSubscriptionInfoID.Valid {
			moniest.MoniestSubscriptionInfo = &MoniestSubscriptionInfo{
				Fee:       user.Fee.Float64,
				Message:   user.Message.String,
				UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt.Time,
			}
		}

		response.Moniest = moniest
	}

	return
}

func NewGetUserResponse(user db.GetUserByUsernameRow) (response User) {
	response = User{
		Id:                           user.ID,
		Fullname:                     user.Fullname,
		Username:                     user.Username,
		EmailVerified:                user.EmailVerified,
		Location:                     user.Location.String,
		ProfilePhotoLink:             user.ProfilePhotoLink.(string),
		ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
		BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
		BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
		CreatedAt:                    &user.CreatedAt,
		UpdatedAt:                    &user.UpdatedAt,
	}

	if user.MoniestID.String != "" {
		moniest := &Moniest{
			ID:          user.MoniestID.String,
			Bio:         user.Bio.String,
			Description: user.Description.String,
			Score:       user.Score.Float64,
		}

		if user.MoniestSubscriptionInfoID.Valid {
			moniest.MoniestSubscriptionInfo = &MoniestSubscriptionInfo{
				Fee:       user.Fee.Float64,
				Message:   user.Message.String,
				UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt.Time,
			}
		}

		response.Moniest = moniest
	}

	return
}

func NewGetUsersResponse(users []db.GetUserByUsernameRow) []User {
	responses := make([]User, 0, len(users))

	for _, user := range users {
		response := User{
			Id:                           user.ID,
			Fullname:                     user.Fullname,
			Username:                     user.Username,
			EmailVerified:                user.EmailVerified,
			Location:                     user.Location.String,
			ProfilePhotoLink:             user.ProfilePhotoLink.(string),
			ProfilePhotoThumbnailLink:    user.ProfilePhotoThumbnailLink.(string),
			BackgroundPhotoLink:          user.BackgroundPhotoLink.(string),
			BackgroundPhotoThumbnailLink: user.BackgroundPhotoThumbnailLink.(string),
			CreatedAt:                    &user.CreatedAt,
			UpdatedAt:                    &user.UpdatedAt,
		}

		if user.MoniestID.String != "" {
			moniest := &Moniest{
				ID:          user.MoniestID.String,
				Bio:         user.Bio.String,
				Description: user.Description.String,
				Score:       user.Score.Float64,
			}

			if user.MoniestSubscriptionInfoID.Valid {
				moniest.MoniestSubscriptionInfo = &MoniestSubscriptionInfo{
					Fee:       user.Fee.Float64,
					Message:   user.Message.String,
					UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt.Time,
				}
			}

			response.Moniest = moniest
		}

		responses = append(responses, response)
	}

	return responses
}
