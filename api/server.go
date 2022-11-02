package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
)

// Server serves HTTP requests
type Server struct {
	config     util.Config
	service    *service.Service
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, service *service.Service) (*Server, error) {
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

	accountRouters.POST("/register", server.registerUser)
	accountRouters.POST("/login", server.loginUser)
	accountRouters.GET("/usernames/:username/check", server.checkUsername)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
