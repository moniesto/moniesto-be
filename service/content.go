package service

import (
	"database/sql"
	"net/http"
	"unsafe"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) GetContentPosts(ctx *gin.Context, userID string, subscribed, active bool, sortBy string, limit, offset int) ([]model.GetContentPostResponse, error) {
	// STEP: get subscribed moniest posts
	if subscribed {
		// STEP: get user's is moniest or not status
		userIsMoniest, err := service.CheckUserIsMoniestByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}

		// OPTION 1: subscribed moniest -> active posts
		if active { // active
			params := db.GetSubscribedActivePostsParams{
				UserID: userID,
				Limit:  int32(limit),
				Offset: int32(offset),
			}

			// STEP: if user is moniest, append own posts to the response
			if userIsMoniest {
				postsFromDB, err := service.Store.GetSubscribedActivePostsWithOwn(ctx, db.GetSubscribedActivePostsWithOwnParams(params))
				if err != nil {
					systemError.Log("server error on get subscribed active posts with own", err.Error())
					return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
				}

				posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

				return model.NewGetContentPostResponse(posts), nil
			} else {
				postsFromDB, err := service.Store.GetSubscribedActivePosts(ctx, params)
				if err != nil {
					systemError.Log("server error on get subscribed active posts", err.Error())
					return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
				}

				posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

				return model.NewGetContentPostResponse(posts), nil
			}
		}

		// OPTION 2: subscribed moniest -> deactive(old) posts
		params := db.GetSubscribedDeactivePostsParams{
			UserID: userID,
			Limit:  int32(limit),
			Offset: int32(offset),
		}

		// STEP: if user is moniest, append own posts to the response
		if userIsMoniest {
			postsFromDB, err := service.Store.GetSubscribedDeactivePostsWithOwn(ctx, db.GetSubscribedDeactivePostsWithOwnParams(params))
			if err != nil {
				systemError.Log("server error on get subscribed deactive posts with own", err.Error())
				return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
			}

			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
			return model.NewGetContentPostResponse(posts), nil

		} else {
			postsFromDB, err := service.Store.GetSubscribedDeactivePosts(ctx, params)

			if err != nil {
				systemError.Log("server error on get subscribed deactive posts", err.Error())
				return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
			}

			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
			return model.NewGetContentPostResponse(posts), nil
		}
	}

	// all moniests -> deactive(old) high score posts

	// OPTION 3: sorted by score
	if sortBy == util.POST_FILTER_SCORE {
		postsFromDB, err := service.Store.GetDeactivePostsByScore(ctx, db.GetDeactivePostsByScoreParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return []model.GetContentPostResponse{}, nil
			}
			systemError.Log("server error on get deactive posts by score", err.Error())
			return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
		}
		posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
		return model.NewGetContentPostResponse(posts), nil
	}

	// OPTION 4: sorted by createdAt =>  sortBy == util.POST_FILTER_CREATED_AT
	postsFromDB, err := service.Store.GetDeactivePostsByCreatedAt(ctx, db.GetDeactivePostsByCreatedAtParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		systemError.Log("server error on get deactive posts by created at", err.Error())
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
	}
	posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
	return model.NewGetContentPostResponse(posts), nil
}

func (service *Service) GetContentMoniests(ctx *gin.Context, user_id string, limit, offset int) ([]model.GetContentMoniestResponse, error) {
	// STEP: get all moniests -> highest score first
	params := db.GetMoniestsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	moniestFromDB, err := service.Store.GetMoniests(ctx, params)
	if err != nil {
		systemError.Log("server error on get content moniests", err.Error())
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetMoniests_ServerErrorGetMoniests)
	}

	moniests := *(*model.ContentMoniestDBResponse)(unsafe.Pointer(&moniestFromDB))
	return model.NewGetContentMoniestResponse(moniests), nil
}

func (service *Service) SearchMoniest(ctx *gin.Context, searchText string, limit, offset int) ([]model.User, error) {
	querySearchText := "%" + searchText + "%"

	params := db.SearchMoniestsParams{
		Fullname: querySearchText,
		Limit:    int32(limit),
		Offset:   int32(offset),
	}

	moniestFromDB, err := service.Store.SearchMoniests(ctx, params)
	if err != nil {
		systemError.Log("server error on search moniests", err.Error())
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_SearchMoniests_ServerErrorSearchMoniest)
	}

	moniests := *(*model.MoniestDBResponse)(unsafe.Pointer(&moniestFromDB))

	return model.NewGetMoniestsResponse(moniests), nil
}
