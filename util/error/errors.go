package error

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	message string
	code    int
}

type InternalErrorMessage struct {
	message string
}

type errorMessageType map[string]func(...string) (int, gin.H)
type internalErrorMessageType map[string]func(error) string

func report(errorMessage ErrorMessage) func(...string) (int, gin.H) {
	return func(moreMessage ...string) (int, gin.H) {

		message := errorMessage.message

		if len(moreMessage) > 0 {
			message += ": " + strings.Join(moreMessage[:], ", ")
		}

		return errorMessage.code, errorResponse(message)
	}
}

func errorResponse(err string) gin.H {
	return gin.H{"error": err}
}

func internalReport(errorMessage InternalErrorMessage) func(error) string {
	return func(err error) string {
		return fmt.Sprintln("INTERNAL ERROR:", errorMessage.message, ":", err.Error())
	}
}

var Messages = errorMessageType{
	"NotProvided_AuthorizationHeader": report(ErrorMessage{
		"Authorization header is not provided",
		http.StatusUnauthorized,
	}),
	"Invalid_AuthorizationHeader": report(ErrorMessage{
		"Invalid authorization header format",
		http.StatusUnauthorized,
	}),
	"Invalid_Token": report(ErrorMessage{
		"Token is invalid",
		http.StatusUnauthorized,
	}),
	"Unsupported_AuthorizationType": report(ErrorMessage{
		"Unsupported authorization type",
		http.StatusUnauthorized,
	}),
	"Invalid_RequestBody_Register": report(ErrorMessage{
		"Register request body is invalid",
		http.StatusNotAcceptable,
	}),
}

var InternalMessages = internalErrorMessageType{
	"RunService": internalReport(InternalErrorMessage{
		message: "Running service failed",
	}),
}
