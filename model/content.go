package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type PaginationRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type GetContentPostRequest struct {
	Subscribed bool   `form:"subscribed" json:"subscribed"`
	Active     bool   `form:"active" json:"active"`
	SortBy     string `form:"sortBy" json:"sortBy"` // created_at | score
	Limit      int    `form:"limit" json:"limit"`
	Offset     int    `form:"offset" json:"offset"`
}

type PostDBResponse []db.GetSubscribedActivePostsRow
type ContentMoniestDBResponse []db.GetMoniestsRow

type MoniestDBResponse []db.SearchMoniestsRow

type OwnPostDBResponse []db.GetOwnActivePostsByUsernameRow

type GetContentPostResponse struct {
	ID          string              `json:"id"`
	Currency    string              `json:"currency"`
	StartPrice  float64             `json:"start_price"`
	Duration    time.Time           `json:"duration"`
	Target1     float64             `json:"target1"`
	Target2     float64             `json:"target2"`
	Target3     float64             `json:"target3"`
	Stop        float64             `json:"stop"`
	Description string              `json:"description,omitempty"`
	Direction   db.EntryPosition    `json:"direction"`
	Finished    bool                `json:"finished"`
	Status      db.PostCryptoStatus `json:"status"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	User        User                `json:"user"`
}

type GetOwnPostResponse struct {
	ID          string              `json:"id"`
	Currency    string              `json:"currency"`
	StartPrice  float64             `json:"start_price"`
	Duration    time.Time           `json:"duration"`
	Target1     float64             `json:"target1"`
	Target2     float64             `json:"target2"`
	Target3     float64             `json:"target3"`
	Stop        float64             `json:"stop"`
	Description string              `json:"description,omitempty"`
	Direction   db.EntryPosition    `json:"direction"`
	Score       float64             `json:"score,omitempty"`
	Finished    bool                `json:"finished"`
	Status      db.PostCryptoStatus `json:"status"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	User        User                `json:"user"`
}

type GetContentMoniestResponse struct {
	Id                           string          `json:"id,omitempty"`
	Name                         string          `json:"name,omitempty"`
	Surname                      string          `json:"surname,omitempty"`
	Username                     string          `json:"username,omitempty"`
	EmailVerified                bool            `json:"email_verified"`
	Location                     string          `json:"location,omitempty"`
	ProfilePhotoLink             string          `json:"profile_photo_link,omitempty"`
	ProfilePhotoThumbnailLink    string          `json:"profile_photo_thumbnail_link,omitempty"`
	BackgroundPhotoLink          string          `json:"background_photo_link,omitempty"`
	BackgroundPhotoThumbnailLink string          `json:"background_photo_thumbnail_link,omitempty"`
	CreatedAt                    *time.Time      `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time      `json:"updated_at,omitempty"`
	Moniest                      *contentMoniest `json:"moniest,omitempty"`
}

type contentMoniest struct {
	ID               string            `json:"id,omitempty"`
	Bio              string            `json:"bio,omitempty"`
	Description      string            `json:"description,omitempty"`
	SubscriberCount  int64             `json:"subscriber_count"`
	Score            float64           `json:"score"`
	SubscriptionInfo *SubscriptionInfo `json:"subscription_info,omitempty"`
}

type GetContentMoniestRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type SearchMoniestRequest struct {
	SearchText string `form:"searchText" json:"searchText" binding:"required,min=1"`
	Limit      int    `form:"limit" json:"limit"`
	Offset     int    `form:"offset" json:"offset"`
}

func NewGetContentPostResponse(posts PostDBResponse) []GetContentPostResponse {
	response := make([]GetContentPostResponse, 0, len(posts))

	for _, post := range posts {
		response = append(response, GetContentPostResponse{
			ID:          post.ID,
			Currency:    post.Currency,
			StartPrice:  post.StartPrice,
			Duration:    post.Duration,
			Target1:     post.Target1,
			Target2:     post.Target2,
			Target3:     post.Target3,
			Stop:        post.Stop,
			Description: post.PostDescription.String,
			Direction:   post.Direction,
			Finished:    post.Finished,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
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

func NewGetOwnPostResponse(posts OwnPostDBResponse) []GetOwnPostResponse {
	response := make([]GetOwnPostResponse, 0, len(posts))

	for _, post := range posts {
		response = append(response, GetOwnPostResponse{
			ID:          post.ID,
			Currency:    post.Currency,
			StartPrice:  post.StartPrice,
			Duration:    post.Duration,
			Target1:     post.Target1,
			Target2:     post.Target2,
			Target3:     post.Target3,
			Stop:        post.Stop,
			Description: post.PostDescription.String,
			Direction:   post.Direction,
			Score:       post.Score,
			Finished:    post.Finished,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
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

func NewGetMoniestsResponse(moniests MoniestDBResponse) []User {
	response := make([]User, 0, len(moniests))

	for _, user := range moniests {
		response = append(response, User{
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
			CreatedAt:                    &user.CreatedAt,
			UpdatedAt:                    &user.UpdatedAt,
			Moniest: &Moniest{
				ID:          user.MoniestID,
				Bio:         user.Bio.String,
				Description: user.Description.String,
				Score:       user.Score,
				SubscriptionInfo: &SubscriptionInfo{
					Fee:       user.Fee,
					Message:   user.Message.String,
					UpdatedAt: user.SubscriptionInfoUpdatedAt,
				},
			},
		})
	}

	return response
}

func NewGetContentMoniestResponse(moniests ContentMoniestDBResponse) []GetContentMoniestResponse {
	response := make([]GetContentMoniestResponse, 0, len(moniests))

	for _, user := range moniests {
		response = append(response, GetContentMoniestResponse{
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
			CreatedAt:                    &user.CreatedAt,
			UpdatedAt:                    &user.UpdatedAt,
			Moniest: &contentMoniest{
				ID:              user.MoniestID,
				Bio:             user.Bio.String,
				Description:     user.Description.String,
				SubscriberCount: user.UserSubscriptionCount,
				Score:           user.Score,
				SubscriptionInfo: &SubscriptionInfo{
					Fee:       user.Fee,
					Message:   user.Message.String,
					UpdatedAt: user.SubscriptionInfoUpdatedAt,
				},
			},
		})
	}

	return response
}
