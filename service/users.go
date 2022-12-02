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
