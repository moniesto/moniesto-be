package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/util"
	"github.com/stretchr/testify/require"
)

var loginUsers []model.RegisterRequest
var registerUsers []model.RegisterRequest
var checkUsernameUsers []model.RegisterRequest

type LoginCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	body       any
	check      func(t *testing.T, recorder *httptest.ResponseRecorder)
}

type RegisterCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	body       any
	check      func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder)
}

type CheckUsernameCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	username   string
	check      func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder)
}

func TestLogin(t *testing.T) {
	loginUsers = getRandomUsersData(6)

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

func TestRegister(t *testing.T) {
	registerUsers = getRandomUsersData(6)

	registerTestCases := getRegisterCases()

	for _, testCase := range registerTestCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {
			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			url := "/account/register"

			request, err := http.NewRequest(http.MethodPost, url, createBody(testCase.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, ctx_test, server.service, recorder)
		})
	}
}

func TestCheckUsername(t *testing.T) {
	checkUsernameUsers = getRandomUsersData(2)

	checkUsernameTestCases := getCheckUsernameCases()

	for _, testCase := range checkUsernameTestCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {
			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			url := fmt.Sprintf("/account/usernames/%s/check", testCase.username)

			request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, ctx_test, server.service, recorder)
		})

	}

}

// HELPERS
// createBody converts body into byte format
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

func createUserDBLevel(t *testing.T, ctx *gin.Context, service *service.Service, registerRequest model.RegisterRequest) {
	dbUser := db.CreateUserParams{
		ID:       util.CreateID(),
		Name:     registerRequest.Name,
		Surname:  registerRequest.Surname,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	_, err := service.Store.CreateUser(ctx, dbUser)
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

func checkUserNotInSystemByEmail(t *testing.T, ctx *gin.Context, service *service.Service, email string) {
	_, err := service.Store.GetUserByEmail(ctx, email)
	require.Equal(t, err, sql.ErrNoRows)
}

func checkUserInSystemByEmail(t *testing.T, ctx *gin.Context, service *service.Service, email string) {
	_, err := service.Store.GetUserByEmail(ctx, email)
	require.NoError(t, err)
}

func checkUserInSystemByUsername(t *testing.T, ctx *gin.Context, service *service.Service, username string) {
	_, err := service.Store.GetUserByUsername(ctx, username)
	require.NoError(t, err)
}

func checkUserNotInSystemByUsername(t *testing.T, ctx *gin.Context, service *service.Service, username string) {
	_, err := service.Store.GetUserByUsername(ctx, username)
	require.Equal(t, err, sql.ErrNoRows)
}

func checkUsernameValidityBody(t *testing.T, body *bytes.Buffer, expected bool) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotResponse model.CheckUsernameResponse
	err = json.Unmarshal(data, &gotResponse)
	require.NoError(t, err)

	require.Equal(t, gotResponse.Validity, expected)
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
				Identifier: "test@.c", // is invalid
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
				Identifier: "test test", // is invalid
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
				Identifier: "", // is invalid
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
				Password:   "tes", // is invalid
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
			name: "Successful Login with Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[0])
			},
			body: model.LoginRequest{
				Identifier: loginUsers[0].Email,
				Password:   loginUsers[0].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				checkSuccessLoginResponse(t, recorder.Body)
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Successful Login with Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[1])
			},
			body: model.LoginRequest{
				Identifier: loginUsers[1].Username,
				Password:   loginUsers[1].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				checkSuccessLoginResponse(t, recorder.Body)
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Wrong username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[2])
			},
			body: model.LoginRequest{
				Identifier: loginUsers[2].Username + util.RandomString(2),
				Password:   loginUsers[2].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[3])
			},
			body: model.LoginRequest{
				Identifier: util.RandomString(2) + loginUsers[3].Username,
				Password:   loginUsers[3].Password,
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong password with Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[4])
			},
			body: model.LoginRequest{
				Identifier: loginUsers[4].Username,
				Password:   loginUsers[4].Password + util.RandomString(2),
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Wrong password with Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, loginUsers[5])
			},
			body: model.LoginRequest{
				Identifier: loginUsers[5].Email,
				Password:   loginUsers[5].Password + util.RandomString(2),
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}
}

func getRegisterCases() RegisterCases {
	return RegisterCases{
		{
			name:       "Invalid Body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: struct {
				Invalid string
			}{
				Invalid: "invalid",
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Empty Body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body:       model.RegisterRequest{},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid Email",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.RegisterRequest{
				Name:     registerUsers[0].Name,
				Surname:  registerUsers[0].Surname,
				Username: registerUsers[0].Username,
				Email:    "test@.c", // is invalid
				Password: registerUsers[0].Password,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)

				checkUserNotInSystemByEmail(t, ctx, service, "test@.c")
				checkUserNotInSystemByUsername(t, ctx, service, registerUsers[0].Username)
			},
		},
		{
			name:       "Invalid Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.RegisterRequest{
				Name:     registerUsers[1].Name,
				Surname:  registerUsers[1].Surname,
				Username: registerUsers[1].Username,
				Email:    registerUsers[1].Email,
				Password: "tes", // is invalid
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)

				checkUserNotInSystemByEmail(t, ctx, service, registerUsers[1].Email)
				checkUserNotInSystemByUsername(t, ctx, service, registerUsers[1].Username)
			},
		},
		{
			name:       "Invalid Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			body: model.RegisterRequest{
				Name:     registerUsers[2].Name,
				Surname:  registerUsers[2].Surname,
				Username: "Test Username", // is invalid
				Email:    registerUsers[2].Email,
				Password: registerUsers[2].Password,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)

				checkUserNotInSystemByEmail(t, ctx, service, registerUsers[2].Email)
				checkUserNotInSystemByUsername(t, ctx, service, "Test Username")
			},
		},
		{
			name: "Email already in system",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				checkUserNotInSystemByEmail(t, ctx, service, registerUsers[3].Email)

				createUserDBLevel(t, ctx, service, model.RegisterRequest{
					Name:     registerUsers[3].Name,
					Surname:  registerUsers[3].Surname,
					Username: registerUsers[3].Username,
					Email:    registerUsers[3].Email,
					Password: registerUsers[3].Password,
				})
			},
			body: model.RegisterRequest{
				Name:     registerUsers[3].Name,
				Surname:  registerUsers[3].Surname,
				Username: registerUsers[3].Username,
				Email:    registerUsers[3].Email,
				Password: registerUsers[3].Password,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)

				checkUserInSystemByEmail(t, ctx, service, registerUsers[3].Email)
			},
		},
		{
			name: "Username already in system",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {

				checkUserNotInSystemByUsername(t, ctx, service, registerUsers[4].Username)

				createUserDBLevel(t, ctx, service, model.RegisterRequest{
					Name:     registerUsers[4].Name,
					Surname:  registerUsers[4].Surname,
					Username: registerUsers[4].Username,
					Email:    registerUsers[4].Email,
					Password: registerUsers[4].Password,
				})
			},
			body: model.RegisterRequest{
				Name:     registerUsers[4].Name,
				Surname:  registerUsers[4].Surname,
				Username: registerUsers[4].Username,
				Email:    registerUsers[4].Email,
				Password: registerUsers[4].Password,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)

				checkUserInSystemByUsername(t, ctx, service, registerUsers[4].Username)
			},
		},
		{
			name: "Successful Register",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				checkUserNotInSystemByEmail(t, ctx, service, registerUsers[5].Email)
				checkUserNotInSystemByUsername(t, ctx, service, registerUsers[5].Username)
			},
			body: model.RegisterRequest{
				Name:     registerUsers[5].Name,
				Surname:  registerUsers[5].Surname,
				Username: registerUsers[5].Username,
				Email:    registerUsers[5].Email,
				Password: registerUsers[5].Password,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				checkSuccessLoginResponse(t, recorder.Body)
				require.Equal(t, http.StatusOK, recorder.Code)

				checkUserInSystemByEmail(t, ctx, service, registerUsers[5].Email)

				checkUserInSystemByUsername(t, ctx, service, registerUsers[5].Username)
			},
		},
	}
}

func getCheckUsernameCases() CheckUsernameCases {
	return CheckUsernameCases{
		{
			name:       "Invalid Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			username:   "",
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "Invalid Username",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			username:   "test username",
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check there is an error message
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Validity true",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
			},
			username: checkUsernameUsers[0].Username,
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				checkUserNotInSystemByUsername(t, ctx, service, checkUsernameUsers[0].Username)
				require.Equal(t, http.StatusOK, recorder.Code)
				checkUsernameValidityBody(t, recorder.Body, true)
			},
		},
		{
			name: "Validity false",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				createUser(t, ctx, service, checkUsernameUsers[1])
			},
			username: checkUsernameUsers[1].Username,
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				checkUserInSystemByUsername(t, ctx, service, checkUsernameUsers[1].Username)
				require.Equal(t, http.StatusOK, recorder.Code)
				checkUsernameValidityBody(t, recorder.Body, false)
			},
		},
	}
}
