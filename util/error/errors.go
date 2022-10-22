package error

import "fmt"

type ErrorMessage struct {
	message string
	code    int
}

type InternalErrorMessage struct {
	message string
}

type errorMessageType map[string]func() (int, string)
type internalErrorMessageType map[string]func(error) string

func report(errorMessage ErrorMessage) func() (int, string) {
	return func() (int, string) {
		return errorMessage.code, errorMessage.message
	}
}

func internalReport(errorMessage InternalErrorMessage) func(err error) string {
	return func(err error) string {
		return fmt.Sprintln("-INTERNAL ERROR-", errorMessage.message, ":", err.Error())
	}
}

var Messages = errorMessageType{
	"Authorization": report(ErrorMessage{
		"Authorization error happened",
		401,
	}),
}

var InternalMessages = internalErrorMessageType{
	"TokenParse": internalReport(InternalErrorMessage{
		message: "Token parse problem",
	}),
}
