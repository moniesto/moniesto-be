package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (server *Server) GetCurrencies(ctx *gin.Context) {
	var req model.GetCurrenciesRequest

	// STEP: bind/validation
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Crypto_GetCurrencies_InvalidParam))
		return
	}

	// STEP: get currencies with name
	currencies, err := server.service.GetCurrenciesWithName(req.Name)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	ctx.JSON(http.StatusOK, currencies)
}
