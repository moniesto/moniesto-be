package model

import db "github.com/moniesto/moniesto-be/db/sqlc"

type MetricsResponse struct {
	UserMetrics    db.UserMetricsRow    `json:"user_metrics"`
	PostMetrics    db.PostMetricsRow    `json:"post_metrics"`
	PaymentMetrics db.PaymentMetricsRow `json:"payment_metrics"`
	PayoutMetrics  db.PayoutMetricsRow  `json:"payout_metrics"`
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
	}
}
