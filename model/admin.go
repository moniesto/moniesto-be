package model

import (
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

type MetricsResponse struct {
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

func NewMetricsResponse(
	userMetrics db.UserMetricsRow,
	postMetrics db.PostMetricsRow,
	paymentMetrics db.PaymentMetricsRow,
	payoutMetrics db.PayoutMetricsRow,
	feedbackMetrics db.FeedbackMetricsRow,
	feedbacks []db.GetFeedbacksRow) MetricsResponse {

	return MetricsResponse{
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
