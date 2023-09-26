package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"

	"github.com/robfig/cron/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests
type Server struct {
	config     config.Config
	service    *service.Service
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config config.Config, service *service.Service) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		service:    service,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Account routes
	accountRouters := router.Group("/account")
	{
		// No Need Auth
		accountRouters.POST("/register", server.registerUser)
		accountRouters.POST("/login", server.loginUser)
		accountRouters.GET("/usernames/:username/check", server.checkUsername)

		accountRouters.POST("/password/send_email", server.sendResetPasswordEmail)
		accountRouters.POST("/password/verify_token", server.verifyToken)
		accountRouters.POST("/password/change_password", server.verifyTokenChangePassword)

		accountRouters.POST("/email/verify_email", server.verifyEmail)

		// Need auth
		accountRoutersAuth := accountRouters.Group("/").Use(authMiddleware(server.tokenMaker))
		accountRoutersAuth.PUT("/password", server.changePassword)
		accountRoutersAuth.POST("/email/send_verification_email", server.sendVerificationEmail)
		accountRoutersAuth.PATCH("/username", server.changeUsername)
	}

	// Moniests routes - [need Auth]
	moniestsRouters := router.Group("/moniests").Use(authMiddleware(server.tokenMaker))
	{
		moniestsRouters.POST("", server.createMoniest)
		moniestsRouters.POST("/posts", server.createPost)
		moniestsRouters.POST("/posts/calculate-pnl-roi", server.calculatePnlRoi)
		moniestsRouters.PATCH("/profile", server.updateMoniestProfile)
		moniestsRouters.GET("/payout", server.getMoniestPayoutInfo)
		moniestsRouters.PATCH("/payout", server.updateMoniestPayoutInfo)
		moniestsRouters.POST("/:username/subscribe", server.subscribeMoniest)
		moniestsRouters.GET("/:username/subscription-info", server.getSubscriptionInfo)
		moniestsRouters.POST("/:username/unsubscribe", server.unsubscribeMoniest)
		moniestsRouters.GET("/:username/subscribers", server.getSubscribers)
		moniestsRouters.GET("/:username/posts", server.getMoniestPosts)
	}

	// User routes
	usersRouters := router.Group("/users").Use(authMiddleware(server.tokenMaker))
	{
		usersRouters.PATCH("/profile", server.updateUserProfile)
		usersRouters.GET("/:username/subscriptions", server.getSubscriptions)
		usersRouters.GET("/:username/summary-stats", server.getUserStats)

		// PRIMARY TODO: make this endpoint public instead of auth
		usersRouters.GET("/:username", server.getUserByUsername)
	}

	// Crypto routes
	cryptoRouters := router.Group("/crypto").Use(authMiddleware(server.tokenMaker))
	{
		cryptoRouters.GET("/currencies", server.getCurrencies)
	}

	// Assets routes
	assetRouters := router.Group("/assets")
	{
		assetRouters.GET("/configs", server.getConfigs)
		assetRouters.GET("/error-codes", server.getErrorCodes)
		assetRouters.GET("/validations", server.getValidationConfigs)
		assetRouters.GET("/general-info", server.getGeneralInfoConfigs)
	}

	// Feedback routes
	feedbackRouters := router.Group("/feedback").Use(authMiddlewareOptional(server.tokenMaker))
	{
		feedbackRouters.POST("", server.createFeedback)
	}

	// Content routes
	contentRouters := router.Group("/content").Use(authMiddleware(server.tokenMaker))
	{
		contentRouters.GET("/posts", server.getContentPosts)
		contentRouters.GET("/moniests", server.getContentMoniests)
		contentRouters.GET("/moniests/search", server.searchMoniest)
	}

	// Admin routes
	adminRouters := router.Group("/admin").Use(authMiddleware(server.tokenMaker))
	{
		adminRouters.POST("/update_posts_status", server.ADMIN_UpdatePostsStatusManual)
		adminRouters.POST("/update_moniest_post_crypto_statistics", server.ADMIN_UpdateMoniestPostCryptoStatisticsManual)
		adminRouters.GET("/metrics", server.ADMIN_Metrics)
	}

	// Payment routes
	paymentRouters := router.Group("/payment").Use(authMiddleware(server.tokenMaker))
	{
		// TODO: info endpoint
		paymentRouters.POST("/binance/transactions/check/:transaction_id", server.CheckBinancePaymentTransaction)
	}

	// Webhooks
	webhookRouters := router.Group("/webhooks")
	{
		webhookRouters.POST("/binance/transactions", server.TriggerBinanceTransactionWebhook)
	}

	healthRouters := router.Group("/health")
	{
		healthRouters.GET("/", server.HealthCheck)
	}

	// Swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
}

func (server *Server) setupCRONJobs() {
	job := cron.New()

	// JOB: updating post status
	job.AddFunc(util.JOB_TYPE_EVERY_30TH_MINUTE, server.Analyzer)

	// JOB: updating all moniests' post statistics
	job.AddFunc(util.JOB_TYPE_EVERY_5TH_MINUTE_OF_HOUR, server.UpdateMoniestPostCryptoStatistics)

	// JOB: payout to moniest
	job.AddFunc(util.JOB_TYPE_EVERY_12AM, server.PayoutToMoniest)

	// JOB: looking at the all transactions in pending state, and update if more than 10 minutes
	job.AddFunc(util.JOB_TYPE_EVERY_HOUR, server.DetectExpiredPendingTransaction)

	// JOB: checking ended subscription
	job.AddFunc(util.JOB_TYPE_EVERY_HOUR, server.DetectExpiredActiveSubscriptions)

	job.Start()
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) StartCRONJobService() {
	server.setupCRONJobs()
}

// CORSMiddleware, allow all origins for the initial state
// FUTURE TODO: make origin specific
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
