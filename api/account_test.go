package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {

}

var users []model.RegisterRequest

type LoginCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	body       any
	check      func(t *testing.T, recoder *httptest.ResponseRecorder)
}

func TestLogin(t *testing.T) {
	users = getRandomUsersData(6)

	loginTestCases := getLoginCases()

	for _, testCase := range loginTestCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {
			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			url := "/account/login"

			request, err := http.NewRequest(http.MethodPost, url, createBody(testCase.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, recorder)
		})
	}
}

func createBody(body any) *bytes.Reader {
	bodyBytes := new(bytes.Buffer)
	json.NewEncoder(bodyBytes).Encode(body)

	return bytes.NewReader(bodyBytes.Bytes())
}

func getRandomUserData() model.RegisterRequest {
	return model.RegisterRequest{
		Name:     util.RandomName(),
		Surname:  util.RandomSurname(),
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
	}
}

func getRandomUsersData(count int) []model.RegisterRequest {
	all_users := []model.RegisterRequest{}

	for i := 0; i < count; i++ {
		all_users = append(all_users, getRandomUserData())
	}

	return all_users
}

func createUser(t *testing.T, ctx *gin.Context, service *service.Service, registerRequest model.RegisterRequest) {
	_, err := service.CreateUser(ctx, registerRequest)
	require.NoError(t, err)
}

func checkSuccessLoginResponse(t *testing.T, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotResponse model.LoginResponse
	err = json.Unmarshal(data, &gotResponse)
	require.NoError(t, err)

	require.NotEmpty(t, gotResponse.Token)
	require.NotEmpty(t, gotResponse.User)
	require.NotEmpty(t, gotResponse.User.Id)
	require.NotEmpty(t, gotResponse.User.Name)
	require.NotEmpty(t, gotResponse.User.Surname)
	require.NotEmpty(t, gotResponse.User.Username)
	require.NotEmpty(t, gotResponse.User.Email)

	require.NotNil(t, gotResponse.User.EmailVerified)
	require.IsType(t, false, gotResponse.User.EmailVerified)

	require.NotEmpty(t, gotResponse.User.CreatedAt)
	require.IsType(t, time.Time{}, gotResponse.User.CreatedAt)

	require.NotEmpty(t, gotResponse.User.UpdatedAt)
	require.IsType(t, time.Time{}, gotResponse.User.UpdatedAt)
}

// TEST CASES
// LOGIN test cases
func getLoginCases() LoginCases {
	return LoginCases{
		{
			name:       "Invalid Body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: struct {
				Invalid string
			}{
				Invalid: "invalid",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Empty Body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body:       model.LoginRequest{},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "test@t.c",
				Password:   "testtest",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:       "Invalid Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "test test",
				Password:   "testtest",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:       "Invalid Identifier",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "",
				Password:   "tesstest",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "testtest",
				Password:   "tes",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Unauthorized with Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "test@gmail.com",
				Password:   "testtest",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:       "Unauthorized with Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.LoginRequest{
				Identifier: "testusername",
				Password:   "testtest",
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Success with Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[0])
			},
			body: model.LoginRequest{
				Identifier: users[0].Email,
				Password:   users[0].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				checkSuccessLoginResponse(t, recorder.Body)
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Success with Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[1])
			},
			body: model.LoginRequest{
				Identifier: users[1].Username,
				Password:   users[1].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				checkSuccessLoginResponse(t, recorder.Body)
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Wrong username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[2])
			},
			body: model.LoginRequest{
				Identifier: users[2].Username + util.RandomString(2),
				Password:   users[2].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[3])
			},
			body: model.LoginRequest{
				Identifier: util.RandomString(2) + users[3].Username,
				Password:   users[3].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong password with Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[4])
			},
			body: model.LoginRequest{
				Identifier: users[4].Username,
				Password:   users[4].Password + util.RandomString(2),
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong password with Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, users[5])
			},
			body: model.LoginRequest{
				Identifier: users[5].Email,
				Password:   users[5].Password + util.RandomString(2),
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}
}
