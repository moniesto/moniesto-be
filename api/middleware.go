package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

const (
	authorizationHeaderKey          = "authorization"
	authorizationTypeBearer         = "bearer"
	authorizationPayloadKey         = "authorization_payload"
	authorizationPayloadValidityKey = "authorization_payload_validity"
)

func authMiddlewareOptional(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.Set(authorizationPayloadValidityKey, false)
			return
		}
		ctx.Set(authorizationPayloadValidityKey, true)
		authMiddleware(token)(ctx)
	}
}

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, clientError.GetError(clientError.Account_Authorization_NotProvidedHeader))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.JSON(http.StatusUnauthorized, clientError.GetError(clientError.Account_Authorization_InvalidHeader))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.JSON(http.StatusUnauthorized, clientError.GetError(clientError.Account_Authorization_UnsupportedType))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, clientError.GetError(clientError.Account_Authorization_InvalidToken))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
