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
        "/cities": {
            "get": {
                "tags": [
                    "города"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.City"
                            }
                        }
                    }
                }
            }
        },
        "/images": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "изображения"
                ],
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "multi",
                        "description": "изображения",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/images/{path}": {
            "get": {
                "tags": [
                    "изображения"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "пути к изображению",
                        "name": "path",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/meDetails": {
            "get": {
                "tags": [
                    "пользователи"
                ],
                "responses": {}
            }
        },
        "/sign-in": {
            "post": {
                "tags": [
                    "пользователи"
                ],
                "parameters": [
                    {
                        "description": "реквизиты для входа",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Credentials"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/sign-up": {
            "post": {
                "tags": [
                    "пользователи"
                ],
                "parameters": [
                    {
                        "description": "пользователь",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/{id}": {
            "get": {
                "tags": [
                    "пользователи"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "идентификатор",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entities.City": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Credentials": {
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
        "entities.UserDTO": {
            "type": "object",
            "properties": {
                "cityID": {
                    "type": "string"
                },
                "imagePath": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
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
