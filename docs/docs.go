// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-20 20:49:18.510905 +0900 KST m=+0.029160042

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account service provider",
                        "name": "provider",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account password (email provider only)",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "account social_id",
                        "name": "social_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "description": "Create Account data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dto.Account"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "update account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Account data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dto.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "delete account by id",
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            }
        },
        "/files": {
            "post": {
                "description": "create file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "accountId",
                        "name": "account_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "file usage",
                        "name": "usage",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Profile image file",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {}
                }
            }
        },
        "/files/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "file Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.File"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Account": {
            "type": "object",
            "required": [
                "email",
                "provider"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "image_key": {
                    "type": "string",
                    "example": "image_key"
                },
                "interest": {
                    "type": "string",
                    "example": "develop"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "provider": {
                    "type": "string",
                    "example": "gmail"
                },
                "social_id": {
                    "type": "string",
                    "example": "social_id"
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "accesstoken"
                },
                "created_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "image_url": {
                    "type": "string",
                    "example": "profile.image_url.com"
                },
                "interest": {
                    "type": "string",
                    "example": "develop"
                },
                "provider": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                }
            }
        },
        "model.File": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "created_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "id": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "image_url": {
                    "type": "string",
                    "example": "profile.image_url.com"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "usage": {
                    "type": "string",
                    "example": "profile"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
