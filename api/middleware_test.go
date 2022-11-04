package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/moniesto/moniesto-be/token"
	"github.com/moniesto/moniesto-be/util"
)

type TestCases []struct {
	name          string
	setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	checkOptional func(t *testing.T, ctx *gin.Context)
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

func TestAuthMiddleware(t *testing.T) {

	testCases := TestCases{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {

				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, "unsupported", generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, "", generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)

			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestAuthMiddlewareOptional(t *testing.T) {
	testCases := TestCases{
		{
			name: "OK [with auth]",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {

				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			checkOptional: func(t *testing.T, ctx *gin.Context) {
				validAuth := ctx.MustGet(authorizationPayloadValidityKey).(bool)
				require.Equal(t, validAuth, true)
			},
		},
		{
			name: "OK [without auth]",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			checkOptional: func(t *testing.T, ctx *gin.Context) {
				validAuth := ctx.MustGet(authorizationPayloadValidityKey).(bool)
				require.Equal(t, validAuth, false)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, "unsupported", generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, "", generalPayload, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				generalPayload := token.GeneralPaylod{
					UserPayload: token.UserPayload{
						ID:       util.CreateID(),
						Username: "default_username",
					},
				}

				addAuthorzation(t, request, tokenMaker, authorizationTypeBearer, generalPayload, -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			authPath := "/authOptional"
			server.router.GET(
				authPath,
				authMiddlewareOptional(server.tokenMaker),
				func(ctx *gin.Context) {

					tc.checkOptional(t, ctx)

					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
