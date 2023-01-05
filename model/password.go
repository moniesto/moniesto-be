package model

type ChangePasswordRequest struct {
	NewPassword string `json:"new" binding:"required"`
	OldPassword string `json:"old" binding:"required"`
}

type SendResetPasswordEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type VerifyPasswordResetRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new" binding:"required"`
}

type VerifyTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
