package service

import (
	"database/sql"
	"net/http"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) SubscribeMoniest(ctx *gin.Context, moniestID string, userID string, latestTransactionID string, numberOfMonths int) error {
	// STEP: get subscription status
	exist, subscription, err := service.GetUserSubscriptionStatus(ctx, moniestID, userID)
	if err != nil {
		return err
	}

	// STEP: check user is not already subscribed OR in deactive status
	if exist {
		if subscription.Active {
			return clientError.CreateError(http.StatusBadRequest, clientError.Moniest_Subscribe_AlreadySubscribed)
		}
	}

	subscriptionStartDate := time.Now().UTC()
	subscriptionEndDate := subscriptionStartDate.AddDate(0, numberOfMonths, 0)

	// STEP: add | update db
	if exist {
		// STEP: update db
		params := db.ActivateSubscriptionParams{
			UserID:                userID,
			MoniestID:             moniestID,
			LatestTransactionID:   sql.NullString{Valid: true, String: latestTransactionID},
			SubscriptionStartDate: subscriptionStartDate,
			SubscriptionEndDate:   subscriptionEndDate,
		}

		err = service.Store.ActivateSubscription(ctx, params)
		if err != nil {
			systemError.Log("server error on activate subscription", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorActivateSubscription)
		}
	} else {
		// STEP: add to db
		params := db.CreateSubscriptionParams{
			ID:                    core.CreateID(),
			UserID:                userID,
			MoniestID:             moniestID,
			LatestTransactionID:   sql.NullString{Valid: true, String: latestTransactionID},
			SubscriptionStartDate: subscriptionStartDate,
			SubscriptionEndDate:   subscriptionEndDate,
		}

		_, err := service.Store.CreateSubscription(ctx, params)
		if err != nil {
			systemError.Log("server error on create subscription", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorCreateSubscriptionDB)
		}
	}

	// STEP: payment

	// TODO: start payout operation
	service.StartPayout() // TODO: update

	return nil
}

func (service *Service) UnsubscribeMoniest(ctx *gin.Context, moniestID string, userID string) error {
	// STEP: get subscription status
	exist, subscription, err := service.GetUserSubscriptionStatus(ctx, moniestID, userID)
	if err != nil {
		return err
	}

	// STEP: check user is not already unsubscribed OR in deactive status
	if exist {
		if !subscription.Active {
			return clientError.CreateError(http.StatusBadRequest, clientError.Moniest_Unsubscribe_NotSubscribed)
		}
	} else {
		return clientError.CreateError(http.StatusBadRequest, clientError.Moniest_Unsubscribe_NotSubscribed)
	}

	// STEP: deactivate subscription
	params := db.EndsubscriptionParams{
		UserID:    userID,
		MoniestID: moniestID,
	}

	err = service.Store.Endsubscription(ctx, params)
	if err != nil {
		systemError.Log("server error on end subscription", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Unsubscribe_ServerErrorUnsubscribe)
	}

	// STEP: stop subscription on payment
	// PAYMENT FUTURE TODO: unsubscribe from moniest
	// add db if payment failed

	return nil
}

func (service *Service) GetUserSubscriptionStatus(ctx *gin.Context, moniestID string, userID string) (bool, db.UserSubscription, error) {

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

		systemError.Log("server error on get subscription", err.Error())
		return false, db.UserSubscription{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_Subscribe_ServerErrorGetSubscription)
	}

	return true, subscription, nil
}

func (service *Service) CheckUserSubscriptionByMoniestUsername(ctx *gin.Context, user_id, moniest_username string) (bool, error) {

	// STEP: check user is subscribed to moniest
	param := db.CheckSubscriptionByMoniestUsernameParams{
		UserID:   user_id,
		Username: moniest_username,
	}

	userIsSubscribed, err := service.Store.CheckSubscriptionByMoniestUsername(ctx, param)
	if err != nil {
		systemError.Log("server error on check subscription by moniest username", err.Error())
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_SubscribeCheck_ServerErrorCheck)
	}

	return userIsSubscribed, nil
}

func (service *Service) GetSubscribers(ctx *gin.Context, moniestID string, limit, offset int) ([]model.User, error) {

	// STEP: get subscribers
	param := db.GetSubscribersParams{
		MoniestID: moniestID,
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	usersFromDB, err := service.Store.GetSubscribers(ctx, param)
	if err != nil {
		systemError.Log("server error on get subscribers", err.Error())
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_GetSubscriber_ServerErrorGetSubscribers)
	}

	users := *(*model.UserDBResponse)(unsafe.Pointer(&usersFromDB))
	return model.NewGetUsersResponse(users), nil
}
