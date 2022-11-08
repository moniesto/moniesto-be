package clientError

type errorMessagesType map[string]string

const (
	Account_Login_InvalidBody         = "Account_Login_InvalidBody"
	Account_Login_InvalidEmail        = "Account_Login_InvalidEmail"
	Account_Login_InvalidUsername     = "Account_Login_InvalidUsername"
	Account_Login_NotFoundEmail       = "Account_Login_NotFoundEmail"
	Account_Login_NotFoundUsername    = "Account_Login_NotFoundUsername"
	Account_Login_ServerErrorEmail    = "Account_Login_ServerErrorEmail"
	Account_Login_ServerErrorUsername = "Account_Login_ServerErrorUsername"
	Account_Login_WrongPassword       = "Account_Login_WrongPassword"
	Account_Login_ServerErrorToken    = "Account_Login_ServerErrorToken"

	Account_Register_InvalidBody              = "Account_Register_InvalidBody"
	Account_Register_InvalidEmail             = "Account_Register_InvalidEmail"
	Account_Register_InvalidUsername          = "Account_Register_InvalidUsername"
	Account_Register_InvalidPassword          = "Account_Register_InvalidPassword"
	Account_Register_ServerErrorCheckEmail    = "Account_Register_ServerErrorCheckEmail"
	Account_Register_ServerErrorCheckUsername = "Account_Register_ServerErrorCheckUsername"
	Account_Register_RegisteredEmail          = "Account_Register_RegisteredEmail"
	Account_Register_RegisteredUsername       = "Account_Register_RegisteredUsername"
	Account_Register_ServerErrorPassword      = "Account_Register_ServerErrorPassword"
	Account_Register_ServerErrorCreateUser    = "Account_Register_ServerErrorCreateUser"
)

var errorMessages errorMessagesType = errorMessagesType{
	Account_Login_InvalidBody:         "Login request body is invalid",
	Account_Login_InvalidEmail:        "Email is invalid",
	Account_Login_InvalidUsername:     "Username is invalid",
	Account_Login_NotFoundEmail:       "Email is not in the system",
	Account_Login_NotFoundUsername:    "Username is not in the system",
	Account_Login_ServerErrorEmail:    "Server error on login with email",
	Account_Login_ServerErrorUsername: "Server error on login with username",
	Account_Login_WrongPassword:       "Wrong password",
	Account_Login_ServerErrorToken:    "Server error token operation",

	Account_Register_InvalidBody:              "Register request body is invalid",
	Account_Register_InvalidEmail:             "Email is invalid",
	Account_Register_InvalidUsername:          "Username is invalid",
	Account_Register_InvalidPassword:          "Password is invalid",
	Account_Register_ServerErrorCheckEmail:    "Server error on checking email",
	Account_Register_ServerErrorCheckUsername: "Server error on checking username",
	Account_Register_RegisteredEmail:          "Email is already registered",
	Account_Register_RegisteredUsername:       "Username is already registered",
	Account_Register_ServerErrorPassword:      "Server error on password operation",
	Account_Register_ServerErrorCreateUser:    "Server error on creating user",

	"Account_Login_InvalidPassword": "Password is invalid",
}
