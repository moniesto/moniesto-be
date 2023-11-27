package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) ADMIN_Metrics(ctx *gin.Context) (model.ADMIN_MetricsResponse, error) {
	userMetrics, err := service.Store.UserMetrics(ctx)
	if err != nil || len(userMetrics) == 0 {
		system.LogError("user metrics error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorUserMetrics)
	}

	postMetrics, err := service.Store.PostMetrics(ctx)
	if err != nil || len(postMetrics) == 0 {
		system.LogError("post metrics error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPostMetrics)
	}

	paymentMetrics, err := service.Store.PaymentMetrics(ctx)
	if err != nil || len(paymentMetrics) == 0 {
		system.LogError("payment metrics error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPaymentMetrics)
	}

	payoutMetrics, err := service.Store.PayoutMetrics(ctx, sql.NullFloat64{Valid: true, Float64: service.config.OperationFeePercentage})
	if err != nil || len(payoutMetrics) == 0 {
		system.LogError("payout metrics error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorPayoutMetrics)
	}

	feedbackMetrics, err := service.Store.FeedbackMetrics(ctx)
	if err != nil || len(feedbackMetrics) == 0 {
		system.LogError("feedback metrics error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorFeedbackMetrics)
	}

	feedbacks, err := service.Store.GetFeedbacks(ctx)
	if err != nil {
		system.LogError("get feedbacks error", err.Error())
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetMetrics_ServerErrorGetFeedbacks)
	}

	return model.NewADMIN_MetricsResponse(userMetrics[0], postMetrics[0], paymentMetrics[0], payoutMetrics[0], feedbackMetrics[0], feedbacks), nil
}

func (service *Service) ADMIN_DataUser(ctx *gin.Context, limit, offset int) (any, error) {
	params := db.ADMIN_GetAllUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	usersFromDB, err := service.Store.ADMIN_GetAllUsers(ctx, params)
	if err != nil {
		return model.ADMIN_MetricsResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.Admin_GetData_ServerErrorGetData)
	}

	users := model.NewADMIN_DataUserResponse(usersFromDB)

	return users, nil
}

func (service *Service) ADMIN_CreatePost(username string) {
	err := service.CreateRandomPost(username)
	if err != nil {
		system.LogError(err.Error())
	}
}
