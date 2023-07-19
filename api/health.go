package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary HealthCheck
// @Description Health Check
// @Tags Health
// @Success 200
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /health [get]
func (server *Server) HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
