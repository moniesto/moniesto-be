package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
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

	// any user with same email
	checkEmail, err := service.Store.CheckEmail(ctx, validEmail)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CheckEmail"](err))
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorCheckEmail)
	}
	if !checkEmail {
		return db.User{}, clientError.CreateError(http.StatusForbidden, clientError.Account_Register_RegisteredEmail)
	}

	// any user with same username
	checkUsername, err := service.Store.CheckUsername(ctx, registerRequest.Username)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CheckUsername"](err))
		return db.User{}, clientError.CreateError(http.StatusInternalServerError, clientError.Account_Register_ServerErrorCheckUsername)
	}
	if !checkUsername {
		return db.User{}, clientError.CreateError(http.StatusForbidden, clientError.Account_Register_RegisteredUsername)
	}

	// hash password
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		// TODO: add server error
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
		systemError.Log(systemError.InternalMessages["CreateUser"](err))
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
		systemError.Log(systemError.InternalMessages["CheckUsername"](err))
		return false, clientError.CreateError(http.StatusInternalServerError, clientError.Account_CheckUsername_ServerErrorCheckUsername)
	}
	if !checkUsername {
		return false, nil
	}

	return true, nil
}
