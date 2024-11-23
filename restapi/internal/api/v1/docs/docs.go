// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "List all categories",
                "produces": [
                    "application/json"
                ],
                "summary": "List asset categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Category"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create an asset category",
                "parameters": [
                    {
                        "description": "Category object",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Category"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Category"
                        }
                    }
                }
            }
        },
        "/categories/{categoryTitle}": {
            "get": {
                "description": "Get a category",
                "produces": [
                    "application/json"
                ],
                "summary": "Get an asset category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Title",
                        "name": "categoryTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Category"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an asset category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Title",
                        "name": "categoryTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category object",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Category"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete a category",
                "summary": "Delete an asset category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Title",
                        "name": "categoryTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/consumables": {
            "get": {
                "description": "List all consumables",
                "produces": [
                    "application/json"
                ],
                "summary": "List consumables",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Consumable"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a consumable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a consumable",
                "parameters": [
                    {
                        "description": "Consumable object",
                        "name": "consumable",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Consumable"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Consumable"
                        }
                    }
                }
            }
        },
        "/consumables/{consumableTitle}": {
            "get": {
                "description": "Get a consumable",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a consumable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Consumable ID",
                        "name": "consumableTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Consumable"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a consumable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a consumable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Consumable ID",
                        "name": "consumableTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Consumable object",
                        "name": "consumable",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Consumable"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Consumable"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a consumable",
                "summary": "Delete a consumable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Consumable ID",
                        "name": "consumableTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/correctiveTasks": {
            "get": {
                "description": "List all correctiveTasks",
                "produces": [
                    "application/json"
                ],
                "summary": "List correctiveTasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.CorrectiveTask"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a correctiveTask",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a correctiveTask",
                "parameters": [
                    {
                        "description": "CorrectiveTask object",
                        "name": "correctiveTask",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CorrectiveTask"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.CorrectiveTask"
                        }
                    }
                }
            }
        },
        "/correctiveTasks/{correctiveTaskTitle}": {
            "get": {
                "description": "Get a correctiveTask",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a correctiveTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CorrectiveTask Title",
                        "name": "correctiveTaskTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CorrectiveTask"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a correctiveTask",
                "consumes": [
                    "application/json"
                ],
                "summary": "Update a correctiveTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CorrectiveTask Title",
                        "name": "correctiveTaskTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a correctiveTask",
                "summary": "Delete a correctiveTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CorrectiveTask Title",
                        "name": "correctiveTaskTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/groups": {
            "get": {
                "description": "List all groups",
                "produces": [
                    "application/json"
                ],
                "summary": "List asset groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Group"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create an asset group",
                "parameters": [
                    {
                        "description": "Group object",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Group"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Group"
                        }
                    }
                }
            }
        },
        "/groups/{groupTitle}": {
            "get": {
                "description": "Get a group",
                "produces": [
                    "application/json"
                ],
                "summary": "Get an asset group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Title",
                        "name": "groupTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Group"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an asset group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Title",
                        "name": "groupTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Group object",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Group"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Group"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a group",
                "summary": "Delete an asset group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Title",
                        "name": "groupTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/preventativeTasks": {
            "get": {
                "description": "List all preventativeTasks",
                "produces": [
                    "application/json"
                ],
                "summary": "List preventativeTasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.PreventativeTask"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a preventativeTask",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a preventativeTask",
                "parameters": [
                    {
                        "description": "PreventativeTask object",
                        "name": "preventativeTask",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.PreventativeTask"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.PreventativeTask"
                        }
                    }
                }
            }
        },
        "/preventativeTasks/{preventativeTaskTitle}": {
            "get": {
                "description": "Get a preventativeTask",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a preventativeTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PreventativeTask Title",
                        "name": "preventativeTaskTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.PreventativeTask"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a preventativeTask",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a preventativeTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PreventativeTask Title",
                        "name": "preventativeTaskTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PreventativeTask object",
                        "name": "preventativeTask",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.PreventativeTask"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "delete": {
                "description": "Delete a preventativeTask",
                "summary": "Delete a preventativeTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PreventativeTask Title",
                        "name": "preventativeTaskTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/time-units": {
            "get": {
                "description": "List all time units",
                "produces": [
                    "application/json"
                ],
                "summary": "List time units",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.TimeUnit"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a time unit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a time unit",
                "parameters": [
                    {
                        "description": "Time Unit object",
                        "name": "timeUnit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TimeUnit"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.TimeUnit"
                        }
                    }
                }
            }
        },
        "/time-units/{timeUnitTitle}": {
            "get": {
                "description": "Get a time unit",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a time unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Time Unit Title",
                        "name": "timeUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TimeUnit"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a time unit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a time unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Time Unit Title",
                        "name": "timeUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TimeUnit"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a time unit",
                "summary": "Delete a time unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Time Unit Title",
                        "name": "timeUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/usage-units": {
            "get": {
                "description": "List all usage units",
                "produces": [
                    "application/json"
                ],
                "summary": "List usage units",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UsageUnit"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a usage unit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a usage unit",
                "parameters": [
                    {
                        "description": "Usage Unit object",
                        "name": "usageUnit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UsageUnit"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.UsageUnit"
                        }
                    }
                }
            }
        },
        "/usage-units/{usageUnitTitle}": {
            "get": {
                "description": "Get a usage unit",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a usage unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Usage Unit Title",
                        "name": "usageUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.UsageUnit"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a usage unit",
                "consumes": [
                    "application/json"
                ],
                "summary": "Update a usage unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Usage Unit Title",
                        "name": "usageUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a usage unit",
                "summary": "Delete a usage unit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Usage Unit Title",
                        "name": "usageUnitTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/work-order-statuses": {
            "get": {
                "description": "List all work order statuses",
                "produces": [
                    "application/json"
                ],
                "summary": "List work order statuses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.WorkOrderStatus"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a work order status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a work order status",
                "parameters": [
                    {
                        "description": "Work Order Status object",
                        "name": "workOrderStatus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.WorkOrderStatus"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.WorkOrderStatus"
                        }
                    }
                }
            }
        },
        "/work-order-statuses/{workOrderStatusTitle}": {
            "get": {
                "description": "Get a work order status",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a work order status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work Order Status Title",
                        "name": "workOrderStatusTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.WorkOrderStatus"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a work order status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a work order status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work Order Status Title",
                        "name": "workOrderStatusTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Work Order Status object",
                        "name": "workOrderStatus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.WorkOrderStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.WorkOrderStatus"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a work order status",
                "summary": "Delete a work order status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work Order Status Title",
                        "name": "workOrderStatusTitle",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Category": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "types.Consumable": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "types.CorrectiveTask": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "types.Group": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "types.PreventativeTask": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "types.TimeUnit": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "types.UsageUnit": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "types.WorkOrderStatus": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "PersonaCMMS API",
	Description:      "This is the Personal Computer Maintenance Management System REST API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
