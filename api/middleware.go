package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
)

const (
	authorizationHeaderKey = "authorization"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
	}
}
