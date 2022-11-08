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

	"Account_Login_InvalidPassword": "Password is invalid",
}
