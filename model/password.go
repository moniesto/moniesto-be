package model

type ChangePasswordRequest struct {
	// Send Email case
	Email string `json:"email"`

	// Verify & Change case [unauth]
	Token       string `json:"token"`
	NewPassword string `json:"new"`

	// Change Case [auth]
	OldPassword string `json:"old"`
}
