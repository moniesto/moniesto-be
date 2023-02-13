package model

type CreateFeedbackRequest struct {
	Type    string `json:"type"` // optional
	Message string `json:"message" binding:"required"`
}
