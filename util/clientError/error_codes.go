package clientError

type errorMessagesType map[string]string

const (
	Account_Authorization_NotProvidedHeader = "Account_Authorization_NotProvidedHeader"
	Account_Authorization_InvalidHeader     = "Account_Authorization_InvalidHeader"
	Account_Authorization_UnsupportedType   = "Account_Authorization_UnsupportedType"
	Account_Authorization_InvalidToken      = "Account_Authorization_InvalidToken"

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

	Account_GetUser_NotFound    = "Account_GetUser_NotFound"
	Account_GetUser_ServerError = "Account_GetUser_ServerError"

	Account_CheckUsername_InvalidUsername          = "Account_CheckUsername_InvalidUsername"
	Account_CheckUsername_ServerErrorCheckUsername = "Account_CheckUsername_ServerErrorCheckUsername"

	Account_ChangePassword_InvalidBody               = "Account_ChangePassword_InvalidBody"
	Account_ChangePassword_InvalidOldPassword        = "Account_ChangePassword_InvalidOldPassword"
	Account_ChangePassword_InvalidNewPassword        = "Account_ChangePassword_InvalidNewPassword"
	Account_ChangePassword_InvalidEmail              = "Account_ChangePassword_InvalidEmail"
	Account_ChangePassword_NotFoundEmail             = "Account_ChangePassword_NotFoundEmail"
	Account_ChangePassword_ServerErrorCheckPassword  = "Account_ChangePassword_ServerErrorCheckPassword"
	Account_ChangePassword_ServerErrorCheckEmail     = "Account_ChangePassword_ServerErrorCheckEmail"
	Account_ChangePassword_WrongPassword             = "Account_ChangePassword_WrongPassword"
	Account_ChangePassword_ServerErrorPassword       = "Account_ChangePassword_ServerErrorPassword"
	Account_ChangePassword_ServerErrorUpdatePassword = "Account_ChangePassword_ServerErrorUpdatePassword"
	Account_ChangePassword_ServerErrorCreateToken    = "Account_ChangePassword_ServerErrorCreateToken"
	Account_ChangePassowrd_SendEmail                 = "Account_ChangePassowrd_SendEmail"
	Account_ChangePassword_InvalidToken              = "Account_ChangePassword_InvalidToken"
	Account_ChangePassword_NotFoundToken             = "Account_ChangePassword_NotFoundToken"
	Account_ChangePassword_ServerErrorGetToken       = "Account_ChangePassword_ServerErrorGetToken"
	Account_ChangePassword_ExpiredToken              = "Account_ChangePassword_ExpiredToken"
	Account_ChangePassword_ServerErrorDeleteToken    = "Account_ChangePassword_ServerErrorDeleteToken"

	Moniest_CreateMoniest_InvalidBody              = "Moniest_CreateMoniest_InvalidBody"
	Moniest_CreateMoniest_ServerErrorUserIsMoniest = "Moniest_CreateMoniest_ServerErrorUserIsMoniest"
	Moniest_CreateMoniest_UserIsAlreadyMoniest     = "Moniest_CreateMoniest_UserIsAlreadyMoniest"
	Moniest_CreateMoniest_InvalidBio               = "Moniest_CreateMoniest_InvalidBio"
	Moniest_CreateMoniest_InvalidDescription       = "Moniest_CreateMoniest_InvalidDescription"
	Moniest_CreateMoniest_ServerErrorCreateMoniest = "Moniest_CreateMoniest_ServerErrorCreateMoniest"
	Moniest_CreateMoniest_UnverifiedEmail          = "Moniest_CreateMoniest_UnverifiedEmail"

	Moniest_GetMoniest_NoMoniest             = "Moniest_GetMoniest_NoMoniest"
	Moniest_GetMoniest_ServerErrorGetMoniest = "Moniest_GetMoniest_ServerErrorGetMoniest"

	Moniest_CreateSubscriptionInfo_InvalidFee                 = "Moniest_CreateSubscriptionInfo_InvalidFee"
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage = "Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage"
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate        = "Moniest_CreateSubscriptionInfo_ServerErrorOnCreate"

	User_GetUser_NotFoundUser       = "User_GetUser_NotFoundUser"
	User_GetUser_ServerErrorGetUser = "User_GetUser_ServerErrorGetUser"

	Post_CreatePost_InvalidBody     = "Post_CreatePost_InvalidBody"
	Post_CreatePost_InvalidCurrency = "Post_CreatePost_InvalidCurrency"

	Crypto_GetCurrencies_InvalidParam       = "Crypto_GetCurrencies_InvalidParam"
	Crypto_GetCurrenciesFromAPI_ServerError = "Crypto_GetCurrenciesFromAPI_ServerError"
)

var errorMessages errorMessagesType = errorMessagesType{
	Account_Authorization_NotProvidedHeader: "Authorization Header is not provided",
	Account_Authorization_InvalidHeader:     "Authorization Header is invalid",
	Account_Authorization_UnsupportedType:   "Authorization Type is not supported",
	Account_Authorization_InvalidToken:      "Authorization Token is invalid",

	Account_Login_InvalidBody:         "Login request body is invalid",
	Account_Login_InvalidEmail:        "Email is invalid",
	Account_Login_InvalidUsername:     "Username is invalid",
	Account_Login_NotFoundEmail:       "Email is not in the system",
	Account_Login_NotFoundUsername:    "Username is not in the system",
	Account_Login_ServerErrorEmail:    "Server error on login with email",
	Account_Login_ServerErrorUsername: "Server error on login with username",
	Account_Login_WrongPassword:       "Wrong password",
	Account_Login_ServerErrorToken:    "Server error on token operation",

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

	Account_GetUser_NotFound:    "User not found",
	Account_GetUser_ServerError: "Server error on getting user",

	Account_CheckUsername_InvalidUsername:          "Username is invalid",
	Account_CheckUsername_ServerErrorCheckUsername: "Server error on checking username",

	Account_ChangePassword_InvalidBody:               "Change password request body is invalid",
	Account_ChangePassword_InvalidOldPassword:        "Old password is invalid",
	Account_ChangePassword_InvalidNewPassword:        "New password is invalid",
	Account_ChangePassword_InvalidEmail:              "Email is invalid",
	Account_ChangePassword_NotFoundEmail:             "Email is not in the system",
	Account_ChangePassword_ServerErrorCheckPassword:  "Server error on checking password",
	Account_ChangePassword_ServerErrorCheckEmail:     "Server error on checking email",
	Account_ChangePassword_WrongPassword:             "Wrong old Password",
	Account_ChangePassword_ServerErrorPassword:       "Server error on password operation",
	Account_ChangePassword_ServerErrorUpdatePassword: "Server error on updating password",
	Account_ChangePassword_ServerErrorCreateToken:    "Server error on creating token",
	Account_ChangePassowrd_SendEmail:                 "Server error on sending email",
	Account_ChangePassword_InvalidToken:              "Password Reset Token is invalid",
	Account_ChangePassword_NotFoundToken:             "Password Reset Token is not in the system",
	Account_ChangePassword_ServerErrorGetToken:       "Server error on getting Password Reset Token from system",
	Account_ChangePassword_ExpiredToken:              "Password Reset Token is expired",
	Account_ChangePassword_ServerErrorDeleteToken:    "Server error on deleting Password Reset Token",

	Moniest_CreateMoniest_InvalidBody:              "Create moniest request body is invalid",
	Moniest_CreateMoniest_ServerErrorUserIsMoniest: "Server error on user is moniest check",
	Moniest_CreateMoniest_UserIsAlreadyMoniest:     "This user is already a moniest",
	Moniest_CreateMoniest_InvalidBio:               "Bio is invalid",
	Moniest_CreateMoniest_InvalidDescription:       "Description is invalid",
	Moniest_CreateMoniest_ServerErrorCreateMoniest: "Server error on create moniest",
	Moniest_CreateMoniest_UnverifiedEmail:          "Email is not verified yet",

	Moniest_GetMoniest_NoMoniest:             "Not any moniest exist",
	Moniest_GetMoniest_ServerErrorGetMoniest: "Server error on getting moniest",

	Moniest_CreateSubscriptionInfo_InvalidFee:                 "Fee is invalid",
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage: "Subscription message is invalid",
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate:        "Server error on create subscription info",

	User_GetUser_NotFoundUser:       "User not found",
	User_GetUser_ServerErrorGetUser: "Server error on getting user",

	Post_CreatePost_InvalidBody:     "Create post request body is invalid",
	Post_CreatePost_InvalidCurrency: "Currency is invalid",

	Crypto_GetCurrencies_InvalidParam:       "Get Currencies request params is invalid",
	Crypto_GetCurrenciesFromAPI_ServerError: "Server error on getting currencies from API",
}
