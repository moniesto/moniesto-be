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

// @Summary Get Validation Configs
// @Description Get Validation Configs of system
// @Tags Assets
// @Accept json
// @Produce json
// @Success 200 {object} model.GetValidationConfigsResponse
// @Router /assets/validations [get]
func (server *Server) getValidationConfigs(ctx *gin.Context) {

	validation_configs := server.service.GetValidationConfigs()

	ctx.JSON(http.StatusOK, validation_configs)
}

// @Summary Get General Info Configs
// @Description Get General Information Configs
// @Tags Assets
// @Accept json
// @Produce json
// @Success 200 {object} model.GetGeneralInfoResponse
// @Router /assets/general-info [get]
func (server *Server) getGeneralInfoConfigs(ctx *gin.Context) {

	general_info := server.service.GetGeneralInfoConfigs()

	ctx.JSON(http.StatusOK, general_info)
}

// @Summary Get All Configs
// @Description Get All Configs of system
// @Tags Assets
// @Accept json
// @Produce json
// @Success 200 {object} model.GetConfigsResponse
// @Router /assets/configs [get]
func (server *Server) getConfigs(ctx *gin.Context) {

	configs := server.service.GetConfigs()

	ctx.JSON(http.StatusOK, configs)
}
