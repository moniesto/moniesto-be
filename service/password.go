package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) CheckPassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: get old password [in hashed form]
	hashedPasword, err := service.Store.GetPasswordByID(ctx, user_id)
	if err != nil {
		systemError.Log(systemError.InternalMessages["GetPassword"](err))
		return errors.New("server error check password")
	}

	// STEP: check the password
	err = util.CheckPassword(password, hashedPasword)
	if err != nil {
		systemError.Log(systemError.InternalMessages["LoginFail"](err))
		return fmt.Errorf("wrong password")
	}

	return
}

func (service *Service) UpdatePassword(ctx *gin.Context, user_id, password string) (err error) {
	// STEP: hash password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return
	}

	setPasswordParams := db.SetPasswordParams{
		ID:       user_id,
		Password: hashedPassword,
	}

	// STEP: update/set password
	err = service.Store.SetPassword(ctx, setPasswordParams)
	if err != nil {
		systemError.Log(systemError.InternalMessages["UpdatePassword"](err))
		return fmt.Errorf("server error on update password")
	}

	return
}
