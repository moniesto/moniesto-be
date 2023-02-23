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
)

func (service *Service) GetOwnUserByUsername(ctx *gin.Context, username string) (db.GetOwnUserByUsernameRow, error) {

	user, err := service.Store.GetOwnUserByUsername(ctx, username)
	if err != nil {

		if err == sql.ErrNoRows {
			return db.GetOwnUserByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByUsername)
		}

		// TODO: log server error
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

		// TODO: log server error
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

		return db.GetOwnUserByIDRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorGetUser)
	}

	return user, nil
}

// UpdateUserProfile updates profile fields [name | surname | location]
func (service *Service) UpdateUserProfile(ctx *gin.Context, user_id string, req model.UpdateUserProfileRequest) error {
	difference := false

	// STEP: get user data
	user, err := service.Store.GetUserByID(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return clientError.CreateError(http.StatusNotFound, clientError.General_UserNotFoundByID)
		} else {
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetUser)
		}
	}

	// STEP: update the ones that provided on parameters
	var param db.UpdateUserParams = db.UpdateUserParams{
		ID:       user_id,
		Name:     user.Name,
		Surname:  user.Surname,
		Location: user.Location,
	}

	if req.Name != "" || req.Surname != "" || req.Location != "" {
		difference = true

		if req.Name != "" {
			param.Name = req.Name
		}

		if req.Surname != "" {
			param.Surname = req.Surname
		}

		if req.Location != "" {
			param.Location = sql.NullString{
				Valid:  true,
				String: req.Location,
			}
		}
	}

	// STEP: check if there are new values or not
	if difference {
		// STEP: update user
		err = service.Store.UpdateUser(ctx, param)
		if err != nil {
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
	_, err := service.Store.GetProfilePhoto(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			profilePhotoExists = false
		} else {
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetProfilePhoto)
		}
	}

	// STEP: upload to storage, get links
	profile_photo_links, err := service.Storage.UploadProfilePhoto(ctx, image_base64)
	if err != nil {
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
	_, err := service.Store.GetBackgroundPhoto(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			backgroundPhotoExists = false
		} else {
			return clientError.CreateError(http.StatusInternalServerError, clientError.Account_UpdateUserProfile_ServerErrorGetBackgroundPhoto)
		}
	}

	// STEP: upload to storage, get links
	background_photo_links, err := service.Storage.UploadBackgroundPhoto(ctx, image_base64)
	if err != nil {
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

		return nil, clientError.CreateError(http.StatusInternalServerError, clientError.User_GetSubscriptions_ServerErrorGetSubscriptions)
	}

	moniests := *(*model.MoniestDBResponse)(unsafe.Pointer(&subscriptions))
	return model.NewGetContentMoniestResponse(moniests), nil
}
