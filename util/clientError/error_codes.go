package clientError

type ErrorMessagesType map[string]string

const (
	General_Maintenance = "General_Maintenance"

	General_UserNotMoniest                        = "General_UserNotMoniest"
	General_UserNotFoundByID                      = "General_UserNotFoundByID"
	General_UserNotFoundByUsername                = "General_UserNotFoundByUsername"
	General_MoniestNotFoundByUsername             = "General_MoniestNotFoundByUsername"
	General_ServerErrorGetMoniestByUsername       = "General_ServerErrorGetMoniestByUsername"
	General_ServerErrorCheckMoniestByUsername     = "General_ServerErrorCheckMoniestByUsername"
	General_ServerErrorCheckMoniestByUserID       = "General_ServerErrorCheckMoniestByUserID"
	General_ServerErrorGettingUserLanguageByEmail = "General_ServerErrorGettingUserLanguageByEmail"
	General_ServerErrorUserLanguageNotFound       = "General_ServerErrorUserLanguageNotFound"
	General_CalculatePNLandROI                    = "General_CalculatePNLandROI"
	General_Not_Admin                             = "General_Not_Admin"

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
	Account_Register_InvalidFullname          = "Account_Register_InvalidFullname"
	Account_Register_UnsupportedLanguage      = "Account_Register_UnsupportedLanguage"
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
	Account_UpdateUserProfile_InvalidFullname                  = "Account_UpdateUserProfile_InvalidFullname"
	Account_UpdateUserProfile_InvalidLocation                  = "Account_UpdateUserProfile_InvalidLocation"
	Account_UpdateUserProfile_UnsupportedLanguage              = "Account_UpdateUserProfile_UnsupportedLanguage"
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

	Moniest_UpdateMoniest_InvalidBody                       = "Moniest_UpdateMoniest_InvalidBody"
	Moniest_UpdateMoniest_ServerErrorGetUser                = "Moniest_UpdateMoniest_ServerErrorGetUser"
	Moniest_UpdateMoniest_InvalidBio                        = "Moniest_UpdateMoniest_InvalidBio"
	Moniest_UpdateMoniest_InvalidDescription                = "Moniest_UpdateMoniest_InvalidDescription"
	Moniest_UpdateMoniest_ServerErrorUpdateMoniest          = "Moniest_UpdateMoniest_ServerErrorUpdateMoniest"
	Moniest_UpdateMoniest_InvalidFee                        = "Moniest_UpdateMoniest_InvalidFee"
	Moniest_UpdateMoniest_InvalidSubscriptionMessage        = "Moniest_UpdateMoniest_InvalidSubscriptionMessage"
	Moniest_UpdateMoniest_ServerErrorGetSubscriptionInfo    = "Moniest_UpdateMoniest_ServerErrorGetSubscriptionInfo"
	Moniest_UpdateMoniest_ServerErrorUpdateSubscriptionInfo = "Moniest_UpdateMoniest_ServerErrorUpdateSubscriptionInfo"

	Moniest_GetPayoutInfo_ServerErrorGetMoniest    = "Moniest_GetPayoutInfo_ServerErrorGetMoniest"
	Moniest_GetPayoutInfo_ServerErrorGetPayoutInfo = "Moniest_GetPayoutInfo_ServerErrorGetPayoutInfo"
	Moniest_GetPayoutInfo_PayoutInfoNotFound       = "Moniest_GetPayoutInfo_PayoutInfoNotFound"

	Moniest_UpdatePayout_InvalidBody                 = "Moniest_UpdatePayout_InvalidBody"
	Moniest_UpdatePayout_ServerErrorGetMoniest       = "Moniest_UpdatePayout_ServerErrorGetMoniest"
	Moniest_UpdatePayout_ServerErrorUpdatePayoutInfo = "Moniest_UpdatePayout_ServerErrorUpdatePayoutInfo"

	Moniest_GetMoniest_NoMoniest             = "Moniest_GetMoniest_NoMoniest"
	Moniest_GetMoniest_ServerErrorGetMoniest = "Moniest_GetMoniest_ServerErrorGetMoniest"

	Moniest_Subscribe_InvalidBody                         = "Moniest_Subscribe_InvalidBody"
	Moniest_Subscribe_SubscribeOwn                        = "Moniest_Subscribe_SubscribeOwn"
	Moniest_Subscribe_ServerErrorGetSubscription          = "Moniest_Subscribe_ServerErrorGetSubscription"
	Moniest_Subscribe_AlreadySubscribed                   = "Moniest_Subscribe_AlreadySubscribed"
	Moniest_Subscribe_ServerErrorCreateBinanceOrder       = "Moniest_Subscribe_ServerErrorCreateBinanceOrder"
	Moniest_Subscribe_ServerErrorGetProductName           = "Moniest_Subscribe_ServerErrorGetProductName"
	Moniest_Subscribe_ServerErrorCreateBinanceTransaction = "Moniest_Subscribe_ServerErrorCreateBinanceTransaction"
	Moniest_Subscribe_ServerErrorActivateSubscription     = "Moniest_Subscribe_ServerErrorActivateSubscription"
	Moniest_Subscribe_ServerErrorCreateSubscriptionDB     = "Moniest_Subscribe_ServerErrorCreateSubscriptionDB"

	Moniest_Unsubscribe_UnsubscribeOwn              = "Moniest_Unsubscribe_UnsubscribeOwn"
	Moniest_Unsubscribe_NotSubscribed               = "Moniest_Unsubscribe_NotSubscribed"
	Moniest_Unsubscribe_ServerErrorUnsubscribe      = "Moniest_Unsubscribe_ServerErrorUnsubscribe"
	Moniest_Unsubscribe_ServerErrorGetPayoutHistory = "Moniest_Unsubscribe_ServerErrorGetPayoutHistory"
	Moniest_Unsubscribe_ServerErrorRefund           = "Moniest_Unsubscribe_ServerErrorRefund"

	Moniest_CreateSubscriptionInfo_InvalidFee                 = "Moniest_CreateSubscriptionInfo_InvalidFee"
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage = "Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage"
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate        = "Moniest_CreateSubscriptionInfo_ServerErrorOnCreate"

	Moniest_CreatePayoutInfo_InvalidBinanceID    = "Moniest_CreatePayoutInfo_InvalidBinanceID"
	Moniest_CreatePayoutInfo_ServerErrorOnCreate = "Moniest_CreatePayoutInfo_ServerErrorOnCreate"

	Moniest_CreatePostStatistics_ServerErrorOnCreate = "Moniest_CreatePostStatistics_ServerErrorOnCreate"

	Moniest_SubscribeCheck_ServerErrorCheck = "Moniest_SubscribeCheck_ServerErrorCheck"

	Moniest_GetSubscriber_InvalidParam              = "Moniest_GetSubscriber_InvalidParam"
	Moniest_GetSubscriber_ServerErrorGetSubscribers = "Moniest_GetSubscriber_ServerErrorGetSubscribers"

	Moniest_GetMoniestPosts_InvalidParam        = "Moniest_GetMoniestPosts_InvalidParam"
	Moniest_GetMoniestPosts_ForbiddenAccess     = "Moniest_GetMoniestPosts_ForbiddenAccess"
	Moniest_GetMoniestPosts_ServerErrorGetPosts = "Moniest_GetMoniestPosts_ServerErrorGetPosts"

	User_GetUser_ServerErrorGetUser = "User_GetUser_ServerErrorGetUser"

	Post_CreatePost_InvalidBody                  = "Post_CreatePost_InvalidBody"
	Post_CreatePost_InvalidCurrency              = "Post_CreatePost_InvalidCurrency"
	Post_CreatePost_InvalidDuration              = "Post_CreatePost_InvalidDuration"
	Post_CreatePost_InvalidCurrencyPrice         = "Post_CreatePost_InvalidCurrencyPrice"
	Post_CreatePost_InvalidTargets               = "Post_CreatePost_InvalidTargets"
	Post_CreatePost_InvalidTakeProfit            = "Post_CreatePost_InvalidTakeProfit"
	Post_CreatePost_InvalidStop                  = "Post_CreatePost_InvalidStop"
	Post_CreatePost_InvalidMarketType            = "Post_CreatePost_InvalidMarketType"
	Post_CreatePost_InvalidLeverage              = "Post_CreatePost_InvalidLeverage"
	Post_CreatePost_ServerErrorCreatePost        = "Post_CreatePost_ServerErrorCreatePost"
	Post_CreatePost_ServerErrorCreateDescription = "Post_CreatePost_ServerErrorCreateDescription"
	Post_CreatePost_ServerErrorPostPhotoUpload   = "Post_CreatePost_ServerErrorPostPhotoUpload"

	Crypto_GetCurrencies_InvalidParam       = "Crypto_GetCurrencies_InvalidParam"
	Crypto_GetCurrenciesFromAPI_ServerError = "Crypto_GetCurrenciesFromAPI_ServerError"
	Crypto_GetCurrencyFromAPI_ServerError   = "Crypto_GetCurrencyFromAPI_ServerError"

	Feedback_CreateFeedback_InvalidBody               = "Feedback_CreateFeedback_InvalidBody"
	Feedback_CreateFeedback_ServerErrorCreateFeedback = "Feedback_CreateFeedback_ServerErrorCreateFeedback"

	Content_GetPosts_InvalidParam        = "Content_GetPosts_InvalidParam"
	Content_GetPosts_ServerErrorGetPosts = "Content_GetPosts_ServerErrorGetPosts"

	Content_GetMoniests_InvalidParam           = "Content_GetMoniests_InvalidParam"
	Content_GetMoniests_ServerErrorGetMoniests = "Content_GetMoniests_ServerErrorGetMoniests"

	Content_SearchMoniests_InvalidParam             = "Content_SearchMoniests_InvalidParam"
	Content_SearchMoniests_ServerErrorSearchMoniest = "Content_SearchMoniests_ServerErrorSearchMoniest"

	User_GetSubscriptions_InvalidParam                = "User_GetSubscriptions_InvalidParam"
	User_GetSubscriptions_ServerErrorGetSubscriptions = "User_GetSubscriptions_ServerErrorGetSubscriptions"

	User_GetStats_ServerErrorGetStats    = "User_GetStats_ServerErrorGetStats"
	Moniest_GetStats_ServerErrorGetStats = "Moniest_GetStats_ServerErrorGetStats"

	Payment_CheckBinanceTransaction_TransactionIDNotFound          = "Payment_CheckBinanceTransaction_TransactionIDNotFound"
	Payment_CheckBinanceTransaction_ServerErrorGetTransaction      = "Payment_CheckBinanceTransaction_ServerErrorGetTransaction"
	Payment_CheckBinanceTransaction_ServerErrorQueryTransaction    = "Payment_CheckBinanceTransaction_ServerErrorQueryTransaction"
	Payment_CheckBinanceTransaction_ServerErrorUpdateStatusSuccess = "Payment_CheckBinanceTransaction_ServerErrorUpdateStatusSuccess"
	Payment_CheckBinanceTransaction_ServerErrorUpdateStatusFail    = "Payment_CheckBinanceTransaction_ServerErrorUpdateStatusFail"

	Moniest_GetUserSubscriptionInfo_ServerErrorGetSubscriptionInfo = "Moniest_GetUserSubscriptionInfo_ServerErrorGetSubscriptionInfo"

	Admin_GetMetrics_ServerErrorUserMetrics     = "Admin_GetMetrics_ServerErrorUserMetrics"
	Admin_GetMetrics_ServerErrorPostMetrics     = "Admin_GetMetrics_ServerErrorPostMetrics"
	Admin_GetMetrics_ServerErrorPaymentMetrics  = "Admin_GetMetrics_ServerErrorPaymentMetrics"
	Admin_GetMetrics_ServerErrorPayoutMetrics   = "Admin_GetMetrics_ServerErrorPayoutMetrics"
	Admin_GetMetrics_ServerErrorFeedbackMetrics = "Admin_GetMetrics_ServerErrorFeedbackMetrics"
	Admin_GetMetrics_ServerErrorGetFeedbacks    = "Admin_GetMetrics_ServerErrorGetFeedbacks"
	Admin_GetData_InvalidParam                  = "Admin_GetData_InvalidParam"
	Admin_GetData_ServerErrorGetData            = "Admin_GetData_ServerErrorGetData"
)

var errorMessages ErrorMessagesType = ErrorMessagesType{
	General_Maintenance: "Server is in maintenance mode",

	General_UserNotMoniest:                        "User is not moniest",
	General_UserNotFoundByID:                      "User not found with this user ID",
	General_UserNotFoundByUsername:                "User not found with this username",
	General_MoniestNotFoundByUsername:             "No moniest with this username",
	General_ServerErrorGetMoniestByUsername:       "Server error on getting moniest by username",
	General_ServerErrorCheckMoniestByUsername:     "Server error on checking user is moniest by username",
	General_ServerErrorCheckMoniestByUserID:       "Server error on checking user is moniest by user ID",
	General_ServerErrorGettingUserLanguageByEmail: "Server error on getting user language by email",
	General_ServerErrorUserLanguageNotFound:       "User language not found",
	General_CalculatePNLandROI:                    "Error while calculating pnl and roi",
	General_Not_Admin:                             "Not admin",

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
	Account_Register_InvalidFullname:          "Fullname is invalid",
	Account_Register_UnsupportedLanguage:      "Language is not supported. Supported languages: [en, tr]",
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
	Account_UpdateUserProfile_InvalidFullname:                  "Fullname is invalid",
	Account_UpdateUserProfile_InvalidLocation:                  "Location is invalid",
	Account_UpdateUserProfile_UnsupportedLanguage:              "Language is not supported. Supported languages: [en, tr]",
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

	Moniest_UpdateMoniest_InvalidBody:                       "Update moniest profile request body is invalid",
	Moniest_UpdateMoniest_ServerErrorGetUser:                "Server error on getting user",
	Moniest_UpdateMoniest_InvalidBio:                        "Bio is invalid",
	Moniest_UpdateMoniest_InvalidDescription:                "Description is invalid",
	Moniest_UpdateMoniest_ServerErrorUpdateMoniest:          "Server error on updating moniest",
	Moniest_UpdateMoniest_InvalidFee:                        "Fee is invalid",
	Moniest_UpdateMoniest_InvalidSubscriptionMessage:        "Subscription message is invalid",
	Moniest_UpdateMoniest_ServerErrorGetSubscriptionInfo:    "Server error on getting subscription info",
	Moniest_UpdateMoniest_ServerErrorUpdateSubscriptionInfo: "Server error on updating subscription info",

	Moniest_GetPayoutInfo_ServerErrorGetMoniest:    "Server error on get moniest",
	Moniest_GetPayoutInfo_ServerErrorGetPayoutInfo: "Server error on getting payout infos",
	Moniest_GetPayoutInfo_PayoutInfoNotFound:       "Moniest does not have any payout info",

	Moniest_UpdatePayout_InvalidBody:                 "Update moniest payout info request body is invalid",
	Moniest_UpdatePayout_ServerErrorGetMoniest:       "Server error on getting moniest",
	Moniest_UpdatePayout_ServerErrorUpdatePayoutInfo: "Server error on updating payout info",

	Moniest_GetMoniest_NoMoniest:             "Not any moniest exist",
	Moniest_GetMoniest_ServerErrorGetMoniest: "Server error on getting moniest",

	Moniest_Subscribe_InvalidBody:                     "Subscribe to moniest request body is invalid",
	Moniest_Subscribe_SubscribeOwn:                    "User can't subscribe to own",
	Moniest_Subscribe_ServerErrorGetSubscription:      "Server error on getting subscription",
	Moniest_Subscribe_AlreadySubscribed:               "Already subscribed to moniest",
	Moniest_Subscribe_ServerErrorCreateBinanceOrder:   "Server error on creating binance payment link",
	Moniest_Subscribe_ServerErrorGetProductName:       "Server error on getting product name",
	Moniest_Subscribe_ServerErrorActivateSubscription: "Server error on activating subscription",
	Moniest_Subscribe_ServerErrorCreateSubscriptionDB: "Server error on creating subscription on DB",

	Moniest_Unsubscribe_UnsubscribeOwn:              "User can't unsubscribe from own",
	Moniest_Unsubscribe_NotSubscribed:               "User is not subscribed to moniest",
	Moniest_Unsubscribe_ServerErrorUnsubscribe:      "Server error on unsubscribe",
	Moniest_Unsubscribe_ServerErrorGetPayoutHistory: "Server error on getting payout histories",
	Moniest_Unsubscribe_ServerErrorRefund:           "Server error on refund user",

	Moniest_CreateSubscriptionInfo_InvalidFee:                 "Fee is invalid",
	Moniest_CreateSubscriptionInfo_InvalidSubscriptionMessage: "Subscription message is invalid",
	Moniest_CreateSubscriptionInfo_ServerErrorOnCreate:        "Server error on create subscription info",

	Moniest_CreatePayoutInfo_InvalidBinanceID:    "Binance ID is not valid",
	Moniest_CreatePayoutInfo_ServerErrorOnCreate: "Server error on create payout info",

	Moniest_CreatePostStatistics_ServerErrorOnCreate: "Server error on create post statistics",

	Moniest_SubscribeCheck_ServerErrorCheck: "Server error on checking moniest subscribe status",

	Moniest_GetSubscriber_InvalidParam:              "Get subscribers invalid params",
	Moniest_GetSubscriber_ServerErrorGetSubscribers: "Server error on getting subscribed users",

	Moniest_GetMoniestPosts_InvalidParam:        "Get moniest posts invalid params",
	Moniest_GetMoniestPosts_ForbiddenAccess:     "User can not access active posts of moniests that are not subscribed",
	Moniest_GetMoniestPosts_ServerErrorGetPosts: "Server error on getting posts of moniest",

	User_GetUser_ServerErrorGetUser: "Server error on getting user",

	Post_CreatePost_InvalidBody:                  "Create post request body is invalid",
	Post_CreatePost_InvalidCurrency:              "Currency is invalid",
	Post_CreatePost_InvalidDuration:              "Duration is invalid",
	Post_CreatePost_InvalidCurrencyPrice:         "Currency price is invalid",
	Post_CreatePost_InvalidTargets:               "Targets are invalid",
	Post_CreatePost_InvalidTakeProfit:            "Take Profit is invalid",
	Post_CreatePost_InvalidStop:                  "Stop is invalid",
	Post_CreatePost_InvalidMarketType:            "Market type is invalid. Supported market types: [spot, futures]",
	Post_CreatePost_InvalidLeverage:              "Leverage is not valid",
	Post_CreatePost_ServerErrorCreatePost:        "Server error on creating post",
	Post_CreatePost_ServerErrorCreateDescription: "Server error on creating description",
	Post_CreatePost_ServerErrorPostPhotoUpload:   "Server error on uploading post photos (maybe invalid format)",

	Crypto_GetCurrencies_InvalidParam:       "Get Currencies request params is invalid",
	Crypto_GetCurrenciesFromAPI_ServerError: "Server error on getting currencies from API",
	Crypto_GetCurrencyFromAPI_ServerError:   "Server error on getting currency from API",

	Feedback_CreateFeedback_InvalidBody:               "Create feedback request body is invalid",
	Feedback_CreateFeedback_ServerErrorCreateFeedback: "Server error on creating feedback",

	Content_GetPosts_InvalidParam:        "Get Posts request param is invalid",
	Content_GetPosts_ServerErrorGetPosts: "Server error on getting posts",

	Content_GetMoniests_InvalidParam:           "Get Moniests request param is invalid",
	Content_GetMoniests_ServerErrorGetMoniests: "Server error on getting moniests",

	Content_SearchMoniests_InvalidParam:             "Search moniest request param is invalid",
	Content_SearchMoniests_ServerErrorSearchMoniest: "Server error on searching moniest",

	User_GetSubscriptions_InvalidParam:                "Get user subscriptions request param is invalid",
	User_GetSubscriptions_ServerErrorGetSubscriptions: "Server error on get user subscriptions",

	User_GetStats_ServerErrorGetStats:    "Server error on getting user stats",
	Moniest_GetStats_ServerErrorGetStats: "Server error on getting moniest stats",

	Payment_CheckBinanceTransaction_TransactionIDNotFound:          "TransactionID not found",
	Payment_CheckBinanceTransaction_ServerErrorGetTransaction:      "Server error on getting transaction data",
	Payment_CheckBinanceTransaction_ServerErrorQueryTransaction:    "Server error on query transaction",
	Payment_CheckBinanceTransaction_ServerErrorUpdateStatusSuccess: "Server error on updating transaction status [success case]",
	Payment_CheckBinanceTransaction_ServerErrorUpdateStatusFail:    "Server error on updating transaction status [fail case]",

	Moniest_GetUserSubscriptionInfo_ServerErrorGetSubscriptionInfo: "Server error on getting user subscription info",

	Admin_GetMetrics_ServerErrorUserMetrics:     "Server error on getting user metrics",
	Admin_GetMetrics_ServerErrorPostMetrics:     "Server error on getting post metrics",
	Admin_GetMetrics_ServerErrorPaymentMetrics:  "Server error on getting payment metrics",
	Admin_GetMetrics_ServerErrorPayoutMetrics:   "Server error on getting payout metrics",
	Admin_GetMetrics_ServerErrorFeedbackMetrics: "Server error on getting feedback metrics",
	Admin_GetMetrics_ServerErrorGetFeedbacks:    "Server error on getting feedbacks",
	Admin_GetData_InvalidParam:                  "Get data request param is invalid",
	Admin_GetData_ServerErrorGetData:            "Server error on getting data",
}
