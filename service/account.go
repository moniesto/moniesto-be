package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) CreateUser(ctx *gin.Context, registerRequest model.RegisterRequest) error {
	fmt.Println("Creating user", registerRequest)

	validEmail, err := validation.Email(registerRequest.Email)
	if err != nil {
		return err
	}
	_ = validEmail

	err = validation.Password(registerRequest.Password)
	if err != nil {
		return err
	}

	// TODO: complete register

	// any user with same email
	checkEmail, err := service.store.CheckEmail(ctx, validEmail)
	fmt.Println("checkemail", checkEmail)

	// any user with same username

	// hash password

	user := db.CreateUserParams{
		ID:       util.CreateID(),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	_ = user

	// store password
	// service.store.CreateUser(ctx, user)

	return nil
}
