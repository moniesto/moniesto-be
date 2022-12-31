// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/login": {
            "post": {
                "description": "Login with [email \u0026 password] OR [username \u0026 password]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "identifier can be email OR username",
                        "name": "LoginBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    },
                    "403": {
                        "description": "wrong password",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "email OR username not found",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid body \u0026 data",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/password": {
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Authenticated user password change",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "description": " ",
                        "name": "ChangePasswordBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "wrong password",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid body \u0026 data",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/password/send_email": {
            "post": {
                "description": "Unauthenticated user send reset password email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password"
                ],
                "summary": "Send Reset Password Email",
                "parameters": [
                    {
                        "description": " ",
                        "name": "SendEmailBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendResetPasswordEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "406": {
                        "description": "invalid body \u0026 data",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/password/verify_token": {
            "post": {
                "description": "Unauthenticated verify token \u0026 change password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password"
                ],
                "summary": "Verify Token \u0026 Change Password",
                "parameters": [
                    {
                        "description": "token \u0026 new fiels are required",
                        "name": "VerifyTokenBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VerifyPasswordResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "token is expired",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "reset token not found",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid body \u0026 data",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/register": {
            "post": {
                "description": "Register as user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body\"",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterResponse"
                        }
                    },
                    "403": {
                        "description": "wrong password",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid body \u0026 data",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/crypto/currencies": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Search crypto currencies by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crypto"
                ],
                "summary": "Crypto Currency Search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Currency"
                            }
                        }
                    },
                    "406": {
                        "description": "invalid name",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error \u0026 crypto api error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/moniests": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Turn into moniest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Moniest"
                ],
                "summary": "Be Moniest",
                "parameters": [
                    {
                        "description": " ",
                        "name": "CreateMoniest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMoniestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.OwnUser"
                        }
                    },
                    "400": {
                        "description": "user is already moniest",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "forbidden operation: email is not verified",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "not found user",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid body",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/usernames/:username/check": {
            "get": {
                "description": "Check username is valid of not",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Check Username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CheckUsernameResponse"
                        }
                    },
                    "406": {
                        "description": "invalid username",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/:username": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "get user info with username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by Username",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'email' field will be visible if user request for own account",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "404": {
                        "description": "not any user with this username",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "invalid username",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "$ref": "#/definitions/clientError.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "clientError.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "error_code": {
                    "type": "string"
                }
            }
        },
        "model.ChangePasswordRequest": {
            "type": "object",
            "required": [
                "new",
                "old"
            ],
            "properties": {
                "new": {
                    "type": "string"
                },
                "old": {
                    "type": "string"
                }
            }
        },
        "model.CheckUsernameResponse": {
            "type": "object",
            "properties": {
                "validity": {
                    "type": "boolean"
                }
            }
        },
        "model.CreateMoniestRequest": {
            "type": "object",
            "required": [
                "card_id",
                "fee"
            ],
            "properties": {
                "bio": {
                    "description": "optional",
                    "type": "string"
                },
                "card_id": {
                    "type": "string"
                },
                "description": {
                    "description": "optional",
                    "type": "string"
                },
                "fee": {
                    "type": "number"
                },
                "message": {
                    "description": "optional",
                    "type": "string"
                }
            }
        },
        "model.Currency": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "identifier",
                "password"
            ],
            "properties": {
                "identifier": {
                    "type": "string",
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.OwnUser"
                }
            }
        },
        "model.Moniest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                },
                "subscription_info": {
                    "$ref": "#/definitions/model.SubscriptionInfo"
                }
            }
        },
        "model.OwnUser": {
            "type": "object",
            "properties": {
                "background_photo_link": {
                    "type": "string"
                },
                "background_photo_thumbnail_link": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "email_verified": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "moniest": {
                    "$ref": "#/definitions/model.Moniest"
                },
                "name": {
                    "type": "string"
                },
                "profile_photo_link": {
                    "type": "string"
                },
                "profile_photo_thumbnail_link": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "surname",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "password": {
                    "type": "string"
                },
                "surname": {
                    "type": "string",
                    "minLength": 1
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.RegisterResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.OwnUser"
                }
            }
        },
        "model.SendResetPasswordEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "model.SubscriptionInfo": {
            "type": "object",
            "properties": {
                "fee": {
                    "type": "number"
                },
                "message": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "background_photo_link": {
                    "type": "string"
                },
                "background_photo_thumbnail_link": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email_verified": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "moniest": {
                    "$ref": "#/definitions/model.Moniest"
                },
                "name": {
                    "type": "string"
                },
                "profile_photo_link": {
                    "type": "string"
                },
                "profile_photo_thumbnail_link": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.VerifyPasswordResetRequest": {
            "type": "object",
            "required": [
                "new",
                "token"
            ],
            "properties": {
                "new": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
