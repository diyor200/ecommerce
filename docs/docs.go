// Code generated by swaggo/swag. DO NOT EDIT.

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/add/product": {
            "post": {
                "description": "This API endpoint is used to add a new product to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Add a new product",
                "parameters": [
                    {
                        "description": "Product Name and Price",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AddProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Muvaffaqiyatli qo'shildi!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "message: Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/buy": {
            "post": {
                "description": "This API endpoint is used to buy a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Buy a product",
                "parameters": [
                    {
                        "description": "Product ID and user's firstname and lastname",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BuyRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Muvaffaqiyatli harid qilindi!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "message: Harid qilishda muammo yuzga keldi, qaytadan urinib ko'ring !",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "message: Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/delete/{id}": {
            "delete": {
                "description": "Delete a product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete a Product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error"
                    }
                }
            }
        },
        "/api/download": {
            "get": {
                "description": "DownloadExcel creates an excel file with all the orders data and sends it as a response to the client.",
                "summary": "DownloadExcel returns an excel file with all the orders data.",
                "responses": {}
            }
        },
        "/api/products": {
            "get": {
                "description": "Get a list of all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get Products",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error"
                    }
                }
            }
        },
        "/api/products/{id}": {
            "post": {
                "description": "Update an existing product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated product details",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Muvaffaqiyatli o'zgartirildi!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "message: Bunday id dagi product yo'q!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AddProductRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "controllers.BuyRequestBody": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Ecommerce Web api",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
