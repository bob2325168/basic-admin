// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/account/add": {
            "post": {
                "description": "创建账户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "创建账户",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/account/del": {
            "post": {
                "description": "删除账户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "删除账户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/account/detail": {
            "get": {
                "description": "获取账户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "获取账户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/account/edit": {
            "post": {
                "description": "编辑账户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "编辑账户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/account/list": {
            "get": {
                "description": "获取账户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "获取账户列表",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/add": {
            "post": {
                "description": "创建角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "创建角色",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/del": {
            "post": {
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "删除id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/detail": {
            "get": {
                "description": "查看角色详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "查看角色详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/edit": {
            "post": {
                "description": "编辑角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "编辑角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/list": {
            "get": {
                "description": "查看角色列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "查看角色列表",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-api/del": {
            "post": {
                "description": "删除api接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "接口管理"
                ],
                "summary": "删除api接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-api/edit": {
            "post": {
                "description": "更新api接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "接口管理"
                ],
                "summary": "更新api接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-api/list": {
            "get": {
                "description": "api接口列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "接口管理"
                ],
                "summary": "api接口列表",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-opera-log/del": {
            "post": {
                "description": "删除操作记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "删除操作记录",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-opera-log/detail": {
            "get": {
                "description": "操作记录详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作记录详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sys-opera-log/list": {
            "get": {
                "description": "操作记录列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作记录列表",
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/detail": {
            "get": {
                "description": "获取用户的详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "获取用户的详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取用户的详情",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login/address": {
            "post": {
                "description": "钱包登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "钱包登录",
                "parameters": [
                    {
                        "description": "login with param",
                        "name": "loginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginAddressParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login/username": {
            "post": {
                "description": "账密登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "账密登录",
                "parameters": [
                    {
                        "description": "login with param",
                        "name": "loginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/logout": {
            "get": {
                "description": "用户登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": "用户登出成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "registerParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.registerParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register/verify": {
            "get": {
                "description": "验证注册信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "验证注册信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功，请进入登录页面",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.loginAddressParam": {
            "type": "object",
            "required": [
                "address",
                "message",
                "signature"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                }
            }
        },
        "user.loginParam": {
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
        "user.registerParam": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
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
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "BasicAdmin",
	Description: "通用管理后台",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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