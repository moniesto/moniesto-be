package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/scoring"
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
	score := scoring.CalculateApproxScore(req.Duration, currency_price, req.Target3, req.Direction, service.config)

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

	// STEP: convert image base64's to URL (upload to storage)
	descriptionWithPhoto, err := service.postDescriptionImageReplacer(ctx, description)
	if err != nil {
		return db.PostCryptoDescription{}, clientError.CreateError(http.StatusInternalServerError, clientError.Post_CreatePost_ServerErrorPostPhotoUpload)
	}

	createDescription := db.AddPostDescriptionParams{
		ID:          core.CreateID(),
		PostID:      postID,
		Description: descriptionWithPhoto,
	}

	// STEP: create description
	createdDescription, err := service.Store.AddPostDescription(ctx, createDescription)
	if err != nil {
		return db.PostCryptoDescription{}, clientError.CreateError(http.StatusInternalServerError, clientError.Post_CreatePost_ServerErrorCreateDescription)
	}

	return createdDescription, nil
}

// postDescriptionImageReplacer replace images base64's with URL that upladed to storage
func (service *Service) postDescriptionImageReplacer(ctx *gin.Context, description string) (string, error) {
	var post model.PostDescriptionType

	err := json.Unmarshal([]byte(description), &post)
	if err != nil {
		return "", err
	}

	for i := range post.Blocks {
		if post.Blocks[i].Type == "image" {

			file := post.Blocks[i].Data["file"]

			_, ok := file.(map[string]interface{})
			if !ok {
				return "", nil
			}

			base64Image := file.(map[string]interface{})["url"]

			uploadedPostPhoto, err := service.Storage.UploadPostDescriptionPhoto(ctx, base64Image.(string))
			if err != nil {
				return "", err
			}

			file.(map[string]interface{})["url"] = uploadedPostPhoto.URL

			post.Blocks[i].Data["file"] = file
		}
	}

	descriptionByte, err := json.Marshal(post)
	if err != nil {
		return "", err
	}

	return string(descriptionByte), nil
}

func (service *Service) GetMoniestPosts(ctx *gin.Context, moniest_username string, userIsSubscribed, active bool, limit, offset int) ([]model.GetContentPostResponse, error) {

	// OPTION 0: user is not subscribed, but requesting for `active` posts, causes error
	if !userIsSubscribed && active {
		return nil, clientError.CreateError(http.StatusForbidden, clientError.Moniest_GetMoniestPosts_ForbiddenAccess)
	}

	if userIsSubscribed {

		// OPTION 1: user is subscribed & only `active` posts
		if active {
			params := db.GetMoniestActivePostsByUsernameParams{
				Username: moniest_username,
				Limit:    int32(limit),
				Offset:   int32(offset),
			}

			postsFromDB, err := service.Store.GetMoniestActivePostsByUsername(ctx, params)
			if err != nil {
				return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniestPosts_ServerErrorGetPosts)
			}

			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

			return model.NewGetContentPostResponse(posts), nil

		} else { // OPTION 2: user is subscribed & all posts of moniest
			params := db.GetMoniestAllPostsByUsernameParams{
				Username: moniest_username,
				Limit:    int32(limit),
				Offset:   int32(offset),
			}
			postsFromDB, err := service.Store.GetMoniestAllPostsByUsername(ctx, params)
			if err != nil {
				return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniestPosts_ServerErrorGetPosts)
			}

			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

			return model.NewGetContentPostResponse(posts), nil
		}
	}

	// OPTION 3: user is not subscribed & all `not active` posts
	params := db.GetMoniestDeactivePostsByUsernameParams{
		Username: moniest_username,
		Limit:    int32(limit),
		Offset:   int32(offset),
	}

	postsFromDB, err := service.Store.GetMoniestDeactivePostsByUsername(ctx, params)
	if err != nil {
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniestPosts_ServerErrorGetPosts)
	}

	posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

	return model.NewGetContentPostResponse(posts), nil
}

func (service *Service) GetOwnPosts(ctx *gin.Context, moniest_username string, active bool, limit, offset int) ([]model.GetOwnPostResponse, error) {

	if active {
		params := db.GetOwnActivePostsByUsernameParams{
			Username: moniest_username,
			Limit:    int32(limit),
			Offset:   int32(offset),
		}

		postsFromDB, err := service.Store.GetOwnActivePostsByUsername(ctx, params)
		if err != nil {
			return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniestPosts_ServerErrorGetPosts)
		}

		posts := *(*model.OwnPostDBResponse)(unsafe.Pointer(&postsFromDB))

		return model.NewGetOwnPostResponse(posts), nil
	} else {

		params := db.GetOwnAllPostsByUsernameParams{
			Username: moniest_username,
			Limit:    int32(limit),
			Offset:   int32(offset),
		}

		postsFromDB, err := service.Store.GetOwnAllPostsByUsername(ctx, params)
		if err != nil {
			return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetMoniestPosts_ServerErrorGetPosts)
		}

		posts := *(*model.OwnPostDBResponse)(unsafe.Pointer(&postsFromDB))

		return model.NewGetOwnPostResponse(posts), nil
	}
}
