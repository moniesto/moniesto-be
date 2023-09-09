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
			// Id:                           user.ID, // TODO: remove after confirmation
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
			// ID:          user.MoniestID.String, // TODO: remove after confirmation
			Bio:         user.Bio.String,
			Description: user.Description.String,
		}

		if user.MoniestSubscriptionInfoID.Valid {
			moniest.MoniestSubscriptionInfo = &MoniestSubscriptionInfo{
				Fee:       user.Fee.Float64,
				Message:   user.Message.String,
				UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt.Time,
			}
		}

		if user.Pnl7days.Valid {
			moniest.CryptoPostStatistics = &CryptoPostStatistics{
				Pnl7days:      user.Pnl7days.Float64,
				Roi7days:      user.Roi7days.Float64,
				WinRate7days:  user.WinRate7days.Float64,
				Pnl30days:     user.Pnl30days.Float64,
				Roi30days:     user.Roi30days.Float64,
				WinRate30days: user.WinRate30days.Float64,
				PnlTotal:      user.PnlTotal.Float64,
				RoiTotal:      user.RoiTotal.Float64,
				WinRateTotal:  user.WinRateTotal.Float64,
			}
		}

		response.User.Moniest = moniest
	}

	return
}
