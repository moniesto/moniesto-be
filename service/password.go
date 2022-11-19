package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CheckPassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: get old password [in hashed form]
	hashedPasword, err := service.Store.GetPasswordByID(ctx, user_id)
	if err != nil {
		systemError.Log(systemError.InternalMessages["GetPassword"](err))
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckPassword)
	}

	// STEP: check the password
	err = util.CheckPassword(password, hashedPasword)
	if err != nil {
		systemError.Log(systemError.InternalMessages["LoginFail"](err))
		return clientError.CreateError(http.StatusForbidden, clientError.Account_ChangePassword_WrongPassword)
	}

	return
}

func (service *Service) UpdatePassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorPassword)
	}

	setPasswordParams := db.SetPasswordParams{
		ID:       user_id,
		Password: hashedPassword,
	}

	// STEP: update/set password
	err = service.Store.SetPassword(ctx, setPasswordParams)
	if err != nil {
		systemError.Log(systemError.InternalMessages["UpdatePassword"](err))
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorUpdatePassword)
	}

	return
}

func (service *Service) CheckEmailExistidy(ctx *gin.Context, email string) (validEmail string, err error) {
	// STEP: validate email
	validEmail, err = validation.Email(email)
	if err != nil {
		return "", clientError.CreateError(http.StatusNotAcceptable, clientError.Account_ChangePassword_InvalidEmail)
	}

	// STEP: get email existidy in the system
	checkEmail, err := service.Store.CheckEmail(ctx, validEmail)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CheckEmail"](err))
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckEmail)
	}

	// STEP: return error if email is not in the system
	if checkEmail {
		return "", clientError.CreateError(http.StatusNotFound, clientError.Account_ChangePassword_NotFoundEmail)
	}
	return validEmail, nil
}

// TODO: update function
func (service *Service) CreatePasswordResetToken(ctx *gin.Context, email string, expiryAt time.Duration) (string, db.PasswordResetToken, error) {
	// STEP: get user by email (need ID)
	user, err := service.Store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusNotFound, clientError.Account_ChangePassword_NotFoundEmail)
		}
		systemError.Log(systemError.InternalMessages["GetUserByEmail"](err))
		return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckEmail)
	}

	// STEP: create password reset token object
	token := token.CreateValidatingToken()

	params := db.CreatePasswordResetTokenParams{
		ID:          core.CreateID(),
		UserID:      user.ID,
		Token:       token,
		TokenExpiry: time.Now().Add(expiryAt),
	}

	// STEP: insert DB
	password_reset_token, err := service.Store.CreatePasswordResetToken(ctx, params)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CreatePasswordResetToken"](err))
		return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCreateToken)
	}

	return user.Name, password_reset_token, nil
}
