package service

import (
	"database/sql"
	"net/http"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) GetOwnUserByUsername(ctx *gin.Context, username string) (db.GetOwnUserByUsernameRow, error) {

	user, err := service.Store.GetOwnUserByUsername(ctx, username)
	if err != nil {

		if err == sql.ErrNoRows {
			return db.GetOwnUserByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByUsername)
		}

		systemError.Log("server error on getting own user by username", err.Error())
		return db.GetOwnUserByUsernameRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.User_GetUser_ServerErrorGetUser)

	}

	return user, nil
}

func (service *Service) GetUserByUsername(ctx *gin.Context, username string) (db.GetUserByUsernameRow, error) {

	user, err := service.Store.GetUserByUsername(ctx, username)
	if err != nil {

		if err == sql.ErrNoRows {
			return db.GetUserByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByUsername)
		}

		systemError.Log("server error on getting user by username", err.Error())
		return db.GetUserByUsernameRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.User_GetUser_ServerErrorGetUser)

	}

	return user, nil
}

func (service *Service) GetOwnUserByID(ctx *gin.Context, userID string) (db.GetOwnUserByIDRow, error) {

	user, err := service.Store.GetOwnUserByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.GetOwnUserByIDRow{}, clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByID)
		}

		systemError.Log("server error on getting user by id", err.Error())
		return db.GetOwnUserByIDRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorGetUser)
	}

	return user, nil
}

// UpdateUserProfile updates profile fields [name | location]
func (service *Service) UpdateUserProfile(ctx *gin.Context, user_id string, req model.UpdateUserProfileRequest) error {
	difference := false

	// STEP: get user data
	user, err := service.Store.GetOwnUserByID(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByID)
		} else {
			systemError.Log("server error on getting user by id", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetUser)
		}
	}

	// STEP: update the ones that provided on parameters
	var param db.UpdateUserParams = db.UpdateUserParams{
		ID:       user_id,
		Fullname: user.Fullname,
		Location: user.Location,
		Language: user.Language,
	}

	if req.Fullname != "" || req.Location != "" || req.Language != "" {
		difference = true

		if req.Fullname != "" {
			err := validation.Fullname(req.Fullname)
			if err != nil {
				return clientError.CreateError(http.StatusNotAcceptable, clientError.Account_UpdateUserProfile_InvalidFullname)
			}

			param.Fullname = req.Fullname
		}

		if req.Location != "" {
			err := validation.Location(req.Location)
			if err != nil {
				return clientError.CreateError(http.StatusNotAcceptable, clientError.Account_UpdateUserProfile_InvalidLocation)
			}

			param.Location = sql.NullString{
				Valid:  true,
				String: req.Location,
			}
		}

		if req.Language != "" {
			err := validation.Language(string(req.Language))
			if err != nil {
				return clientError.CreateError(http.StatusNotAcceptable, clientError.Account_UpdateUserProfile_UnsupportedLanguage)
			}

			param.Language = req.Language
		}
	}

	// STEP: check if there are new values or not
	if difference {
		// STEP: update user
		err = service.Store.UpdateUser(ctx, param)
		if err != nil {
			systemError.Log("server error on update user", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorUpdateUser)
		}
	}

	return nil
}

// UpdateProfilePhoto create/updates profile photos
func (service *Service) UpdateProfilePhoto(ctx *gin.Context, user_id string, image_base64 string) error {

	if image_base64 == "" {
		return nil
	}

	profilePhotoExists := true

	// STEP: get current one if exists
	profilePhotos, err := service.Store.GetProfilePhoto(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			profilePhotoExists = false
		} else {
			systemError.Log("server error on get profile photo", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetProfilePhoto)
		}
	}

	if profilePhotos.Link == image_base64 { // same url sended from UI
		return nil
	}

	// STEP: upload to storage, get links
	profile_photo_links, err := service.Storage.UploadProfilePhoto(ctx, image_base64)
	if err != nil {
		systemError.Log("server error on upload profile photo", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorUploadProfilePhoto)
	}

	if profilePhotoExists {
		// STEP: update db
		param := db.UpdateProfilePhotoParams{
			UserID:        user_id,
			Link:          profile_photo_links.URL,
			ThumbnailLink: profile_photo_links.ThumbnailURL,
		}

		_, err := service.Store.UpdateProfilePhoto(ctx, param)
		if err != nil {
			systemError.Log("server error on update profile photo", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorUpdateProfilePhoto)
		}
	} else {
		// STEP: insert db
		param := db.AddImageParams{
			ID:            core.CreateID(),
			UserID:        user_id,
			Link:          profile_photo_links.URL,
			ThumbnailLink: profile_photo_links.ThumbnailURL,
			Type:          db.ImageTypeProfilePhoto,
		}

		_, err := service.Store.AddImage(ctx, param)
		if err != nil {
			systemError.Log("server error on add image", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorInsertProfilePhoto)
		}
	}

	return nil
}

// UpdateBackgroundPhoto create/updates background photos
func (service *Service) UpdateBackgroundPhoto(ctx *gin.Context, user_id string, image_base64 string) error {

	if image_base64 == "" {
		return nil
	}

	backgroundPhotoExists := true

	// STEP: get current one if exists
	backgroundPhotos, err := service.Store.GetBackgroundPhoto(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			backgroundPhotoExists = false
		} else {
			systemError.Log("server error on get background photo", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetBackgroundPhoto)
		}
	}

	if backgroundPhotos.Link == image_base64 { // same url sended from UI
		return nil
	}

	// STEP: upload to storage, get links
	background_photo_links, err := service.Storage.UploadBackgroundPhoto(ctx, image_base64)
	if err != nil {
		systemError.Log("server error on upload background photo", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorUploadBackgroundPhoto)
	}

	if backgroundPhotoExists {
		// STEP: update db
		param := db.UpdateBackgroundPhotoParams{
			UserID:        user_id,
			Link:          background_photo_links.URL,
			ThumbnailLink: background_photo_links.ThumbnailURL,
		}

		_, err := service.Store.UpdateBackgroundPhoto(ctx, param)
		if err != nil {
			systemError.Log("server error on update background photo", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorUpdateBackgroundPhoto)
		}
	} else {
		// STEP: insert db
		param := db.AddImageParams{
			ID:            core.CreateID(),
			UserID:        user_id,
			Link:          background_photo_links.URL,
			ThumbnailLink: background_photo_links.ThumbnailURL,
			Type:          db.ImageTypeBackgroundPhoto,
		}

		_, err := service.Store.AddImage(ctx, param)
		if err != nil {
			systemError.Log("server error on add image", err.Error())
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorInsertBackgroundPhoto)
		}
	}

	return nil
}

func (service *Service) GetSubscriptions(ctx *gin.Context, user_id string, limit, offset int) ([]model.User, error) {
	params := db.GetSubscriptionsParams{
		UserID: user_id,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	subscriptions, err := service.Store.GetSubscriptions(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.User{}, nil
		}

		systemError.Log("server error on get subscriptions", err.Error())
		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.User_GetSubscriptions_ServerErrorGetSubscriptions)
	}

	moniests := *(*model.MoniestDBResponse)(unsafe.Pointer(&subscriptions))
	return model.NewGetMoniestsResponse(moniests), nil
}

func (service *Service) GetUserStats(ctx *gin.Context, username string) (model.UserStatResponse, error) {

	stats, err := service.Store.GetUserStatsByUsername(ctx, username)
	if err != nil {
		systemError.Log("server error on get stats by username", err.Error())
		return model.UserStatResponse{}, clientError.CreateError(http.StatusInternalServerError, clientError.User_GetStats_ServerErrorGetStats)
	}

	return model.UserStatResponse{SubscriptionCount: stats}, nil
}
