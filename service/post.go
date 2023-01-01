package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

// CreatePost creates post
func (service *Service) CreatePost(req model.CreatePostRequest, currency model.Currency, moniestID string, ctx *gin.Context) (db.CreatePostRow, error) {
	// STEP: duration is valid
	if time.Now().After(req.Duration) {
		return db.CreatePostRow{}, clientError.CreateError(http.StatusMethodNotAllowed, clientError.Post_CreatePost_InvalidDuration)
	}

	// STEP: currency price is invalid
	currency_price, err := strconv.ParseFloat(currency.Price, 64)
	if err != nil {
		return db.CreatePostRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Post_CreatePost_InvalidCurrencyPrice)
	}

	// STEP: targets are valid
	err = validation.Target(currency_price, req.Target1, req.Target2, req.Target3, db.EntryPosition(req.Direction))
	if err != nil {
		return db.CreatePostRow{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Post_CreatePost_InvalidTargets)
	}

	// STEP: stop is valid
	err = validation.Stop(currency_price, req.Stop, db.EntryPosition(req.Direction))
	if err != nil {
		return db.CreatePostRow{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Post_CreatePost_InvalidStop)
	}

	// STEP: get score
	// TODO: update calculate score
	score := core.CalculateApproxScore()

	// STEP: create post
	createPost := db.CreatePostParams{
		ID:         core.CreateID(),
		MoniestID:  moniestID,
		Currency:   currency.Currency,
		StartPrice: currency_price,
		Duration:   req.Duration,
		Target1:    req.Target1,
		Target2:    req.Target2,
		Target3:    req.Target3,
		Stop:       req.Stop,
		Direction:  db.EntryPosition(req.Direction),
		Score:      score,
	}

	post, err := service.Store.CreatePost(ctx, createPost)
	if err != nil {
		return db.CreatePostRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Post_CreatePost_ServerErrorCreatePost)
	}

	return post, nil
}

// CreatePostDescription creates description for the post
func (service *Service) CreatePostDescription(postID, description string, ctx *gin.Context) (db.PostCryptoDescription, error) {

	createDescription := db.AddPostDescriptionParams{
		ID:          core.CreateID(),
		PostID:      postID,
		Description: description,
	}

	// STEP: create description
	createdDescription, err := service.Store.AddPostDescription(ctx, createDescription)
	if err != nil {
		return db.PostCryptoDescription{}, clientError.CreateError(http.StatusInternalServerError, clientError.Post_CreatePost_ServerErrorCreateDescription)
	}

	return createdDescription, nil
}
