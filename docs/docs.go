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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Geiwoyiquanbei",
            "url": "http://www.swagger.io/support",
            "email": "914832432@qq.com"
        },
        "license": {
            "name": "Api_support",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可创建帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "创建帖子相关接口"
                ],
                "summary": "创建帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "author_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "create_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._ResponseSuccessful"
                        }
                    }
                }
            }
        },
        "/post/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按帖子id号来查询帖子内容",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子详细内容",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "score",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按页码来查询帖子内容",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "score",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "score",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._ResponsePostList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeUserExist",
                "CodeUserNotExist",
                "CodeInvalidPassword",
                "CodeServerBusy",
                "CodeInvalidToken",
                "CodeneedLogin"
            ]
        },
        "controllers._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controllers.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据内容",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "msg": {
                    "description": "信息",
                    "type": "string"
                }
            }
        },
        "controllers._ResponseSuccessful": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controllers.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据内容"
                },
                "msg": {
                    "description": "信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "authuor_name": {
                    "type": "string"
                },
                "community_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "votes": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "这里写版本号",
	Host:             "127.0.0.1:8084",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Gin_web_app",
	Description:      "网络论坛",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
