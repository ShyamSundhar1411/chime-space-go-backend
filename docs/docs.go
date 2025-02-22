// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/auth/login/": {
            "post": {
                "description": "Logs a user in by validating credentials and returning access and refresh tokens.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Request Payload",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, returns access and refresh tokens",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid credentials or user not found",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup/": {
            "post": {
                "description": "Signs up a new user by validating the fields and returns their access token and refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Signup",
                "parameters": [
                    {
                        "description": "SignUp Request Payload",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Signup successful, returns access and refresh tokens",
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid or missing parameters in the request",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Issue processing the request",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    }
                }
            }
        },
        "/chimes/": {
            "get": {
                "description": "Fetch all Chimes from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chimes"
                ],
                "summary": "Get all Chimes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.ChimeListResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeListResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new chime by providing a request body with necessary details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chimes"
                ],
                "summary": "Creates a new Chime",
                "parameters": [
                    {
                        "description": "Chime Create Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeCreateOrUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Chime"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeResponse"
                        }
                    }
                }
            }
        },
        "/chimes/user/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetch all Chimes from logged in user from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chimes"
                ],
                "summary": "Get all Chimes of logged in user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Chime"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeListResponse"
                        }
                    }
                }
            }
        },
        "/chimes/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing chime by providing the chime ID and updated details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chimes"
                ],
                "summary": "Update an existing Chime",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chime ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Chime Update Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeCreateOrUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Chime"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ChimeResponse"
                        }
                    }
                }
            }
        },
        "/user/me/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrives the profile details of the user based on access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Me Endpoint",
                "responses": {
                    "201": {
                        "description": "Returns Profile of the user",
                        "schema": {
                            "$ref": "#/definitions/domain.ProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Profile Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Chime": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "chimeContent": {
                    "type": "string"
                },
                "chimeTitle": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                }
            }
        },
        "domain.ChimeCreateOrUpdateRequest": {
            "type": "object",
            "required": [
                "chimeContent",
                "chimeTitle",
                "isPrivate"
            ],
            "properties": {
                "chimeContent": {
                    "type": "string"
                },
                "chimeTitle": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                }
            }
        },
        "domain.ChimeListResponse": {
            "type": "object",
            "properties": {
                "chimes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ChimeWithAuthor"
                    }
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "domain.ChimeResponse": {
            "type": "object",
            "properties": {
                "chime": {
                    "$ref": "#/definitions/domain.ChimeWithAuthor"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "domain.ChimeWithAuthor": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/domain.User"
                },
                "chimeContent": {
                    "type": "string"
                },
                "chimeTitle": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                }
            }
        },
        "domain.BaseResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "domain.ProfileResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/domain.User"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "userName"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "penName": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "penName": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "API documentation for ChimeSpace backend",
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
