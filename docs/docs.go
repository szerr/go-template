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
        "/auth/sig_in": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "SigIn",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SigIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sig_out": {
            "post": {
                "security": [
                    {
                        "SigOut": []
                    }
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete role record by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "List User",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        },
        "/user/retrieve/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Retrieve User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.SigIn": {
            "type": "object",
            "required": [
                "password",
                "user"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "domain.UserListRequest": {
            "type": "object",
            "required": [
                "page_num",
                "page_size"
            ],
            "properties": {
                "date_type": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "page_num": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                }
            }
        },
        "model.SysUser": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "group_name": {
                    "description": "角色名",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nick_name": {
                    "description": "暱称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-template",
	Description:      "description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
