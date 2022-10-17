package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) (server *Server) {

	server = &Server{store: store}
	router := gin.Default()

	router.POST("/hello", func(ctx *gin.Context) {
		fmt.Println("Hey hello")
	})

	server.router = router
	return
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
