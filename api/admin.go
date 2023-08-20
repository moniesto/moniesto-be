package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/mailing"
	"github.com/moniesto/moniesto-be/util/validation"
)

// @Summary Update Posts status
// @Description Can update the status of the posts manually
// @Security bearerAuth
// @Tags Admin
// @Success 200
// @Failure 403
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /admin/update_posts_status [post]
func (server *Server) UpdatePostsStatusManual(ctx *gin.Context) {

	// STEP: get user id from token
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: get user
	user, err := server.service.GetOwnUserByID(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(clientError.ParseError(err))
		return
	}

	if validation.UserIsAdmin(user.Email) {
		server.UpdatePostStatus()
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}

func (server *Server) SendEmailTest(ctx *gin.Context) {

	mailing.SendPayoutEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "eyvzov", "Davut Turug", "111111", 1, 12, 110, 10, 99, db.UserLanguageEn)

	// mailing.SendNewPostEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "Davut Turug", "parvin", "BTCUSDT", model.LANGUAGE_TURKISH)

	// mailing.SendPasswordResetEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "token", model.LANGUAGE_TURKISH)

	//	mailing.SendWelcomingEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", model.LANGUAGE_TURKISH)

	// mailing.SendEmailVerificationEmail("parvvazov@gmail.com", server.config, "Parvin Eyvazov", "token1", model.LANGUAGE_TURKISH)

}
