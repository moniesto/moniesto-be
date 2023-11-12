package model

import (
	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type AdminDataFetcherFunc func(ctx *gin.Context, limit, offset int) (any, error)

type ADMIN_MetricsResponse struct {
	UserMetrics      db.UserMetricsRow    `json:"user_metrics"`
	PostMetrics      db.PostMetricsRow    `json:"post_metrics"`
	PaymentMetrics   db.PaymentMetricsRow `json:"payment_metrics"`
	PayoutMetrics    db.PayoutMetricsRow  `json:"payout_metrics"`
	Feedback         Feedback             `json:"feedback"`
	FinancialMetrics FinancialMetrics     `json:"financial_metrics"`
}

type Feedback struct {
	Metrics   db.FeedbackMetricsRow `json:"metrics"`
	Feedbacks []db.GetFeedbacksRow  `json:"feedbacks"`
}

type FinancialMetrics struct {
	Payments        float64 `json:"payments"`
	Profit          float64 `json:"profit"`
	EstimatedProfit float64 `json:"estimated_profit"`
}

type ADMIN_DataRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

func NewADMIN_MetricsResponse(
	userMetrics db.UserMetricsRow,
	postMetrics db.PostMetricsRow,
	paymentMetrics db.PaymentMetricsRow,
	payoutMetrics db.PayoutMetricsRow,
	feedbackMetrics db.FeedbackMetricsRow,
	feedbacks []db.GetFeedbacksRow) ADMIN_MetricsResponse {

	return ADMIN_MetricsResponse{
		UserMetrics:    userMetrics,
		PostMetrics:    postMetrics,
		PaymentMetrics: paymentMetrics,
		PayoutMetrics:  payoutMetrics,
		Feedback: Feedback{
			Metrics:   feedbackMetrics,
			Feedbacks: feedbacks,
		},
		FinancialMetrics: FinancialMetrics{
			Payments: paymentMetrics.SuccessPaymentAmount,
			Profit: (payoutMetrics.SuccessPayoutsAmount - payoutMetrics.SuccessPayoutsAmountAfterCut) +
				(payoutMetrics.RefundPayoutsAmount - payoutMetrics.RefundPayoutsAmountAfterCut),
			EstimatedProfit: (payoutMetrics.SuccessPayoutsAmount - payoutMetrics.SuccessPayoutsAmountAfterCut) +
				(payoutMetrics.RefundPayoutsAmount - payoutMetrics.RefundPayoutsAmountAfterCut) + (payoutMetrics.PendingPayoutsAmount - payoutMetrics.PendingPayoutsAmountAfterCut),
		},
	}
}

func NewADMIN_DataUserResponse(users []db.ADMIN_GetAllUsersRow) []OwnUser {
	responses := make([]OwnUser, 0, len(users))

	for _, user := range users {
		response := OwnUser{
			Id:                           user.ID,
			Fullname:                     user.Fullname,
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
			moniest := &Moniest{
				Bio:         user.Bio.String,
				Description: user.Description.String,
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
