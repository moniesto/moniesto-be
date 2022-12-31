package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (server *Server) CreatePost(ctx *gin.Context) {
	/*
		STEPS
			user is moniest or not
			currency is right
			duration is not in past
			targets & stop are right (if short, targets has to be lower, stop has to be upper, if long vice-versa)
			calculate score
			insert post
			if description is not null, insert description
	*/

	var req model.CreatePostRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidBody))
		return
	}

	// STEP: get currency
	currencies, err := server.service.GetCurrenciesWithName(req.Currency)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	if len(currencies) != 1 { // not found OR found more than 1
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidCurrency))
		return
	}

}
