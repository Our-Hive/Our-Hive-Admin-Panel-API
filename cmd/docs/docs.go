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
        "/contact-lines": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all contact lines with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact Line"
                ],
                "summary": "Get all contact lines",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start after",
                        "name": "startAfter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ContactLine"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a contact line",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact Line"
                ],
                "summary": "Create a contact line",
                "parameters": [
                    {
                        "description": "Contact Line",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateContactLine"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/contact-lines/{id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a contact line",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact Line"
                ],
                "summary": "Update a contact line",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Contact Line ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Contact Line",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateContactLine"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/generation": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Generate an IA image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "generation"
                ],
                "summary": "Generate IA Image",
                "parameters": [
                    {
                        "description": "Generate IA Image",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GenerateIAImage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/response.UploadImage"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "409": {
                        "description": "Conflict"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/images": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all images with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get all images",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start after",
                        "name": "startAfter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Image"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload an image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Upload an image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/response.UploadImage"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/images/approval": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all images by approved status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Get all images by approved status",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Approved",
                        "name": "approved",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start after",
                        "name": "startAfter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Image"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/images/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Approve an image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Approve an image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/recommended-content": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all recommended content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recommended Content"
                ],
                "summary": "Get all recommended content",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.DigitalContent"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new recommended content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recommended Content"
                ],
                "summary": "Create a new recommended content",
                "parameters": [
                    {
                        "description": "Recommended Content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateDigitalContent"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ContactLine": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.DigitalContent": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "request.CreateContactLine": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 1
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "request.CreateDigitalContent": {
            "type": "object",
            "required": [
                "description",
                "title",
                "url"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 1
                },
                "title": {
                    "type": "string",
                    "minLength": 1
                },
                "url": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "request.GenerateIAImage": {
            "type": "object",
            "required": [
                "file_name",
                "prompt"
            ],
            "properties": {
                "file_name": {
                    "type": "string",
                    "minLength": 1
                },
                "prompt": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "response.Image": {
            "type": "object",
            "properties": {
                "contentType": {
                    "type": "string"
                },
                "createdTime": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isApproved": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "response.UploadImage": {
            "type": "object",
            "properties": {
                "url": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
