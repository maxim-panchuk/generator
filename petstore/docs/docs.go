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
        "/pet": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Update an existing pet by Id\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Update an existing pet\"",
                "parameters": [
                    {
                        "description": "Update an existing pet by Id",
                        "name": "Pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pet.PetDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/pet.PetDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied"
                    },
                    "404": {
                        "description": "Pet not found"
                    },
                    "422": {
                        "description": "Validation exception"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Add a new pet to the store\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Add a new pet to the store\"",
                "parameters": [
                    {
                        "description": "Add a new pet to the store",
                        "name": "Pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pet.PetDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/pet.PetDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input"
                    },
                    "422": {
                        "description": "Validation exception"
                    }
                }
            }
        },
        "/pet/findByStatus": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Multiple status values can be provided with comma separated strings\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Finds Pets by status\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Status values that need to be considered for filter",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pet.PetDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid status value"
                    }
                }
            }
        },
        "/pet/findByTags": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Finds Pets by tags\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tags to filter by",
                        "name": "tags",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pet.PetDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid tag value"
                    }
                }
            }
        },
        "/pet/{petId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Returns a single pet\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Find pet by ID\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of pet to return",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/pet.PetDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied"
                    },
                    "404": {
                        "description": "Pet not found"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"default description\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Updates a pet in the store with form data\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of pet that needs to be updated",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of pet that needs to be updated",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Status of pet that needs to be updated",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Invalid input"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"delete a pet\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"Deletes a pet\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pet id to delete",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Invalid pet value"
                    }
                }
            }
        },
        "/pet/{petId}/uploadImage": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"default description\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "pet"
                ],
                "summary": "\"uploads an image\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of pet to update",
                        "name": "petId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Additional Metadata",
                        "name": "additionalMetadata",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/apiResponse.ApiResponseDTO"
                        }
                    }
                }
            }
        },
        "/store/inventory": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Returns a map of status codes to quantities\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "\"Returns pet inventories by status\"",
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/store/order": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Place a new order in the store\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "\"Place an order for a pet\"",
                "parameters": [
                    {
                        "description": "Place a new order in the store",
                        "name": "Order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.OrderDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/order.OrderDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input"
                    },
                    "422": {
                        "description": "Validation exception"
                    }
                }
            }
        },
        "/store/order/{orderId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"For valid response try integer IDs with value \u003c= 5 or \u003e 10. Other values will generate exceptions.\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "\"Find purchase order by ID\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of order that needs to be fetched",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/order.OrderDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid ID supplied"
                    },
                    "404": {
                        "description": "Order not found"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"For valid response try integer IDs with value \u003c 1000. Anything above 1000 or nonintegers will generate API errors\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "\"Delete purchase order by ID\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the order that needs to be deleted",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Invalid ID supplied"
                    },
                    "404": {
                        "description": "Order not found"
                    }
                }
            }
        },
        "/user": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"This can only be done by the logged in user.\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Create user\"",
                "parameters": [
                    {
                        "description": "This can only be done by the logged in user.",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "default": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    }
                }
            }
        },
        "/user/createWithList": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"Creates list of users with given input array\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Creates list of users with given input array\"",
                "parameters": [
                    {
                        "description": "Creates list of users with given input array",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    },
                    "default": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/user/login": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"default description\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Logs user into the system\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user name for login",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The password for login in clear text",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    },
                    "400": {
                        "description": "Invalid username/password supplied"
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"default description\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Logs out current logged in user session\"",
                "responses": {
                    "default": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/user/{username}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"default description\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Get user by user name\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name that needs to be fetched. Use user1 for testing. ",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid username supplied"
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"This can only be done by the logged in user.\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Update user\"",
                "parameters": [
                    {
                        "description": "This can only be done by the logged in user.",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "name that need to be deleted",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "default": {
                        "description": "successful operation"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "\"This can only be done by the logged in user.\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "\"Delete user\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name that needs to be deleted",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Invalid username supplied"
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "apiResponse.ApiResponseDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "category.CategoryDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "order.OrderDTO": {
            "type": "object",
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "petId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "shipDate": {
                    "type": "string"
                },
                "status": {
                    "description": "Order Status",
                    "type": "string"
                }
            }
        },
        "pet.PetDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/category.CategoryDTO"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photoUrls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "description": "pet status in the store",
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tag.TagDTO"
                    }
                }
            }
        },
        "tag.TagDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.UserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "userStatus": {
                    "description": "User Status",
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "petstore api",
	Description:      "petstore api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
