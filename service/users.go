package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (service *Service) GetOwnUserByUsername(ctx *gin.Context, username string) (db.GetOwnUserByUsernameRow, error) {

	user, err := service.Store.GetOwnUserByUsername(ctx, username)
	if err != nil {

		if err == sql.ErrNoRows {
			return db.GetOwnUserByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.User_GetUser_NotFoundUser)
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
			return db.GetUserByUsernameRow{}, clientError.CreateError(http.StatusNotFound, clientError.User_GetUser_NotFoundUser)
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
			return db.GetOwnUserByIDRow{}, clientError.CreateError(http.StatusNotFound, clientError.UserNotFoundByID)
		}

		return db.GetOwnUserByIDRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorGetUser)
	}

	return user, nil
}
