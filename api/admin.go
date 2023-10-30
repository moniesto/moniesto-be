package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

// @Summary Update Posts status manually
// @Description Can update the status of the posts manually
// @Security bearerAuth
// @Tags Admin
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/update_posts_status [post]
func (server *Server) ADMIN_UpdatePostsStatusManual(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	server.Analyzer()
	ctx.Status(http.StatusOK)
}

// @Summary Update Moniest Post Crypto Statistics manually
// @Description Can update the Moniest Post Crypto Statistics manually
// @Security bearerAuth
// @Tags Admin
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/update_moniest_post_crypto_statistics [post]
func (server *Server) ADMIN_UpdateMoniestPostCryptoStatisticsManual(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	server.UpdateMoniestPostCryptoStatistics()
	ctx.Status(http.StatusOK)
}

// @Summary Get Metrics
// @Description Get Metrics of user, post, payments, payouts, feedbacks
// @Security bearerAuth
// @Tags Admin
// @Success 200 {object} model.MetricsResponse
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/metrics [get]
func (server *Server) ADMIN_Metrics(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	metrics, err := server.service.Metrics(ctx)
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

	runnerTypes := map[string]func(){
		"post-analyzer":                       server.Analyzer,
		"moniest-analyzer":                    server.UpdateMoniestPostCryptoStatistics,
		"payout":                              server.PayoutToMoniest,
		"detect-expired-pending-transaction":  server.DetectExpiredPendingTransaction,
		"detect-expired-active-subscriptions": server.DetectExpiredActiveSubscriptions,
	}

	// STEP: get value based on key
	runner, exists := runnerTypes[runnerType]
	if !exists {
		allTypes := []string{}
		for k := range runnerTypes {
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

func (server *Server) ADMIN_Test(ctx *gin.Context) {
}

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
