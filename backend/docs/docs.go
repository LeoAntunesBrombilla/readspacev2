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
        "/bookList": {
            "get": {
                "description": "Retrieve a list of all bookLists in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookList"
                ],
                "summary": "List all bookList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.UserEntity"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new bookList with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookList"
                ],
                "summary": "Create a new bookList",
                "parameters": [
                    {
                        "description": "BookList input for creation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.BookListInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the book list identified by the given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookList"
                ],
                "summary": "Delete a book list by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book List ID to delete",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Book list deleted with success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            }
        },
        "/books": {
            "post": {
                "description": "Create a new book with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book input for creation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ExternalBook"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete a book",
                "parameters": [
                    {
                        "description": "Book input for deletion",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DeleteBookInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate using username and password to get a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Authenticate a user and obtain a token",
                "parameters": [
                    {
                        "description": "Credentials for login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated, token returned",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Invalidate the user's current session, effectively logging them out",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Log out a user and invalidate the session",
                "responses": {
                    "200": {
                        "description": "Successfully logged out",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Retrieve a list of all users in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.UserEntity"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User details for creation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the user identified by the given ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID to delete",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted with success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update the user identified by the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID to update",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Details to update",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserUpdateDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated with success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "patch": {
                "description": "Update the password of the user identified by the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user password",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID to update password",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "New password details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserUpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User password updated with success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorEntity"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.BookListInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.DeleteBookInput": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "book_list_id": {
                    "type": "integer"
                }
            }
        },
        "entity.ErrorEntity": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ExternalBook": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "bookListId": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "googleBookId": {
                    "type": "string"
                },
                "imageLinks": {
                    "type": "object",
                    "properties": {
                        "smallThumbnail": {
                            "type": "string"
                        },
                        "thumbnail": {
                            "type": "string"
                        }
                    }
                },
                "language": {
                    "type": "string"
                },
                "pageCount": {
                    "type": "integer"
                },
                "publisher": {
                    "type": "string"
                },
                "subtitle": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entity.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.UserEntity": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.UserInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.UserUpdateDetails": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.UserUpdatePassword": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Readspace API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
