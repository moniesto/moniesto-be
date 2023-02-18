package service

import "github.com/gin-gonic/gin"

func (service *Service) GetContentPosts(ctx *gin.Context, userID string, subscribed, active bool, limit, offset int) {
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

		} else { // deactive(old)

		}
	} else {
		// get deactive(old) posts of all moniests
	}

}
