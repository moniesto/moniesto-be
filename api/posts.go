package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util/clientError"
)

// PRIMARY TODO: make post will have status field [pending, failed, success]
// @Summary Create Post
// @Description Create Crypto Post
// @Security bearerAuth
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostBody body model.CreatePostRequest true "`description` is optional"
// @Success 200 {object} model.CreatePostResponse
// @Failure 400 {object} clientError.ErrorResponse "user is not moniest"
// @Failure 406 {object} clientError.ErrorResponse "invalid body"
// @Failure 500 {object} clientError.ErrorResponse "server error"
// @Router /moniests/posts [post]
func (server *Server) CreatePost(ctx *gin.Context) {
	var req model.CreatePostRequest

	// STEP: bind/validation
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotAcceptable, clientError.GetError(clientError.Post_CreatePost_InvalidBody))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user_id := authPayload.User.ID

	// STEP: user is moniest
	userIsMoniest, moniest, err := server.service.UserIsMoniest(ctx, user_id)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}
	if !userIsMoniest {
		ctx.JSON(http.StatusBadRequest, clientError.GetError(clientError.UserNotMoniest))
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

	// STEP: create post
	post, err := server.service.CreatePost(req, currencies[0], moniest.MoniestID, ctx)
	if err != nil {
		ctx.JSON(clientError.ParseError(err))
		return
	}

	var response model.CreatePostResponse

	if req.Description != "" {
		// STEP: create post description
		description, err := server.service.CreatePostDescription(post.ID, req.Description, ctx)
		if err != nil {
			ctx.JSON(clientError.ParseError(err))
			return
		}

		response = model.NewCreatePostResponse(post, description)
	} else {
		response = model.NewCreatePostResponse(post, db.PostCryptoDescription{})
	}

	ctx.JSON(http.StatusOK, response)
}
