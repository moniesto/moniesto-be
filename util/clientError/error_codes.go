package clientError

type ErrorMessagesType map[string]string

const (
	UserNotMoniest   = "UserNotMoniest"
	UserNotFoundByID = "UserNotFoundByID"

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

	Account_EmailVerification_InvalidBody            = "Account_EmailVerification_InvalidBody"
	Account_EmailVerification_ServerErrorGetUser     = "Account_EmailVerification_ServerErrorGetUser"
	Account_EmailVerification_AlreadyVerified        = "Account_EmailVerification_AlreadyVerified"
	Account_EmailVerification_ServerErrorCreateToken = "Account_EmailVerification_ServerErrorCreateToken"
	Account_EmailVerification_SendEmail              = "Account_EmailVerification_SendEmail"
	Account_EmailVerification_InvalidToken           = "Account_EmailVerification_InvalidToken"
	Account_EmailVerification_NotFoundToken          = "Account_EmailVerification_NotFoundToken"
	Account_EmailVerification_ServerErrorGetToken    = "Account_EmailVerification_ServerErrorGetToken"
	Account_EmailVerification_ExpiredToken           = "Account_EmailVerification_ExpiredToken"
	Account_EmailVerification_ServerErrorVerifyEmail = "Account_EmailVerification_ServerErrorVerifyEmail"
	Account_EmailVerification_ServerErrorDeleteToken = "Account_EmailVerification_ServerErrorDeleteToken"

	Account_ChangeUsername_InvalidBody               = "Account_ChangeUsername_InvalidBody"
	Account_ChangeUsername_RegisteredUsername        = "Account_ChangeUsername_RegisteredUsername"
	Account_ChangeUsername_InvalidUsername           = "Account_ChangeUsername_InvalidUsername"
	Account_ChangeUsername_ServerErrorChangeUsername = "Account_ChangeUsername_ServerErrorChangeUsername"
	Account_ChangeUsername_ServerErrorToken          = "Account_ChangeUsername_ServerErrorToken"

	Account_UpdateUserProfile_InvalidBody                      = "Account_UpdateUserProfile_InvalidBody"
	Account_UpdateUserProfile_ServerErrorGetUser               = "Account_UpdateUserProfile_ServerErrorGetUser"
	Account_UpdateUserProfile_ServerErrorUpdateUser            = "Account_UpdateUserProfile_ServerErrorUpdateUser"
	Account_UpdateUserProfile_ServerErrorGetProfilePhoto       = "Account_UpdateUserProfile_ServerErrorGetProfilePhoto"
	Account_UpdateUserProfile_ServerErrorGetBackgroundPhoto    = "Account_UpdateUserProfile_ServerErrorGetBackgroundPhoto"
	Account_UpdateUserProfile_ServerErrorUploadProfilePhoto    = "Account_UpdateUserProfile_ServerErrorUploadProfilePhoto"
	Account_UpdateUserProfile_ServerErrorUploadBackgroundPhoto = "Account_UpdateUserProfile_ServerErrorUploadBackgroundPhoto"
	Account_UpdateUserProfile_ServerErrorUpdateProfilePhoto    = "Account_UpdateUserProfile_ServerErrorUpdateProfilePhoto"
	Account_UpdateUserProfile_ServerErrorUpdateBackgroundPhoto = "Account_UpdateUserProfile_ServerErrorUpdateBackgroundPhoto"
	Account_UpdateUserProfile_ServerErrorInsertProfilePhoto    = "Account_UpdateUserProfile_ServerErrorInsertProfilePhoto"
	Account_UpdateUserProfile_ServerErrorInsertBackgroundPhoto = "Account_UpdateUserProfile_ServerErrorInsertBackgroundPhoto"

	Moniest_CreateMoniest_InvalidBody              = "Moniest_CreateMoniest_InvalidBody"
	Moniest_CreateMoniest_ServerErrorUserIsMoniest = "Moniest_CreateMoniest_ServerErrorUserIsMoniest"
	Moniest_CreateMoniest_UserIsAlreadyMoniest     = "Moniest_CreateMoniest_UserIsAlreadyMoniest"
	Moniest_CreateMoniest_InvalidBio               = "Moniest_CreateMoniest_InvalidBio"
	Moniest_CreateMoniest_InvalidDescription       = "Moniest_CreateMoniest_InvalidDescription"
	Moniest_CreateMoniest_ServerErrorCreateMoniest = "Moniest_CreateMoniest_ServerErrorCreateMoniest"
	Moniest_CreateMoniest_UnverifiedEmail          = "Moniest_CreateMoniest_UnverifiedEmail"

	Moniest_UpdateMoniest_InvalidBody = "Moniest_UpdateMoniest_InvalidBody"

	Moniest_GetMoniest_NoMoniest             = "Moniest_GetMoniest_NoMoniest"
	Moniest_GetMoniest_ServerErrorGetMoniest = "Moniest_GetMoniest_ServerErrorGetMoniest"

	Moniest_CreateSubscriptionInfo_InvalidFee                 = "Moniest_CreateSubscriptionInfo_InvalidFee"
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage = "Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage"
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate        = "Moniest_CreateSubscriptionInfo_ServerErrorOnCreate"

	User_GetUser_NotFoundUser       = "User_GetUser_NotFoundUser"
	User_GetUser_ServerErrorGetUser = "User_GetUser_ServerErrorGetUser"

	Post_CreatePost_InvalidBody                  = "Post_CreatePost_InvalidBody"
	Post_CreatePost_InvalidCurrency              = "Post_CreatePost_InvalidCurrency"
	Post_CreatePost_InvalidDuration              = "Post_CreatePost_InvalidDuration"
	Post_CreatePost_InvalidCurrencyPrice         = "Post_CreatePost_InvalidCurrencyPrice"
	Post_CreatePost_InvalidTargets               = "Post_CreatePost_InvalidTargets"
	Post_CreatePost_InvalidStop                  = "Post_CreatePost_InvalidStop"
	Post_CreatePost_ServerErrorCreatePost        = "Post_CreatePost_ServerErrorCreatePost"
	Post_CreatePost_ServerErrorCreateDescription = "Post_CreatePost_ServerErrorCreateDescription"
	Post_CreatePost_ServerErrorPostPhotoUpload   = "Post_CreatePost_ServerErrorPostPhotoUpload"

	Crypto_GetCurrencies_InvalidParam       = "Crypto_GetCurrencies_InvalidParam"
	Crypto_GetCurrenciesFromAPI_ServerError = "Crypto_GetCurrenciesFromAPI_ServerError"
)

var errorMessages ErrorMessagesType = ErrorMessagesType{
	UserNotMoniest:   "User is not moniest",
	UserNotFoundByID: "User not found with this user ID",

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

	Account_EmailVerification_InvalidBody:            "Email verification request body is invalid",
	Account_EmailVerification_ServerErrorGetUser:     "Server error on getting user",
	Account_EmailVerification_AlreadyVerified:        "The email is already verified",
	Account_EmailVerification_ServerErrorCreateToken: "Server error on creating token",
	Account_EmailVerification_SendEmail:              "Server error on sending email",
	Account_EmailVerification_InvalidToken:           "Email Verification Token is invalid",
	Account_EmailVerification_NotFoundToken:          "Email Verification Token is not in the system",
	Account_EmailVerification_ServerErrorGetToken:    "Server error on getting Email Verification Token from system",
	Account_EmailVerification_ExpiredToken:           "Email Verification Token is expired",
	Account_EmailVerification_ServerErrorVerifyEmail: "Server error on verifying email",
	Account_EmailVerification_ServerErrorDeleteToken: "Server error on deleting Email Verification Token",

	Account_ChangeUsername_InvalidBody:               "Change username request body is invalid",
	Account_ChangeUsername_RegisteredUsername:        "Username is already registered",
	Account_ChangeUsername_InvalidUsername:           "Username is invalid",
	Account_ChangeUsername_ServerErrorChangeUsername: "Server error on changing username",
	Account_ChangeUsername_ServerErrorToken:          "Server error on token operation",

	Account_UpdateUserProfile_InvalidBody:                      "Update user profile request body is invalid",
	Account_UpdateUserProfile_ServerErrorGetUser:               "Server error on getting user",
	Account_UpdateUserProfile_ServerErrorUpdateUser:            "Server error updating user",
	Account_UpdateUserProfile_ServerErrorGetProfilePhoto:       "Server error on getting profile photo",
	Account_UpdateUserProfile_ServerErrorGetBackgroundPhoto:    "Server error on getting background photo",
	Account_UpdateUserProfile_ServerErrorUploadProfilePhoto:    "Server error on uploading profile photo",
	Account_UpdateUserProfile_ServerErrorUploadBackgroundPhoto: "Server error on uploading background photo",
	Account_UpdateUserProfile_ServerErrorUpdateProfilePhoto:    "Server error on updating profile photo",
	Account_UpdateUserProfile_ServerErrorUpdateBackgroundPhoto: "Server error on updating background photo",
	Account_UpdateUserProfile_ServerErrorInsertProfilePhoto:    "Server error on inserting profile photo",
	Account_UpdateUserProfile_ServerErrorInsertBackgroundPhoto: "Server error on inserting background photo",

	Moniest_CreateMoniest_InvalidBody:              "Create moniest request body is invalid",
	Moniest_CreateMoniest_ServerErrorUserIsMoniest: "Server error on user is moniest check",
	Moniest_CreateMoniest_UserIsAlreadyMoniest:     "This user is already a moniest",
	Moniest_CreateMoniest_InvalidBio:               "Bio is invalid",
	Moniest_CreateMoniest_InvalidDescription:       "Description is invalid",
	Moniest_CreateMoniest_ServerErrorCreateMoniest: "Server error on create moniest",
	Moniest_CreateMoniest_UnverifiedEmail:          "Email is not verified yet",

	Moniest_UpdateMoniest_InvalidBody: "Update moniest profile request body is invalid",

	Moniest_GetMoniest_NoMoniest:             "Not any moniest exist",
	Moniest_GetMoniest_ServerErrorGetMoniest: "Server error on getting moniest",

	Moniest_CreateSubscriptionInfo_InvalidFee:                 "Fee is invalid",
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage: "Subscription message is invalid",
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate:        "Server error on create subscription info",

	User_GetUser_NotFoundUser:       "User not found",
	User_GetUser_ServerErrorGetUser: "Server error on getting user",

	Post_CreatePost_InvalidBody:                  "Create post request body is invalid",
	Post_CreatePost_InvalidCurrency:              "Currency is invalid",
	Post_CreatePost_InvalidDuration:              "Duration is invalid",
	Post_CreatePost_InvalidCurrencyPrice:         "Currency price is invalid",
	Post_CreatePost_InvalidTargets:               "Targets are invalid",
	Post_CreatePost_InvalidStop:                  "Stop is invalid",
	Post_CreatePost_ServerErrorCreatePost:        "Server error on creating post",
	Post_CreatePost_ServerErrorCreateDescription: "Server error on creating description",
	Post_CreatePost_ServerErrorPostPhotoUpload:   "Server error on uploading post photos (maybe invalid format)",

	Crypto_GetCurrencies_InvalidParam:       "Get Currencies request params is invalid",
	Crypto_GetCurrenciesFromAPI_ServerError: "Server error on getting currencies from API",
}
