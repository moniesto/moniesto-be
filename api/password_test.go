package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
	"github.com/stretchr/testify/require"
)

type registeredUserWithID struct {
	ID string
	model.RegisterRequest
}

var changePasswordUsers []registeredUserWithID

const (
	SendEmailEndpoint                 = "/account/password/send_email"
	VerifyTokenChangePasswordEndpoint = "/account/password/verify_token"
	ChangePasswordEndpoint            = "/account/password"
)

type ChangePasswordCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	body       any
	check      func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder)
}

func TestChangePassword(t *testing.T) {
	changePasswordUsers = getRandomUsersDataWithID(5)

	changeLoggedInUserPasswordCases := getChangeLoggedInUserPasswordCases()

	for _, testCase := range changeLoggedInUserPasswordCases {
		server := newTestServer(t)

		recorder := httptest.NewRecorder()
		ctx_test, _ := gin.CreateTestContext(recorder)
		testCase.initialize(t, ctx_test, server.service)

		request, err := http.NewRequest(http.MethodPut, ChangePasswordEndpoint, createBody(testCase.body))
		require.NoError(t, err)

		testCase.setupAuth(t, request, server.tokenMaker)

		server.router.ServeHTTP(recorder, request)
		testCase.check(t, ctx_test, server.service, recorder)
	}
}

func _TestChangeLoggedOutUserPassword(t *testing.T) {

	changeLoggedOutUserPasswordCases := getChangeLoggedOutUserPasswordCases()

	for _, testCase := range changeLoggedOutUserPasswordCases {
		server := newTestServer(t)

		recorder := httptest.NewRecorder()
		ctx_test, _ := gin.CreateTestContext(recorder)
		testCase.initialize(t, ctx_test, server.service)

		url := "/account/password"

		request, err := http.NewRequest(http.MethodPut, url, createBody(testCase.body))
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		testCase.check(t, ctx_test, server.service, recorder)
	}

}

func checkUserPasswordIs(t *testing.T, ctx *gin.Context, service *service.Service, user_id, password string) {
	hashedPassword, err := service.Store.GetPasswordByID(ctx, user_id)
	require.NoError(t, err)

	err = util.CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func getChangeLoggedInUserPasswordCases() ChangePasswordCases {
	return ChangePasswordCases{
		{
			name: "Invalid body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, changePasswordUsers[0].RegisterRequest)
				changePasswordUsers[0].ID = users.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       changePasswordUsers[0].ID,
						Username: changePasswordUsers[0].Username,
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name: "Invalid Old Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, changePasswordUsers[1].RegisterRequest)
				changePasswordUsers[1].ID = users.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       changePasswordUsers[1].ID,
						Username: changePasswordUsers[1].Username,
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			body: model.ChangePasswordRequest{
				OldPassword: "",
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name: "Invalid New Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, changePasswordUsers[2].RegisterRequest)
				changePasswordUsers[2].ID = users.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       changePasswordUsers[2].ID,
						Username: changePasswordUsers[2].Username,
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			body: model.ChangePasswordRequest{
				OldPassword: util.RandomPassword(),
				NewPassword: "",
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name: "Wrong Old Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, changePasswordUsers[3].RegisterRequest)
				changePasswordUsers[3].ID = users.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       changePasswordUsers[3].ID,
						Username: changePasswordUsers[3].Username,
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			body: model.ChangePasswordRequest{
				OldPassword: util.RandomPassword(),
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "Successfully Change Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, changePasswordUsers[4].RegisterRequest)
				changePasswordUsers[4].ID = users.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       changePasswordUsers[4].ID,
						Username: changePasswordUsers[4].Username,
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			body: model.ChangePasswordRequest{
				OldPassword: changePasswordUsers[4].Password,
				NewPassword: "testtest",
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusOK, recorder.Code)
				checkUserPasswordIs(t, ctx, service, changePasswordUsers[4].ID, "testtest")
			},
		},
	}

}

func getChangeLoggedOutUserPasswordCases() ChangePasswordCases {
	/*
			cases:

			- invalid body:
		done	- token and email in the same time
		done	- email and new password in the same time
		done	- token alone
		done	- new password alone
			- send email case:
		done		- email is not exist on the system (still returns 202)
					- success with email (moniesto.test@gmail.com email)
			- verify token case:
				- invalid token (can not be decoded)
				- no record with this token
				- expired token with email (moniesto.test@gmail.com email)
				- success with email (moniesto.test@gmail.com email)

	*/

	return ChangePasswordCases{
		{
			name:       "Invalid body",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid body [Token and Email fields provided in the same time]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: struct {
				Email string `json:"email"`
				Token string `json:"token"`
			}{
				Email: util.RandomEmail(),
				Token: util.RandomString(30),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid body [Email and New Password fields provided in the same time]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: struct {
				Email       string `json:"email"`
				NewPassword string `json:"new"`
			}{
				Email:       util.RandomEmail(),
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid body [Only Token field provided (not new password)]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: struct {
				Token string `json:"token"`
			}{
				Token: util.RandomString(30),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Invalid body [Only New Password field provided]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: struct {
				NewPassword string `json:"new"`
			}{
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},

		// SEND EMAIL cases
		{
			name:       "[SEND RESET PASSWORD EMAIL] Not exist email on system",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: struct {
				Email string `json:"email"`
			}{
				Email: util.RandomEmail(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
		{
			name:       "[SUCCESS RESET PASSWORD EMAIL] Email sent successfully",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			// TODO: complete success case of send email
		},
	}
}
