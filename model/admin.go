package model

import db "github.com/moniesto/moniesto-be/db/sqlc"

type MetricsResponse struct {
	UserMetrics      db.UserMetricsRow    `json:"user_metrics"`
	PostMetrics      db.PostMetricsRow    `json:"post_metrics"`
	PaymentMetrics   db.PaymentMetricsRow `json:"payment_metrics"`
	PayoutMetrics    db.PayoutMetricsRow  `json:"payout_metrics"`
	FinancialMetrics FinancialMetrics     `json:"financial_metrics"`
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
	payoutMetrics db.PayoutMetricsRow) MetricsResponse {

	return MetricsResponse{
		UserMetrics:    userMetrics,
		PostMetrics:    postMetrics,
		PaymentMetrics: paymentMetrics,
		PayoutMetrics:  payoutMetrics,
		FinancialMetrics: FinancialMetrics{
			Payments: paymentMetrics.SuccessPaymentAmount,
			Profit: (payoutMetrics.SuccessPayoutsAmount - payoutMetrics.SuccessPayoutsAmountAfterCut) +
				(payoutMetrics.RefundPayoutsAmount - payoutMetrics.RefundPayoutsAmountAfterCut),
			EstimatedProfit: (payoutMetrics.SuccessPayoutsAmount - payoutMetrics.SuccessPayoutsAmountAfterCut) +
				(payoutMetrics.RefundPayoutsAmount - payoutMetrics.RefundPayoutsAmountAfterCut) + (payoutMetrics.PendingPayoutsAmount - payoutMetrics.PendingPayoutsAmountAfterCut),
		},
	}
}
