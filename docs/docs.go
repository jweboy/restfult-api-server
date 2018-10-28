// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-28 16:29:15.245253555 +0800 CST m=+0.035834656

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is restful api server with golang.",
        "title": "Restful API",
        "termsOfService": "https://github.com/jweboy/restfult-api-server",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "localhost:4000",
    "basePath": "/v1",
    "paths": {
        "/qiniu/bucket": {
            "get": {
                "description": "获取所有的存储空间列表，无分页。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "获取存储空间列表",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"ok\", \"data\": []}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/qiniu/file/{bucketName}": {
            "post": {
                "description": "支持任何格式的文件上传，文件大小有限定",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "文件上传",
                "parameters": [
                    {
                        "type": "string",
                        "description": "存储空间名称",
                        "name": "bucketName",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/user/{username}": {
            "post": {
                "description": "新增用户入库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户名1-32个字符,密码4-128个字符，都必填",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"ok\",\"data\":{\"username\":\"Jack\"}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
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
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
