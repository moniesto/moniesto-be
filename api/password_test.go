package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/token"
)

type ChangeLoggedUserPasswordCases []struct {
	name       string
	initialize func(t *testing.T, ctx *gin.Context, service *service.Service)
	setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	body       any
	check      func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestChangeLoggedUserPassword(t *testing.T) {

}

func getChangeLoggedUserPasswordCases() ChangeLoggedUserPasswordCases {
	return ChangeLoggedUserPasswordCases{
		{
			name:       "Authentication Error",
			initialize: func(t *testing.T, ctx *gin.Context, service *service.Service) {},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
		},
	}

}
