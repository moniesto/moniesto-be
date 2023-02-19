package service

import (
	"database/sql"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
)

func (service *Service) GetContentPosts(ctx *gin.Context, userID string, subscribed, active bool, limit, offset int) ([]model.GetContentPostResponse, error) {
	/*
		Steps:
			if subscribed == true:
				if active == true:
					get active posts of subscribed moniests with latest first order
				else:
					get deactive(old) posts of subscribed moniests with latest first order


			if subscribed == false:
				get deactive(old) posts of all moniest
	*/

	// STEP: get subscribed moniest posts
	if subscribed {
		if active { // active
			postsFromDB, err := service.Store.GetSubscribedActivePosts(ctx, userID)
			if err != nil {
				if err == sql.ErrNoRows {
					return []model.GetContentPostResponse{}, nil
				}
				// TODO: return error
			}
			posts := *(*model.PostDBResponse)(unsafe.Pointer(&postsFromDB))

			return model.NewGetContentPostResponse(posts), nil

		} else { // deactive(old)

		}
	} else {
		// get deactive(old) posts of all moniests
	}

	return nil, nil
}
