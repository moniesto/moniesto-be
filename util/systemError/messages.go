package systemError

import (
	"fmt"
)

type InternalErrorMessage struct {
	message string
}

type internalErrorMessageType map[string]func(error) string

func internalReport(errorMessage InternalErrorMessage) func(error) string {
	return func(err error) string {
		return fmt.Sprintln("INTERNAL ERROR:", errorMessage.message, ":", err.Error())
	}
}

var InternalMessages = internalErrorMessageType{
	"RunService": internalReport(InternalErrorMessage{
		message: "Running service failed",
	}),
	"CheckEmail": internalReport(InternalErrorMessage{
		message: "Server error on check email",
	}),
	"CheckUsername": internalReport(InternalErrorMessage{
		message: "Server error on check username",
	}),
	"CreateUser": internalReport(InternalErrorMessage{
		message: "Server error on creating user",
	}),
	"LoginByEmail": internalReport(InternalErrorMessage{
		message: "Server error on login by email",
	}),
	"LoginByUsername": internalReport(InternalErrorMessage{
		message: "Server error on login by username",
	}),
	"LoginFail": internalReport(InternalErrorMessage{
		message: "Login failed",
	}),
	"UpdateLoginStatsFail": internalReport(InternalErrorMessage{
		message: "Server error on updating login stats",
	}),
	"GetPassword": internalReport(InternalErrorMessage{
		message: "Server error on getting password",
	}),
	"UpdatePassword": internalReport(InternalErrorMessage{
		message: "Server error on updating password",
	}),
	"GetUserByEmail": internalReport(InternalErrorMessage{
		message: "Server error on getting user by email",
	}),
	"CreatePasswordResetToken": internalReport(InternalErrorMessage{
		message: "Server error on creating password reset token instance on db",
	}),
	"SendPasswordResetEmail": internalReport(InternalErrorMessage{
		message: "Server error on sending password reset email",
	}),
	"SendEmailVerificationEmail": internalReport(InternalErrorMessage{
		message: "Server error on sending email verification email",
	}),
	"SendWelcomingEmail": internalReport(InternalErrorMessage{
		message: "Server error on sending welcoming email",
	}),
}
