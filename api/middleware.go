package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/error"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(error.Messages["NotProvided_AuthorizationHeader"]())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(error.Messages["Invalid_AuthorizationHeader"]())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(error.Messages["Unsupported_AuthorizationType"](authorizationType))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(error.Messages["Invalid_Token"]())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
