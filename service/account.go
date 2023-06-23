package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/systemError"
	"github.com/moniesto/moniesto-be/util/validation"
)

// CreateUser creates user in register operation
func (service *Service) CreateUser(ctx *gin.Context, registerRequest model.RegisterRequest) (createdUser db.User, err error) {
	validEmail, err := validation.Email(registerRequest.Email)
	if err != nil {
		return db.User{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Register_InvalidEmail)
	}

	err = validation.Password(registerRequest.Password)
	if err != nil {
		return db.User{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Register_InvalidPassword)
	}

	err = validation.Username(registerRequest.Username)
	if err != nil {
		return db.User{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Register_InvalidUsername)
	}

	err = validation.Name(registerRequest.Name)
	if err != nil {
		return db.User{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Register_InvalidName)
	}

	err = validation.Surname(registerRequest.Surname)
	if err != nil {
		return db.User{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Register_InvalidSurname)
	}

	// any user with same email
	checkEmail, err := service.Store.CheckEmail(ctx, validEmail)
	if err != nil {
		systemError.Log("server error on check email", err.Error())
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorCheckEmail)
	}
	if !checkEmail {
		return db.User{}, clientError.CreateError(http.StatusForbidden, clientError.Account_Register_RegisteredEmail)
	}

	// any user with same username
	checkUsername, err := service.Store.CheckUsername(ctx, registerRequest.Username)
	if err != nil {
		systemError.Log("server error on check username", err.Error())
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorCheckUsername)
	}
	if !checkUsername {
		return db.User{}, clientError.CreateError(http.StatusForbidden, clientError.Account_Register_RegisteredUsername)
	}

	// hash password
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		systemError.Log("hash password error", err.Error())
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorPassword)
	}

	user := db.CreateUserParams{
		ID:       core.CreateID(),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}

	createdUser, err = service.Store.CreateUser(ctx, user)
	if err != nil {
		systemError.Log("server error on creating user", err.Error())
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorCreateUser)
	}

	return
}

// GetOwnUser get the users data with identifier[email or username] and password
func (service *Service) GetOwnUser(ctx *gin.Context, identifier, password string) (createdUser db.LoginUserByEmailRow, err error) {

	// STEP: validation
	identifierIsEmail := true
	_, err = validation.Email(identifier)
	if err != nil {

		err = validation.Username(identifier)
		if err != nil {
			return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_Login_InvalidUsername)
		}
		identifierIsEmail = false
	}

	var user db.LoginUserByEmailRow

	// STEP: fetch user from DB by email or username
	if identifierIsEmail { // identifier is email case
		user, err = service.Store.LoginUserByEmail(ctx, identifier)

		if err != nil {

			if err == sql.ErrNoRows {
				systemError.Log(systemError.InternalMessages["LoginFail"](err))
				return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusNotFound, clientError.Account_Login_NotFoundEmail)
			}

			systemError.Log(systemError.InternalMessages["LoginByEmail"](err))
			return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Login_ServerErrorEmail)
		}

	} else { // identifier is username case
		user1, err := service.Store.LoginUserByUsername(ctx, identifier)

		if err != nil {

			if err == sql.ErrNoRows {
				systemError.Log(systemError.InternalMessages["LoginFail"](err))
				return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusNotFound, clientError.Account_Login_NotFoundUsername)
			}

			systemError.Log(systemError.InternalMessages["LoginByUsername"](err))
			return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Login_ServerErrorUsername)
		}

		// assert object always to db.LoginUserByEmailRow
		user = db.LoginUserByEmailRow(user1)
	}

	// STEP: check the password
	err = util.CheckPassword(password, user.Password)
	if err != nil {
		systemError.Log(systemError.InternalMessages["LoginFail"](err))
		return db.LoginUserByEmailRow{}, clientError.CreateError(http.StatusForbidden, clientError.Account_Login_WrongPassword)
	}

	createdUser = user

	return createdUser, nil
}

// GetUserByID get user by user_id
func (service *Service) GetUserByID(ctx *gin.Context, user_id string) (db.GetUserByIDRow, error) {
	// STEP: get user
	user, err := service.Store.GetUserByID(ctx, user_id)
	if err != nil {

		// STEP: user not found
		if err == sql.ErrNoRows {
			return db.GetUserByIDRow{}, clientError.CreateError(http.StatusNotFound, clientError.Account_GetUser_NotFound)
		}

		systemError.Log("get user by id error", err.Error())
		return db.GetUserByIDRow{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_GetUser_ServerError)
	}

	return user, nil
}

// UpdateLoginStats update the latest login status of the user
func (service *Service) UpdateLoginStats(ctx *gin.Context, user_id string) {
	_, err := service.Store.UpdateLoginStats(ctx, user_id)

	if err != nil {
		systemError.Log(systemError.InternalMessages["UpdateLoginStatsFail"](err))
	}
}

// CheckUsername checks the validity of the username [valid/used]
func (service *Service) CheckUsername(ctx *gin.Context, username string) (bool, error) {
	err := validation.Username(username)
	if err != nil {
		return false, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_CheckUsername_InvalidUsername)
	}

	checkUsername, err := service.Store.CheckUsername(ctx, username)
	if err != nil {
		systemError.Log("server error on check username", err.Error())
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.Account_CheckUsername_ServerErrorCheckUsername)
	}
	if !checkUsername {
		return false, nil
	}

	return true, nil
}

// CreateEmailVerificationToken creates email verification token on DB and delete older ones of user
func (service *Service) CreateEmailVerificationToken(ctx *gin.Context, userID, redirectURL string, expiryAt time.Duration) (db.EmailVerificationToken, error) {

	// STEP: delete older email verification tokens
	err := service.Store.DeleteEmailVerificationTokenByUserID(ctx, userID)
	if err != nil {
		systemError.Log("delete email verification token by user id error", err.Error())
	}

	// STEP: create email verification token object
	plain_token := token.CreateValidatingToken()

	params := db.CreateEmailVerificationTokenParams{
		ID:          core.CreateID(),
		UserID:      userID,
		Token:       plain_token,
		TokenExpiry: time.Now().UTC().Add(expiryAt),
		RedirectUrl: redirectURL,
	}

	// STEP: insert DB
	email_verification_token, err := service.Store.CreateEmailVerificationToken(ctx, params)
	if err != nil {
		systemError.Log("create email verification token error", err.Error())
		return db.EmailVerificationToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorCreateToken)
	}

	// STEP: replace with encoded token (send encoded token)
	encoded_token := token.EncodeValidatingToken(plain_token)
	email_verification_token.Token = encoded_token

	return email_verification_token, nil
}

// GetEmailVerificationToken gets email verification token from the DB [check expiry + token validity]
func (service *Service) GetEmailVerificationToken(ctx *gin.Context, encoded_token string) (db.EmailVerificationToken, error) {

	// STEP: get decoded token
	decoded_token, err := token.GetValidatingToken(encoded_token)
	if err != nil {
		return db.EmailVerificationToken{}, clientError.CreateError(http.StatusNotAcceptable, clientError.Account_EmailVerification_InvalidToken)
	}

	// STEP: get email verification token record
	email_verification_token, err := service.Store.GetEmailVerificationTokenByToken(ctx, decoded_token)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.EmailVerificationToken{}, clientError.CreateError(http.StatusNotFound, clientError.Account_EmailVerification_NotFoundToken)
		}

		systemError.Log("get email verification token error", err.Error())
		return db.EmailVerificationToken{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorGetToken)
	}

	// STEP: token is not expired
	if time.Now().UTC().After(email_verification_token.TokenExpiry) {
		return db.EmailVerificationToken{}, clientError.CreateError(http.StatusForbidden, clientError.Account_EmailVerification_ExpiredToken)
	}

	return email_verification_token, nil
}

// VerifyEmail is verifying email
func (service *Service) VerifyEmail(ctx *gin.Context, user_id string) error {

	// STEP: verify email
	err := service.Store.VerifyEmail(ctx, user_id)
	if err != nil {
		systemError.Log("server error on verify email", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorVerifyEmail)
	}

	return nil
}

// DeleteEmailVerificationToken deletes the token from DB by token
func (service *Service) DeleteEmailVerificationToken(ctx *gin.Context, token string) error {

	err := service.Store.DeleteEmailVerificationTokenByToken(ctx, token)

	if err != nil {
		systemError.Log("delete email verification token error", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_EmailVerification_ServerErrorDeleteToken)
	}

	return nil
}

// ChangeUsername set new username to user [+ check username validity]
func (service *Service) ChangeUsername(ctx *gin.Context, user_id, new_username string) error {

	// STEP: check username is valid
	err := validation.Username(new_username)
	if err != nil {
		return clientError.CreateError(http.StatusNotAcceptable, clientError.Account_ChangeUsername_InvalidUsername)
	}

	params := db.SetUsernameParams{
		ID:       user_id,
		Username: new_username,
	}

	// STEP: update/set new username
	err = service.Store.SetUsername(ctx, params)
	if err != nil {
		systemError.Log("server error on set username", err.Error())
		return clientError.CreateError(http.StatusInternalServerError, clientError.Account_ChangeUsername_ServerErrorChangeUsername)
	}

	return nil
}
