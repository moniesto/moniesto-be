package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Error Codes
// @Description Get All Error Codes in the system
// @Tags Assets
// @Accept json
// @Produce json
// @Success 200 {object} clientError.ErrorMessagesType
// @Router /assets/error-codes [get]
func (server *Server) getErrorCodes(ctx *gin.Context) {

	// STEP: get error codes
	errorCodes := server.service.GetErrorCodes()

	ctx.JSON(http.StatusOK, errorCodes)
}
