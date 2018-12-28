// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-29 00:06:54.313389783 +0800 CST m=+0.039124023

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "118.24.155.105:4000",
    "basePath": "/api/v1",
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
                    "200": {}
                }
            }
        },
        "/qiniu/file": {
            "get": {
                "description": "获取指定存储空间的文件列表，带分页。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "获取指定空间的文件列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "镜像空间名",
                        "name": "bucket",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "页数",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.FileModel"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "支持任何格式的文件上传",
                "consumes": [
                    "multipart/form-data"
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
                        "name": "bucket",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "选择文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            },
            "delete": {
                "description": "删除指定空间的文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "删除指定空间的文件",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/qiniu/file/changeMime": {
            "put": {
                "description": "更新文件类型",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "更新文件类型",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "文件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/qiniu/file/detail": {
            "get": {
                "description": "获取文件详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "获取文件详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.FileModel"
                        }
                    }
                }
            }
        },
        "/qiniu/file/edit": {
            "put": {
                "description": "更新文件信息",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qiniu"
                ],
                "summary": "更新文件信息",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UpdateDetailModel"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "model.FileModel": {
            "type": "object",
            "required": [
                "bucket",
                "key",
                "name",
                "size",
                "type"
            ],
            "properties": {
                "bucket": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.UpdateDetailModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "xx.png"
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
