package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

// @Summary Get Metrics
// @Description Get Metrics of user, post, payments, payouts, feedbacks
// @Security bearerAuth
// @Tags Admin
// @Success 200 {object} model.ADMIN_MetricsResponse
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/metrics [get]
func (server *Server) ADMIN_Metrics(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	metrics, err := server.service.ADMIN_Metrics(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, metrics)
}

// @Summary Runner
// @Description Running key operations [`post-analyzer`, `moniest-analyzer`, `payout`, `detect-expired-pending-transaction`, `detect-expired-active-subscriptions`]
// @Security bearerAuth
// @Tags Admin
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 404 "runner type not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/run/:runner [post]
func (server *Server) ADMIN_Runner(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	// STEP: get username from param
	runnerType := ctx.Param("runner")

	runners := map[string]func(){
		"post-analyzer":                       server.Analyzer,
		"moniest-analyzer":                    server.UpdateMoniestPostCryptoStatistics,
		"payout":                              server.PayoutToMoniest,
		"detect-expired-pending-transaction":  server.DetectExpiredPendingTransaction,
		"detect-expired-active-subscriptions": server.DetectExpiredActiveSubscriptions,
	}

	// STEP: get value based on key
	runner, exists := runners[runnerType]
	if !exists {
		allTypes := []string{}
		for k := range runners {
			allTypes = append(allTypes, k)
		}

		ctx.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("invalid runner type '%s'. Available runner types: %s", runnerType, strings.Join(allTypes, ", ")))
		return
	}

	runner()
	ctx.Status(http.StatusOK)
}

func (server *Server) SendEmailTest(ctx *gin.Context) {

	// mailing.SendPayoutEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "eyvzov", "Davut Turug", "111111", 1, 12, 110, 10, 99, db.UserLanguageEn)

	// mailing.SendNewPostEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "Davut Turug", "parvin", "BTCUSDT", model.LANGUAGE_TURKISH)

	// mailing.SendPasswordResetEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "token", model.LANGUAGE_TURKISH)

	//	mailing.SendWelcomingEmail("1justingame@gmail.com", server.config, "Parvin Eyvazov", db.UserLanguageTr)

	// mailing.SendEmailVerificationEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "token1", model.LANGUAGE_TURKISH)
}

// @Summary Get Data
// @Description Data types [`users`]
// @Security bearerAuth
// @Tags Admin
// @Produce json
// @Param limit query int false "default: 10 & max: 50"
// @Param offset query int false "default: 0"
// @Success 200 {object} []model.OwnUser "response for data_type: users"
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 404 "data type not found"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/data/:data_type [get]
func (server *Server) ADMIN_Data(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	dataFetchers := map[string]model.AdminDataFetcherFunc{
		"users": server.service.ADMIN_DataUser,
	}

	// STEP: get fetcher based on data_type
	dataType := ctx.Param("data_type")

	fetcher, exists := dataFetchers[dataType]
	if !exists {
		allTypes := []string{}
		for k := range dataFetchers {
			allTypes = append(allTypes, k)
		}

		ctx.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("invalid runner type '%s'. Available runner types: %s", dataType, strings.Join(allTypes, ", ")))
		return
	}

	// STEP: get limit and offset
	var req model.ADMIN_DataRequest = model.ADMIN_DataRequest{
		Limit:  util.DEFAULT_LIMIT,
		Offset: util.DEFAULT_OFFSET,
	}

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Admin_GetData_InvalidParam))
		return
	}

	// STEP: safe limit & offset
	req.Limit = util.SafeLimit(req.Limit)
	req.Offset = util.SafeOffset(req.Offset)

	// STEP: get data
	data, err := fetcher(ctx, req.Limit, req.Offset)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (server *Server) ADMIN_Test(ctx *gin.Context) {
}

// helper functions
func (server *Server) isAdmin(ctx *gin.Context) bool {
	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: get user
	user, err := server.service.GetOwnUserByID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return false
	}

	if validation.UserIsAdmin(user.Email) {
		return true
	}

	ctx.AbortWithStatusJSON(clientError.ParseError(clientError.CreateError(http.StatusForbidden, clientError.General_Not_Admin)))
	return false
}
