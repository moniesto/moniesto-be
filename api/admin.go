package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/system"
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
// @Param runner path string true "runner type"
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

// @Summary Operation: Subscribe
// @Description Make subscribtion to any moniesto for any user
// @Security bearerAuth
// @Tags Admin
// @Param username path string true "user username"
// @Param moniest_username path string true "moniest username"
// @Success 200
// @Failure 400 {object} clientError.ErrorResponse "already subscribed"
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 404 {object} clientError.ErrorResponse "user or moniest not found"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/operations/:username/subscribe/:moniest_username [post]
func (server *Server) ADMIN_OPERATIONS_Subscribe(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	// STEP: get params [username, moniest_username]
	username := ctx.Param("username")
	moniestUsername := ctx.Param("moniest_username")

	user, err := server.service.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
	}

	moniest, err := server.service.GetMoniestByUsername(ctx, moniestUsername)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
	}

	latestTransactionID := ""
	subscriptionStartDate := time.Now()
	subscriptionEndDate := subscriptionStartDate.AddDate(2, 0, 0)

	err = server.service.SubscribeMoniest(
		ctx,
		moniest.MoniestID,
		user.ID,
		latestTransactionID,
		subscriptionStartDate,
		subscriptionEndDate,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Operation: Create Post
// @Description or `admin/operations/:username/create-post` for specific moniest
// @Security bearerAuth
// @Tags Admin
// @Param username path string true "user username"
// @Success 200
// @Failure 403 {object} clientError.ErrorResponse "not admin"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/operations/:username/create-post [post]
// @Router /admin/operations/create-post [post]
func (server *Server) ADMIN_OPERATIONS_CreatePost(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	// STEP: get params [username, moniest_username]
	username := ctx.Param("username")

	server.service.ADMIN_CreatePost(username)
}

func (server *Server) ADMIN_OPERATIONS_BeMoniest(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	var req model.CreateMoniestRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Moniest_CreateMoniest_InvalidBody))
		return
	}

	// STEP: get params [username, moniest_username]
	username := ctx.Param("username")

	user, err := server.service.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: check user is already moniest or not
	userIsMoniest, err := server.service.CheckUserIsMoniestByUserID(ctx, user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}
	if userIsMoniest {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, clientError.GetError(clientError.Moniest_CreateMoniest_UserIsAlreadyMoniest))
		return
	}

	// STEP: verify user email
	err = server.service.VerifyEmail(ctx, user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create moniest
	moniest, err := server.service.CreateMoniest(ctx, user.ID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create subscription info
	_, err = server.service.CreateSubsriptionInfo(ctx, moniest.ID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create payout info
	_, err = server.service.CreatePayoutInfo(ctx, moniest.ID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: create moniest post crypto statistics
	_, err = server.service.CreateMoniestPostCryptoStatistics(ctx, moniest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: get created moniest data [+ user data]
	createdMoniest, err := server.service.GetMoniestByMoniestID(ctx, moniest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	// STEP: update data form
	response := model.NewCreateMoniestResponse(createdMoniest)

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) ADMIN_OPERATIONS_ShareTwitterPost(ctx *gin.Context) {
	if !server.isAdmin(ctx) {
		return
	}

	// STEP: get params [username, moniest_username]
	postID := ctx.Param("post_id")

	fmt.Println("post ID", postID)

	err := server.service.ShareTwitterPost(ctx, postID)
	if err != nil {
		system.LogError("share twitter post error", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
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
