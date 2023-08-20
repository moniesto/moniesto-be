package model

import (
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type RegisterRequest struct {
	Fullname string          `json:"fullname" binding:"required,min=1"`
	Username string          `json:"username" binding:"required"`
	Email    string          `json:"email" binding:"required"`
	Password string          `json:"password" binding:"required"`
	Language db.UserLanguage `json:"language" binding:"required"`
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

type SendVerificationEmailRequest struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
}

type VerifyEmailRequest struct {
	Token string `json:"token" binding:"required"`
}

type VerifyEmailResponse struct {
	RedirectURL string `json:"redirect_url"`
}

type ChangeUsernameRequest struct {
	NewUsername string `json:"new" binding:"required"`
}

type ChangeUsernameResponse struct {
	Token string `json:"token"`
}

// type UpdateProfileRequest struct {
// 	Fullname            string `json:"fullname"`
// 	Location        string `json:"location"`
// 	ProfilePhoto    string `json:"profile_photo"`
// 	BackgroundPhoto string `json:"background_photo"`

// 	// special for moniest
// 	Bio         string `json:"bio"`
// 	Description string `json:"description"`
// }

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
		},
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

		response.User.Moniest = moniest
	}

	return
}
