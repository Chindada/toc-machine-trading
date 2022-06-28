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
        "/analyze/below-quater": {
            "get": {
                "description": "getQuaterTargets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "analyze"
                ],
                "summary": "getQuaterTargets",
                "operationId": "getQuaterTargets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1.BelowQuaterMA"
                            }
                        }
                    }
                }
            }
        },
        "/basic/stock/repo": {
            "get": {
                "description": "getAllRepoStock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "basic"
                ],
                "summary": "getAllRepoStock",
                "operationId": "getAllRepoStock",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.stockDetailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/basic/stock/sinopac-to-repo": {
            "get": {
                "description": "getAllSinopacStockAndUpdateRepo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "basic"
                ],
                "summary": "getAllSinopacStockAndUpdateRepo",
                "operationId": "getAllSinopacStockAndUpdateRepo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.stockDetailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/basic/system/terminate": {
            "put": {
                "description": "terminateSinopac",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "terminateSinopac",
                "operationId": "terminateSinopac",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/history/day_kbar/{stock}/{start_date}/{interval}": {
            "get": {
                "description": "getKbarData",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "getKbarData",
                "operationId": "getKbarData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "stock",
                        "name": "stock",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "start_date",
                        "name": "start_date",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "interval",
                        "name": "interval",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.HistoryKbar"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "description": "getAllOrder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "getAllOrder",
                "operationId": "getAllOrder",
                "responses": {}
            }
        },
        "/stream": {
            "get": {
                "description": "getTSESnapshot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stream"
                ],
                "summary": "getTSESnapshot",
                "operationId": "getTSESnapshot",
                "responses": {}
            }
        },
        "/targets": {
            "get": {
                "description": "getTargets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "target"
                ],
                "summary": "getTargets",
                "operationId": "getTargets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Target"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.HistoryKbar": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "number"
                },
                "high": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "kbar_time": {
                    "type": "string"
                },
                "low": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "stock": {
                    "$ref": "#/definitions/entity.Stock"
                },
                "stock_num": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "entity.Stock": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "day_trade": {
                    "type": "boolean"
                },
                "exchange": {
                    "type": "string"
                },
                "last_close": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "entity.Target": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "rank": {
                    "type": "integer"
                },
                "real_time_add": {
                    "type": "boolean"
                },
                "stock": {
                    "$ref": "#/definitions/entity.Stock"
                },
                "stock_num": {
                    "type": "string"
                },
                "subscribe": {
                    "type": "boolean"
                },
                "trade_day": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "v1.BelowQuaterMA": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Stock"
                    }
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "v1.stockDetailResponse": {
            "type": "object",
            "properties": {
                "stock_detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Stock"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TOC MACHINE TRADING API",
	Description:      "Auto Trade on sinopac",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
