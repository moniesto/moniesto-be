package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"

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

	// Account routes
	accountRouters := router.Group("/account")
	{
		// No Need Auth
		accountRouters.POST("/register", server.registerUser)
		accountRouters.POST("/login", server.loginUser)
		accountRouters.GET("/usernames/:username/check", server.checkUsername)

		accountRouters.POST("/password/send_email", server.sendResetPasswordEmail)
		accountRouters.POST("/password/verify_token", server.verifyTokenChangePassword)

		// Need auth
		accountRoutersAuth := accountRouters.Group("/").Use(authMiddleware(server.tokenMaker))
		accountRoutersAuth.PUT("/password", server.changePassword)
		accountRoutersAuth.PATCH("/profile", server.updateProfile)
	}

	// Moniests routes - [need Auth]
	moniestsRouters := router.Group("/moniests").Use(authMiddleware(server.tokenMaker))
	{
		moniestsRouters.POST("/", server.CreateMoniest)
	}

	// User routes
	usersRouters := router.Group("/users").Use(authMiddleware(server.tokenMaker))
	{
		usersRouters.GET("/:username", server.GetUserByUsername)
	}

	// Swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
