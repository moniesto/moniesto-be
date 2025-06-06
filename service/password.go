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
	"github.com/moniesto/moniesto-be/util/system"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CheckPassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: get old password [in hashed form]
	hashedPasword, err := service.Store.GetPasswordByID(ctx, user_id)
	if err != nil {
		system.LogError("server error on getting password", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckPassword)
	}

	// STEP: check the password
	err = util.CheckPassword(password, hashedPasword)
	if err != nil {
		system.LogError("login failed", err.Error())
		return clientError.CreateError(http.StatusForbidden, clientError.Account_ChangePassword_WrongPassword)
	}

	return
}

func (service *Service) UpdatePassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		system.LogError("server error on hashing password")
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorPassword)
	}

	setPasswordParams := db.SetPasswordParams{
		ID:       user_id,
		Password: hashedPassword,
	}

	// STEP: update/set password
	err = service.Store.SetPassword(ctx, setPasswordParams)
	if err != nil {
		system.LogError("server error on updating password", err.Error())
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
		system.LogError("server error on check email", err.Error())
		return "", clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckEmail)
	}

	// STEP: return error if email is not in the system
	if checkEmail {
		return "", clientError.CreateError(http.StatusNotFound, clientError.Account_ChangePassword_NotFoundEmail)
	}
	return validEmail, nil
}

func (service *Service) CreatePasswordResetToken(ctx *gin.Context, email string, expiryAt time.Duration) (string, db.PasswordResetToken, error) {
	// STEP: get user by email (need ID)
	user, err := service.Store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusNotFound, clientError.Account_ChangePassword_NotFoundEmail)
		}

		system.LogError("server error on getting user by email", err.Error())
		return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCheckEmail)
	}

	// STEP: delete older password reset tokens
	err = service.Store.DeletePasswordResetTokenByUserID(ctx, user.ID)
	if err != nil {
		system.LogError("server error on deleting password reset token by user id")
	}

	// STEP: create password reset token object
	plain_token := token.CreateValidatingToken()

	params := db.CreatePasswordResetTokenParams{
		ID:          core.CreateID(),
		UserID:      user.ID,
		Token:       plain_token,
		TokenExpiry: util.Now().Add(expiryAt),
	}

	// STEP: insert DB
	password_reset_token, err := service.Store.CreatePasswordResetToken(ctx, params)
	if err != nil {
		system.LogError("server error on creating password reset token on db", err.Error())
		return "", db.PasswordResetToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorCreateToken)
	}

	// STEP: replace with encoded token (send encoded token)
	encoded_token := token.EncodeValidatingToken(plain_token)
	password_reset_token.Token = encoded_token

	return user.Fullname, password_reset_token, nil
}

func (service *Service) GetPasswordResetToken(ctx *gin.Context, encoded_token string) (db.PasswordResetToken, error) {
	// STEP: get decoded token
	decoded_token, err := token.GetValidatingToken(encoded_token)
	if err != nil {
		return db.PasswordResetToken{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_ChangePassword_InvalidToken)
	}

	// STEP: get password reset token record
	password_reset_token, err := service.Store.GetPasswordResetTokenByToken(ctx, decoded_token)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.PasswordResetToken{}, clientError.CreateError(http.StatusNotFound, clientError.Account_ChangePassword_NotFoundToken)
		}

		system.LogError("server error on getting password reset token by token")
		return db.PasswordResetToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorGetToken)
	}

	// STEP: token is not expired
	if util.Now().After(password_reset_token.TokenExpiry) {
		return db.PasswordResetToken{}, clientError.CreateError(http.StatusForbidden, clientError.Account_ChangePassword_ExpiredToken)
	}

	return password_reset_token, nil
}

func (service *Service) DeletePasswordResetToken(ctx *gin.Context, token string) error {
	err := service.Store.DeletePasswordResetTokenByToken(ctx, token)

	if err != nil {
		system.LogError("server error on deleting password reset token by token")
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangePassword_ServerErrorDeleteToken)
	}

	return nil
}
