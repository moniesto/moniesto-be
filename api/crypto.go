package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// @Summary Crypto Currency Search
// @Description Search crypto currencies by name
// @Security bearerAuth
// @Tags Crypto
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Success 200 {object} []model.Currency
// @Failure 406 {object} clientError.ErrorResponse "invalid name"
// @Failure 500 {object} clientError.ErrorResponse "server error & crypto api error"
// @Router /crypto/currencies [get]
func (server *Server) getCurrencies(ctx *gin.Context) {
	var req model.GetCurrenciesRequest

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, clientError.GetError(clientError.Crypto_GetCurrencies_InvalidParam))
		return
	}

	// STEP: get currencies with name
	currencies, err := server.service.GetCurrenciesWithName(req.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, currencies)
}
