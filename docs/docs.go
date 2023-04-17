// Code generated by swaggo/swag. DO NOT EDIT
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
        "/aggregate": {
            "get": {
                "description": "原始数据按时间统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "原始数据按时间统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/area": {
            "get": {
                "description": "按区域统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按区域统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据点",
                        "name": "point",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/day": {
            "get": {
                "description": "按日统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按日统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/group": {
            "get": {
                "description": "按分组统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按分组统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据点",
                        "name": "point",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/hour": {
            "get": {
                "description": "按小时统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按小时统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/minute": {
            "get": {
                "description": "按分钟统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按分钟统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/month": {
            "get": {
                "description": "按月统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按月统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/type": {
            "get": {
                "description": "按类型统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按类型统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据点",
                        "name": "point",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/aggregate/year": {
            "get": {
                "description": "按年统计",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregate"
                ],
                "summary": "按年统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设备区域",
                        "name": "area",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备分组",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "数据点位",
                        "name": "point",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "起始时间",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "设备类型",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-api_Result"
                        }
                    }
                }
            }
        },
        "/history/count": {
            "post": {
                "description": "查询历史",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "查询历史数量",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-int64"
                        }
                    }
                }
            }
        },
        "/history/list": {
            "get": {
                "description": "查询历史",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "查询历史",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyList-types_History"
                        }
                    }
                }
            }
        },
        "/history/search": {
            "post": {
                "description": "查询历史",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "查询历史",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyList-types_History"
                        }
                    }
                }
            }
        },
        "/history/{id}/delete": {
            "get": {
                "description": "删除历史",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "删除历史",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "历史ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_History"
                        }
                    }
                }
            }
        },
        "/job/count": {
            "post": {
                "description": "查询计划任务数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "查询计划任务数量",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-int64"
                        }
                    }
                }
            }
        },
        "/job/create": {
            "post": {
                "description": "创建计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "创建计划任务",
                "parameters": [
                    {
                        "description": "计划任务信息",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Job"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            }
        },
        "/job/export": {
            "get": {
                "description": "导出计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "product"
                ],
                "summary": "导出计划任务",
                "responses": {}
            }
        },
        "/job/import": {
            "post": {
                "description": "导入计划任务",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "导入计划任务",
                "parameters": [
                    {
                        "type": "file",
                        "description": "压缩包",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-int64"
                        }
                    }
                }
            }
        },
        "/job/list": {
            "get": {
                "description": "查询计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "查询计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyList-types_Job"
                        }
                    }
                }
            }
        },
        "/job/search": {
            "post": {
                "description": "查询计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "查询计划任务",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ParamSearch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyList-types_Job"
                        }
                    }
                }
            }
        },
        "/job/{id}": {
            "get": {
                "description": "获取计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "获取计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "计划任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            },
            "post": {
                "description": "修改计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "修改计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "计划任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "计划任务信息",
                        "name": "job",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Job"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            }
        },
        "/job/{id}/delete": {
            "get": {
                "description": "删除计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "删除计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "计划任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            }
        },
        "/job/{id}/disable": {
            "get": {
                "description": "禁用计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "禁用计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "计划任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            }
        },
        "/job/{id}/enable": {
            "get": {
                "description": "启用计划任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "启用计划任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "计划任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ReplyData-types_Job"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ParamSearch": {
            "type": "object",
            "properties": {
                "filter": {
                    "type": "object",
                    "additionalProperties": true
                },
                "keyword": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "skip": {
                    "type": "integer"
                },
                "sort": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "api.ReplyData-api_Result": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/api.Result"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ReplyData-int64": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ReplyData-types_History": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/types.History"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ReplyData-types_Job": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/types.Job"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ReplyList-types_History": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.History"
                    }
                },
                "error": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "api.ReplyList-types_Job": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Job"
                    }
                },
                "error": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "api.Result": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "types.History": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "point": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "types.Job": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "crontab": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "points": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Aggregator"
                    }
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "types.Aggregator": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "description": "store increase average count",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0 版本",
	Host:             "",
	BasePath:         "/app/history/api/",
	Schemes:          []string{},
	Title:            "历史数据接口文档",
	Description:      "API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
