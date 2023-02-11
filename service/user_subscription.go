package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (service *Service) SubscribeMoniest(ctx *gin.Context, moniestID string, userID string) error {
	// STEP: get subscription status
	exist, subscription, err := service.getUserSubscriptionStatus(ctx, moniestID, userID)
	if err != nil {
		return err
	}

	// STEP: check user is not already subscribed OR in deactive status
	if exist {
		if subscription.Active {
			return clientError.CreateError(http.StatusBadRequest, clientError.Moniest_Subscribe_AlreadySubscribed)
		}
	}

	// STEP: add | update db
	if exist {
		// STEP: update db
		params := db.ActivateSubscriptionParams{
			UserID:    userID,
			MoniestID: moniestID,
		}

		err = service.Store.ActivateSubscription(ctx, params)
		if err != nil {
			return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorActivateSubscription)
		}
	} else {
		// STEP: add to db
		params := db.CreateSubscriptionParams{
			ID:        core.CreateID(),
			UserID:    userID,
			MoniestID: moniestID,
		}

		_, err := service.Store.CreateSubscription(ctx, params)
		if err != nil {
			return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorCreateSubscriptionDB)
		}
	}

	// STEP: payment
	// PAYMENT FUTURE TODO: subscribe to moniest
	// remove from db if payment failed

	return nil
}

func (service *Service) getUserSubscriptionStatus(ctx *gin.Context, moniestID string, userID string) (bool, db.UserSubscription, error) {

	params := db.GetSubscriptionParams{
		UserID:    userID,
		MoniestID: moniestID,
	}

	// STEP: fetch subscription
	subscription, err := service.Store.GetSubscription(ctx, params)
	if err != nil {

		if err == sql.ErrNoRows { // if not exist -> return false as exist
			return false, db.UserSubscription{}, nil
		}

		return false, db.UserSubscription{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorGetSubscription)
	}

	return true, subscription, nil
}
