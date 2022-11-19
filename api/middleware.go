package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/systemError"
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
			// TODO: update with new error handler
			ctx.AbortWithStatusJSON(systemError.Messages["NotProvided_AuthorizationHeader"]())
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			// TODO: update with new error handler
			ctx.AbortWithStatusJSON(systemError.Messages["Invalid_AuthorizationHeader"]())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			// TODO: update with new error handler
			ctx.AbortWithStatusJSON(systemError.Messages["Unsupported_AuthorizationType"](authorizationType))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			// TODO: update with new error handler
			ctx.AbortWithStatusJSON(systemError.Messages["Invalid_Token"]())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
