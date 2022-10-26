package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/systemError"
	"github.com/moniesto/moniesto-be/util/validation"
)

// CreateUser creates user in register operation
func (service *Service) CreateUser(ctx *gin.Context, registerRequest model.RegisterRequest) (createdUser db.User, err error) {
	validEmail, err := validation.Email(registerRequest.Email)
	if err != nil {
		return
	}
	_ = validEmail

	err = validation.Password(registerRequest.Password)
	if err != nil {
		return
	}

	err = validation.Username(registerRequest.Username)
	if err != nil {
		return
	}

	// any user with same email
	checkEmail, err := service.store.CheckEmail(ctx, validEmail)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CheckEmail"](err))
		return createdUser, errors.New("server error on check email")
	}
	if !checkEmail {
		return createdUser, errors.New("this email is already in use")
	}

	// any user with same username
	checkUsername, err := service.store.CheckUsername(ctx, registerRequest.Username)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CheckUsername"](err))
		return createdUser, errors.New("server error on check username")
	}
	if !checkUsername {
		return createdUser, errors.New("this username is already in use")
	}

	// hash password
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		return
	}

	user := db.CreateUserParams{
		ID:       util.CreateID(),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}

	fmt.Println("user", user)

	createdUser, err = service.store.CreateUser(ctx, user)
	if err != nil {
		systemError.Log(systemError.InternalMessages["CreateUser"](err))
		return createdUser, errors.New("server error on create user")
	}

	return
}

// GetOwnUser get the users data with identifier[email or username] and password
func (service *Service) GetOwnUser(ctx *gin.Context, identifier, password string) (createdUser db.User, err error) {
	identifierIsEmail := true
	validEmail, err := validation.Email(identifier)
	_ = validEmail
	if err != nil {

		err = validation.Username(identifier)
		if err != nil {
			return
		}
		identifierIsEmail = false
	}

	var user db.LoginUserByEmailRow

	if identifierIsEmail {
		user, err = service.store.LoginUserByEmail(ctx, identifier)

		if err != nil {
			systemError.Log(systemError.InternalMessages["LoginByEmail"](err))
			return createdUser, fmt.Errorf("server error on login by email")
		}

	} else {
		user1, err := service.store.LoginUserByUsername(ctx, identifier)

		if err != nil {
			systemError.Log(systemError.InternalMessages["LoginByUsername"](err))
			return createdUser, fmt.Errorf("server error on login by username")
		}

		user = db.LoginUserByEmailRow(user1)
	}

	fmt.Println("identifierIsEmail", user)
	return createdUser, nil
}
