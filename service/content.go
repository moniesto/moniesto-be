package service

import (
	"database/sql"
	"net/http"
	"unsafe"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (service *Service) GetContentPosts(ctx *gin.Context, userID string, subscribed, active bool, limit, offset int) ([]model.GetContentPostResponse, error) {
	/*
		Steps:
			if subscribed == false:
				get deactive(old) posts of all moniest

				tutmus
				en yuksek scorlu postlar


	*/

	// STEP: get subscribed moniest posts
	if subscribed {
		// STEP: subscribed moniest -> active posts
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

		// STEP: subscribed moniest -> deactive(old) posts
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

	// STEP: all moniests -> deactive(old) posts

	return nil, nil
}
