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
)

func (service *Service) GetContentPosts(ctx *gin.Context, userID string, subscribed, active bool, sortBy string, limit, offset int) ([]model.GetContentPostResponse, error) {
	// STEP: get subscribed moniest posts
	if subscribed {
		// OPTION 1: subscribed moniest -> active posts
		if active { // active
			postsFromDB, err := service.Store.GetSubscribedActivePosts(ctx, db.GetSubscribedActivePostsParams{
				UserID: userID,
				Limit:  int32(limit),
				Offset: int32(offset),
			})
			if err != nil {
				if err == sql.ErrNoRows {
					return []model.GetContentPostResponse{}, nil
				}
				return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
			}
			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

			return model.NewGetContentPostResponse(posts), nil
		}

		// OPTION 2: subscribed moniest -> deactive(old) posts
		postsFromDB, err := service.Store.GetSubscribedDeactivePosts(ctx, db.GetSubscribedDeactivePostsParams{
			UserID: userID,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return []model.GetContentPostResponse{}, nil
			}
			return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
		}

		posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
		return model.NewGetContentPostResponse(posts), nil
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
		if err == sql.ErrNoRows {
			return []model.GetContentPostResponse{}, nil
		}
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetPosts_ServerErrorGetPosts)
	}
	posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))
	return model.NewGetContentPostResponse(posts), nil
}

func (service *Service) GetContentMoniests(ctx *gin.Context, user_id string, limit, offset int) ([]model.User, error) {
	// STEP: get all moniests -> highest score first
	params := db.GetMoniestsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	moniestFromDB, err := service.Store.GetMoniests(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.User{}, nil
		}

		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_GetMoniests_ServerErrorGetMoniests)
	}

	moniests := *(*model.MoniestDBResponse)(unsafe.Pointer(&moniestFromDB))
	return model.NewGetContentMoniestResponse(moniests), nil
}

func (service *Service) SearchMoniest(ctx *gin.Context, searchText string, limit, offset int) ([]model.User, error) {
	querySearchText := "%" + searchText + "%"

	params := db.SearchMoniestsParams{
		Name:   querySearchText,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	moniestFromDB, err := service.Store.SearchMoniests(ctx, params)
	if err != nil {
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Content_SearchMoniests_ServerErrorSearchMoniest)
	}

	moniests := *(*model.MoniestDBResponse)(unsafe.Pointer(&moniestFromDB))

	return model.NewGetContentMoniestResponse(moniests), nil
}
