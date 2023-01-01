package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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

type registeredUserWithID struct {
	ID string
	model.RegisterRequest
}

var changePasswordUsers []registeredUserWithID
var sendPasswordResetEmailUsers []registeredUserWithID
var verifyTokenChangePasswordUsers []registeredUserWithID
var passwordResetTokens []string

const (
	SendPasswordResetEmailEndpoint    = "/account/password/send_email"
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

	changePasswordCases := getChangePasswordCases()

	for _, testCase := range changePasswordCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {
			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			request, err := http.NewRequest(http.MethodPut, ChangePasswordEndpoint, createBody(testCase.body))
			require.NoError(t, err)

			testCase.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, ctx_test, server.service, recorder)
		})

	}
}

func TestSendPasswordResetEmail(t *testing.T) {
	sendPasswordResetEmailUsers = getRandomUsersDataWithID(1)

	sendPasswordResetEmailCases := getSendPasswordResetEmailCases()

	for _, testCase := range sendPasswordResetEmailCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {
			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			request, err := http.NewRequest(http.MethodPost, SendPasswordResetEmailEndpoint, createBody(testCase.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, ctx_test, server.service, recorder)
		})

	}
}

func TestVerifyTokenChangePassword(t *testing.T) {
	verifyTokenChangePasswordUsers = getRandomUsersDataWithID(3)
	passwordResetTokens = createValidatingTokens(3)

	verifyTokenChangePasswordCases := getVerifyTokenChangePasswordCases()

	for _, testCase := range verifyTokenChangePasswordCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {

			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_test, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_test, server.service)

			request, err := http.NewRequest(http.MethodPost, VerifyTokenChangePasswordEndpoint, createBody(testCase.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, ctx_test, server.service, recorder)
		})

	}
}

func checkUserPasswordIs(t *testing.T, ctx *gin.Context, service *service.Service, user_id, password string) {
	hashedPassword, err := service.Store.GetPasswordByID(ctx, user_id)
	require.NoError(t, err)

	err = util.CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func getChangePasswordCases() ChangePasswordCases {
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

func getSendPasswordResetEmailCases() ChangePasswordCases {
	return ChangePasswordCases{
		{
			name:       "Invalid Body [Email does not provided]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Not exist email on system",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.SendResetPasswordEmailRequest{
				Email: util.RandomEmail(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
		{
			name: "Password Reset Email sent successfully",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				users := createUser(t, ctx, service, sendPasswordResetEmailUsers[0].RegisterRequest)
				sendPasswordResetEmailUsers[0].ID = users.ID

				// working directory from ./api -> ./ [to find the template]
				os.Chdir("../")
			},
			body: model.SendResetPasswordEmailRequest{
				Email: sendPasswordResetEmailUsers[0].Email,
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
	}
}

func getVerifyTokenChangePasswordCases() ChangePasswordCases {
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
		{
			name:       "Invalid Token [can't be decoded]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.VerifyPasswordResetRequest{
				Token:       util.RandomString(30),
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name:       "Valid Token [but not found/deleted]",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.VerifyPasswordResetRequest{
				Token:       token.EncodeValidatingToken(token.CreateValidatingToken()), // create new one
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "Expired Token",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				user := createUser(t, ctx, service, verifyTokenChangePasswordUsers[0].RegisterRequest)
				verifyTokenChangePasswordUsers[0].ID = user.ID

				params := db.CreatePasswordResetTokenParams{
					ID:          core.CreateID(),
					UserID:      verifyTokenChangePasswordUsers[0].ID,
					Token:       passwordResetTokens[0],
					TokenExpiry: time.Now().Add(-25 * time.Hour),
				}
				_, err := service.Store.CreatePasswordResetToken(ctx, params)
				require.NoError(t, err)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.VerifyPasswordResetRequest{
				Token:       core.Encode(passwordResetTokens[0]),
				NewPassword: util.RandomPassword(),
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "Invalid Password",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				user := createUser(t, ctx, service, verifyTokenChangePasswordUsers[1].RegisterRequest)
				verifyTokenChangePasswordUsers[1].ID = user.ID

				params := db.CreatePasswordResetTokenParams{
					ID:          core.CreateID(),
					UserID:      verifyTokenChangePasswordUsers[1].ID,
					Token:       passwordResetTokens[1],
					TokenExpiry: time.Now().Add(25 * time.Hour),
				}
				_, err := service.Store.CreatePasswordResetToken(ctx, params)
				require.NoError(t, err)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.VerifyPasswordResetRequest{
				Token:       core.Encode(passwordResetTokens[1]),
				NewPassword: "",
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
		{
			name: "Successful Password Reset Token Validate & Password change",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {
				user := createUser(t, ctx, service, verifyTokenChangePasswordUsers[2].RegisterRequest)
				verifyTokenChangePasswordUsers[2].ID = user.ID

				params := db.CreatePasswordResetTokenParams{
					ID:          core.CreateID(),
					UserID:      verifyTokenChangePasswordUsers[2].ID,
					Token:       passwordResetTokens[2],
					TokenExpiry: time.Now().Add(25 * time.Hour),
				}
				_, err := service.Store.CreatePasswordResetToken(ctx, params)
				require.NoError(t, err)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			body: model.VerifyPasswordResetRequest{
				Token:       core.Encode(passwordResetTokens[2]),
				NewPassword: "newpassword",
			},
			check: func(t *testing.T, ctx *gin.Context, service *service.Service, recorder *httptest.ResponseRecorder) {
				// TODO: check error code
				require.Equal(t, http.StatusOK, recorder.Code)
				checkUserPasswordIs(t, ctx, service, verifyTokenChangePasswordUsers[2].ID, "newpassword")
			},
		},
	}
}
