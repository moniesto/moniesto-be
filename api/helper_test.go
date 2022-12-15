package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/stretchr/testify/require"
)

func getRandomUsersDataWithID(count int) []registeredUserWithID {
	all_users := []registeredUserWithID{}

	for i := 0; i < count; i++ {
		random_user := getRandomUserData()
		all_users = append(all_users, registeredUserWithID{
			ID: "",
			RegisterRequest: model.RegisterRequest{
				Name:     random_user.Name,
				Surname:  random_user.Surname,
				Username: random_user.Username,
				Email:    random_user.Email,
				Password: random_user.Password,
			},
		})
	}
	return all_users
}

func addAuthorzation(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	generalPaylod token.GeneralPaylod,
	duration time.Duration) {

	token, err := tokenMaker.CreateToken(generalPaylod, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func getRandomUsersData(count int) []model.RegisterRequest {
	all_users := []model.RegisterRequest{}

	for i := 0; i < count; i++ {
		all_users = append(all_users, getRandomUserData())
	}

	return all_users
}

func createUser(t *testing.T, ctx *gin.Context, service *service.Service, registerRequest model.RegisterRequest) db.User {
	user, err := service.CreateUser(ctx, registerRequest)
	require.NoError(t, err)

	return user
}

// createUserDBLevel creates user on DB level
func createUserDBLevel(t *testing.T, ctx *gin.Context, service *service.Service, registerRequest model.RegisterRequest) {
	hashed_password, err := util.HashPassword(registerRequest.Password)
	require.NoError(t, err)

	dbUser := db.CreateUserParams{
		ID:       core.CreateID(),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: hashed_password,
	}

	_, err = service.Store.CreateUser(ctx, dbUser)
	require.NoError(t, err)
}

// getRandomUserData returns 1 random model.RegisterRequest
func getRandomUserData() model.RegisterRequest {
	return model.RegisterRequest{
		Name:     util.RandomName(),
		Surname:  util.RandomSurname(),
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
	}
}

// createBody converts body into byte format
func createBody(body any) *bytes.Reader {
	bodyBytes := new(bytes.Buffer)
	json.NewEncoder(bodyBytes).Encode(body)

	return bytes.NewReader(bodyBytes.Bytes())
}

func createValidatingTokens(count int) []string {
	passwordResetTokens := []string{}

	for i := 0; i < count; i++ {
		passwordResetTokens = append(passwordResetTokens, token.CreateValidatingToken())
	}

	return passwordResetTokens
}
