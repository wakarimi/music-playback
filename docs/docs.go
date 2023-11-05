// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dmitry Kolesnikov (Zalimannard)",
            "email": "zalimannard@mail.ru"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rooms": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "Creates a room",
                "parameters": [
                    {
                        "type": "string",
                        "default": "en-US",
                        "description": "Language preference",
                        "name": "Produce-Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "X-Account-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Room data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room_handler.createRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/room_handler.createResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Invalid X-Account-ID header format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/rooms/{roomID}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "Deletes a room",
                "parameters": [
                    {
                        "type": "string",
                        "default": "en-US",
                        "description": "Language preference",
                        "name": "Produce-Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "X-Account-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Trying to delete someone else's room; Invalid X-Account-ID header format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "The room does not exist",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/rooms/{roomID}/share": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ShareCode"
                ],
                "summary": "Receives a code to connect to the room",
                "parameters": [
                    {
                        "type": "string",
                        "default": "en-US",
                        "description": "Language preference",
                        "name": "Produce-Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "X-Account-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/room_handler.generateShareCodeResponse"
                        }
                    },
                    "403": {
                        "description": "Trying to generate a code for someone else's room; Invalid X-Account-ID header format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "The room does not exist",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ShareCode"
                ],
                "summary": "Creates or recreates a code to connect to a room",
                "parameters": [
                    {
                        "type": "string",
                        "default": "en-US",
                        "description": "Language preference",
                        "name": "Produce-Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "X-Account-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/room_handler.generateShareCodeResponse"
                        }
                    },
                    "403": {
                        "description": "Trying to generate a code for someone else's room; Invalid X-Account-ID header format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "The room does not exist",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/rooms/{roomID}/share-reset": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ShareCode"
                ],
                "summary": "Reset a code to connect to a room",
                "parameters": [
                    {
                        "type": "string",
                        "default": "en-US",
                        "description": "Language preference",
                        "name": "Produce-Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "X-Account-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Room ID",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Trying to reset a code for someone else's room; Invalid X-Account-ID header format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "The room does not exist",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.PlaybackOrderType": {
            "type": "string",
            "enum": [
                "IN_ORDER",
                "REPLAY",
                "RANDOM"
            ],
            "x-enum-varnames": [
                "InOrder",
                "Replay",
                "Random"
            ]
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Human-readable error message",
                    "type": "string"
                },
                "reason": {
                    "description": "Internal error description",
                    "type": "string"
                }
            }
        },
        "room_handler.createRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "description": "Desired room name",
                    "type": "string"
                }
            }
        },
        "room_handler.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the created room",
                    "type": "integer"
                },
                "name": {
                    "description": "Name of the created room",
                    "type": "string"
                },
                "ownerID": {
                    "description": "Room owner",
                    "type": "integer"
                },
                "playbackOrderType": {
                    "description": "Playback order in the created room",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PlaybackOrderType"
                        }
                    ]
                }
            }
        },
        "room_handler.generateShareCodeResponse": {
            "type": "object",
            "properties": {
                "shareCode": {
                    "description": "Code to connect to the room",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.4.2",
	Host:             "localhost:8024",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Wakarimi Music Playback API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
