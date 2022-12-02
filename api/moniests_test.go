package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
	"github.com/stretchr/testify/require"
)

type CreateMoniestCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	body       any
	check      func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestCreateMoniest(t *testing.T) {

	createMoniestCases := getCreateMoniestTestCases()

	for _, testCase := range createMoniestCases {

		t.Run(fmt.Sprintf("CASE:%s", testCase.name), func(t *testing.T) {

			server := newTestServer(t)

			recorder := httptest.NewRecorder()
			ctx_text, _ := gin.CreateTestContext(recorder)
			testCase.initialize(t, ctx_text, server.service)

			url := "/moniests"

			request, err := http.NewRequest(http.MethodPost, url, createBody(testCase.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.check(t, recorder)

		})

	}

}

// getCreateMoniestTestCases returns Create Moniest test cases
func getCreateMoniestTestCases() CreateMoniestCases {

	return CreateMoniestCases{
		{
			name: "Unauthenticated",
		},
		{
			name: "Invalid body [no Fee]",
		},
		{
			name: "Invalid body [no CardID]",
		},
		{
			name: "Not found user",
		},
		{
			name: "User is already a moniest",
		},
		{
			name: "User email is not verified",
		},
		{
			name: "Invalid Bio",
		},
		{
			name: "Invalid Description",
		},
		{
			name: "Invalid Fee",
		},
		{
			name: "Invalid Subscription Message",
		},
		{
			name: "Already created Subscription info",
		},
		{
			name: "Successfully Moniest creation",
		},
		// TODO: add card payment info tests
	}
}
