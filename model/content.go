package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type GetContentPostRequest struct {
	Subscribed bool `form:"subscribed" json:"subscribed"`
	Active     bool `form:"active" json:"active"`
	Limit      int  `form:"limit" json:"limit"`
	Offset     int  `form:"offset" json:"offset"`
}

type PostDBResponse []db.GetSubscribedActivePostsRow

type GetContentPostResponse struct {
	ID         string              `json:"id"`
	Currency   string              `json:"currency"`
	StartPrice float64             `json:"start_price"`
	Duration   time.Time           `json:"duration"`
	Target1    float64             `json:"target1"`
	Target2    float64             `json:"target2"`
	Target3    float64             `json:"target3"`
	Stop       float64             `json:"stop"`
	Direction  db.EntryPosition    `json:"direction"`
	Finished   bool                `json:"finished"`
	Status     db.PostCryptoStatus `json:"status"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	User       User                `json:"user"`
}

type GetContentMoniestRequest struct {
	Subscribed bool `form:"subscribed" json:"subscribed"`
	Limit      int  `form:"limit" json:"limit"`
	Offset     int  `form:"offset" json:"offset"`
}

func NewGetContentPostResponse(posts PostDBResponse) []GetContentPostResponse {
	response := make([]GetContentPostResponse, 0, len(posts))

	for _, post := range posts {
		response = append(response, GetContentPostResponse{
			ID:         post.ID,
			Currency:   post.Currency,
			StartPrice: post.StartPrice,
			Duration:   post.Duration,
			Target1:    post.Target1,
			Target2:    post.Target2,
			Target3:    post.Target3,
			Stop:       post.Stop,
			Direction:  post.Direction,
			Finished:   post.Finished,
			Status:     post.Status,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
			User: User{
				Id:                           post.UserID,
				Name:                         post.Name,
				Surname:                      post.Surname,
				Username:                     post.Username,
				EmailVerified:                post.EmailVerified,
				ProfilePhotoLink:             post.ProfilePhotoLink.(string),
				ProfilePhotoThumbnailLink:    post.ProfilePhotoThumbnailLink.(string),
				BackgroundPhotoLink:          post.BackgroundPhotoLink.(string),
				BackgroundPhotoThumbnailLink: post.BackgroundPhotoThumbnailLink.(string),
				Moniest: &Moniest{
					ID:          post.MoniestID,
					Bio:         post.Bio.String,
					Description: post.Description.String,
					Score:       post.MoniestScore,
				},
			},
		})
	}

	return response
}
