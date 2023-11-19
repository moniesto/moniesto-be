package model

import (
	"time"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
)

type PaginationRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type GetContentPostRequest struct {
	Subscribed bool   `form:"subscribed" json:"subscribed"`
	Active     bool   `form:"active" json:"active"`
	SortBy     string `form:"sortBy" json:"sortBy"` // created_at | pnl
	Limit      int    `form:"limit" json:"limit"`
	Offset     int    `form:"offset" json:"offset"`
}

// only difference from User struct is type of "Moniest" field
type UserAsContent struct {
	Id                           string            `json:"id,omitempty"`
	Fullname                     string            `json:"fullname,omitempty"`
	Username                     string            `json:"username,omitempty"`
	EmailVerified                bool              `json:"email_verified"`
	Location                     string            `json:"location,omitempty"`
	ProfilePhotoLink             string            `json:"profile_photo_link,omitempty"`
	ProfilePhotoThumbnailLink    string            `json:"profile_photo_thumbnail_link,omitempty"`
	BackgroundPhotoLink          string            `json:"background_photo_link,omitempty"`
	BackgroundPhotoThumbnailLink string            `json:"background_photo_thumbnail_link,omitempty"`
	CreatedAt                    *time.Time        `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time        `json:"updated_at,omitempty"`
	Moniest                      *MoniestAsContent `json:"moniest,omitempty"`
}

// only difference from Moniest struct is additional "SubscriberCount" field
type MoniestAsContent struct {
	Bio         string `json:"bio,omitempty"`
	Description string `json:"description,omitempty"`

	SubscriberCount int64 `json:"subscriber_count"`

	CryptoPostStatistics    *CryptoPostStatistics    `json:"post_statistics,omitempty"`
	MoniestSubscriptionInfo *MoniestSubscriptionInfo `json:"subscription_info,omitempty"`
}

type PostDBResponse []db.GetSubscribedActivePostsRow
type ContentMoniestDBResponse []db.GetMoniestsRow

type MoniestDBResponse []db.SearchMoniestsRow

type OwnPostDBResponse []db.GetOwnActivePostsByUsernameRow

type GetContentPostResponse struct {
	ID          string                  `json:"id"`
	MarketType  db.PostCryptoMarketType `json:"market_type"`
	Currency    string                  `json:"currency"`
	StartPrice  float64                 `json:"start_price"`
	Duration    time.Time               `json:"duration"`
	TakeProfit  float64                 `json:"take_profit"`
	Stop        float64                 `json:"stop"`
	Target1     *float64                `json:"target1,omitempty"`
	Target2     *float64                `json:"target2,omitempty"`
	Target3     *float64                `json:"target3,omitempty"`
	Direction   db.Direction            `json:"direction"`
	Leverage    int32                   `json:"leverage"`
	Pnl         float64                 `json:"pnl"`
	Roi         float64                 `json:"roi"`
	Finished    bool                    `json:"finished"`
	Status      db.PostCryptoStatus     `json:"status"`
	HitPrice    *float64                `json:"hit_price,omitempty"`
	FinishedAt  *time.Time              `json:"finished_at,omitempty"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	Description string                  `json:"description,omitempty"`

	User User `json:"user"`
}

type GetOwnPostResponse struct {
	ID          string                  `json:"id"`
	MarketType  db.PostCryptoMarketType `json:"market_type"`
	Currency    string                  `json:"currency"`
	StartPrice  float64                 `json:"start_price"`
	Duration    time.Time               `json:"duration"`
	TakeProfit  float64                 `json:"take_profit"`
	Stop        float64                 `json:"stop"`
	Target1     *float64                `json:"target1,omitempty"`
	Target2     *float64                `json:"target2,omitempty"`
	Target3     *float64                `json:"target3,omitempty"`
	Direction   db.Direction            `json:"direction"`
	Leverage    int32                   `json:"leverage"`
	Pnl         float64                 `json:"pnl"`
	Roi         float64                 `json:"roi"`
	Finished    bool                    `json:"finished"`
	Status      db.PostCryptoStatus     `json:"status"`
	HitPrice    *float64                `json:"hit_price,omitempty"`
	FinishedAt  *time.Time              `json:"finished_at,omitempty"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	Description string                  `json:"description,omitempty"`

	User User `json:"user"`
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
			MarketType:  post.MarketType,
			Currency:    post.Currency,
			StartPrice:  post.StartPrice,
			Duration:    post.Duration,
			TakeProfit:  post.TakeProfit,
			Stop:        post.Stop,
			Target1:     util.SafeSQLNullToFloat(post.Target1),
			Target2:     util.SafeSQLNullToFloat(post.Target2),
			Target3:     util.SafeSQLNullToFloat(post.Target3),
			Direction:   post.Direction,
			Leverage:    post.Leverage,
			Pnl:         post.Pnl,
			Roi:         post.Roi,
			Finished:    post.Finished,
			Status:      post.Status,
			HitPrice:    util.SafeSQLNullToFloat(post.HitPrice),
			FinishedAt:  util.SafeSQLNullToTime(post.FinishedAt),
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Description: post.PostDescription.String,
			User: User{
				Id:                           post.UserID,
				Fullname:                     post.Fullname,
				Username:                     post.Username,
				EmailVerified:                post.EmailVerified,
				ProfilePhotoLink:             post.ProfilePhotoLink.(string),
				ProfilePhotoThumbnailLink:    post.ProfilePhotoThumbnailLink.(string),
				BackgroundPhotoLink:          post.BackgroundPhotoLink.(string),
				BackgroundPhotoThumbnailLink: post.BackgroundPhotoThumbnailLink.(string),
				Moniest: &Moniest{
					Bio:         post.Bio.String,
					Description: post.Description.String,
					CryptoPostStatistics: &CryptoPostStatistics{
						Pnl7days:      post.Pnl7days.Float64,
						Roi7days:      post.Roi7days.Float64,
						WinRate7days:  post.WinRate7days.Float64,
						Pnl30days:     post.Pnl30days.Float64,
						Roi30days:     post.Roi30days.Float64,
						WinRate30days: post.WinRate30days.Float64,
						PnlTotal:      post.PnlTotal.Float64,
						RoiTotal:      post.RoiTotal.Float64,
						WinRateTotal:  post.WinRateTotal.Float64,
					},
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
			MarketType:  post.MarketType,
			Currency:    post.Currency,
			StartPrice:  post.StartPrice,
			Duration:    post.Duration,
			TakeProfit:  post.TakeProfit,
			Stop:        post.Stop,
			Target1:     util.SafeSQLNullToFloat(post.Target1),
			Target2:     util.SafeSQLNullToFloat(post.Target2),
			Target3:     util.SafeSQLNullToFloat(post.Target3),
			Direction:   post.Direction,
			Leverage:    post.Leverage,
			Pnl:         post.Pnl,
			Roi:         post.Roi,
			Finished:    post.Finished,
			Status:      post.Status,
			HitPrice:    util.SafeSQLNullToFloat(post.HitPrice),
			FinishedAt:  util.SafeSQLNullToTime(post.FinishedAt),
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Description: post.PostDescription.String,
			User: User{
				Id:                           post.UserID,
				Fullname:                     post.Fullname,
				Username:                     post.Username,
				EmailVerified:                post.EmailVerified,
				ProfilePhotoLink:             post.ProfilePhotoLink.(string),
				ProfilePhotoThumbnailLink:    post.ProfilePhotoThumbnailLink.(string),
				BackgroundPhotoLink:          post.BackgroundPhotoLink.(string),
				BackgroundPhotoThumbnailLink: post.BackgroundPhotoThumbnailLink.(string),
				Moniest: &Moniest{
					Bio:         post.Bio.String,
					Description: post.Description.String,
					CryptoPostStatistics: &CryptoPostStatistics{
						Pnl7days:      post.Pnl7days.Float64,
						Roi7days:      post.Roi7days.Float64,
						WinRate7days:  post.WinRate7days.Float64,
						Pnl30days:     post.Pnl30days.Float64,
						Roi30days:     post.Roi30days.Float64,
						WinRate30days: post.WinRate30days.Float64,
						PnlTotal:      post.PnlTotal.Float64,
						RoiTotal:      post.RoiTotal.Float64,
						WinRateTotal:  post.WinRateTotal.Float64,
					},
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
			Moniest: &Moniest{
				Bio:         user.Bio.String,
				Description: user.Description.String,
				MoniestSubscriptionInfo: &MoniestSubscriptionInfo{
					Fee:       user.Fee,
					Message:   user.Message.String,
					UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt,
				},
				CryptoPostStatistics: &CryptoPostStatistics{
					Pnl7days:      user.Pnl7days.Float64,
					Roi7days:      user.Roi7days.Float64,
					WinRate7days:  user.WinRate7days.Float64,
					Pnl30days:     user.Pnl30days.Float64,
					Roi30days:     user.Roi30days.Float64,
					WinRate30days: user.WinRate30days.Float64,
					PnlTotal:      user.PnlTotal.Float64,
					RoiTotal:      user.RoiTotal.Float64,
					WinRateTotal:  user.WinRateTotal.Float64,
				},
			},
		})
	}

	return response
}

func NewGetContentMoniestResponse(moniests ContentMoniestDBResponse) []UserAsContent {
	response := make([]UserAsContent, 0, len(moniests))

	for _, user := range moniests {
		response = append(response, UserAsContent{
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
			Moniest: &MoniestAsContent{
				Bio:             user.Bio.String,
				Description:     user.Description.String,
				SubscriberCount: user.UserSubscriptionCount,
				MoniestSubscriptionInfo: &MoniestSubscriptionInfo{
					Fee:       user.Fee,
					Message:   user.Message.String,
					UpdatedAt: user.MoniestSubscriptionInfoUpdatedAt,
				},
				CryptoPostStatistics: &CryptoPostStatistics{
					Pnl7days:      user.Pnl7days.Float64,
					Roi7days:      user.Roi7days.Float64,
					WinRate7days:  user.WinRate7days.Float64,
					Pnl30days:     user.Pnl30days.Float64,
					Roi30days:     user.Roi30days.Float64,
					WinRate30days: user.WinRate30days.Float64,
					PnlTotal:      user.PnlTotal.Float64,
					RoiTotal:      user.RoiTotal.Float64,
					WinRateTotal:  user.WinRateTotal.Float64,
				},
			},
		})
	}

	return response
}
