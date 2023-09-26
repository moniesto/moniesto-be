package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (service *Service) Metrics(ctx *gin.Context) (model.MetricsResponse, error) {
	userMetrics, err := service.Store.UserMetrics(ctx)
	if err != nil || len(userMetrics) == 0 {
		return model.MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorUserMetrics)
	}

	postMetrics, err := service.Store.PostMetrics(ctx)
	if err != nil || len(postMetrics) == 0 {
		return model.MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPostMetrics)
	}

	paymentMetrics, err := service.Store.PaymentMetrics(ctx)
	if err != nil || len(paymentMetrics) == 0 {
		return model.MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPaymentMetrics)
	}

	payoutMetrics, err := service.Store.PayoutMetrics(ctx)
	if err != nil || len(payoutMetrics) == 0 {
		return model.MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPayoutMetrics)
	}

	return model.NewMetricsResponse(userMetrics[0], postMetrics[0], paymentMetrics[0], payoutMetrics[0]), nil
}
