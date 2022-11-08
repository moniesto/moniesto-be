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

type ChangeLoggedUserPasswordCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	body       any
	check      func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder)
}

func TestChangeLoggedUserPassword(t *testing.T) {
	changePasswordUsers = getRandomUsersDataWithID(5)

	changeLoggedUserPasswordCases := getChangeLoggedUserPasswordCases()

	for _, testCase := range changeLoggedUserPasswordCases {
		server := newTestServer(t)

		recorder := httptest.NewRecorder()
		ctx_test, _ := gin.CreateTestContext(recorder)
		testCase.initialize(t, ctx_test, server.service)

		url := "/account/password"

		request, err := http.NewRequest(http.MethodPut, url, createBody(testCase.body))
		require.NoError(t, err)

		testCase.setupAuth(t, request, server.tokenMaker)

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

func getChangeLoggedUserPasswordCases() ChangeLoggedUserPasswordCases {
	return ChangeLoggedUserPasswordCases{
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
