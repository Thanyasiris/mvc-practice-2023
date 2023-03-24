package view

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "employee",
      "description": "Testing of MVC's structure"
    }
  ],
  "paths": {
    "/getAllEmployee": {
      "get": {
        "tags": [
          "employee"
        ],
        "summary": "list all employee",
        "responses": {
          "200": {
            "description": "StatusOK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/EmployeeResult"
                }
              }
            }
          }
        }
      }
    },
    "/createEmployee": {
      "post": {
        "tags": [
          "employee"
        ],
        "summary": "create employee",
        "requestBody": {
          "description": "create employee information",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Employee"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Create Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/EmployeeResult"
                }
              }
            }
          }
        }
      }
    },
    "/editEmployee": {
      "post": {
        "tags": [
          "employee"
        ],
        "summary": "edit employee",
        "requestBody": {
          "description": "edit employee information",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/Employee"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "edit Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/EmployeeResult"
                }
              }
            }
          }
        }
      }
    },
    "/deleteEmployee": {
      "post": {
        "tags": [
          "employee"
        ],
        "summary": "delete employee",
        "requestBody": {
          "description": "delete employee information",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/DeleteEmp"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "delete Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/EmployeeResult"
                }
              }
            }
          }
        }
      }
    },
    "/createBird": {
      "post": {
        "tags": [
          "bird"
        ],
        "summary": "create random bird in steel cage",
        "responses": {
          "200": {
            "description": "Create Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Bird"
                }
              }
            }
          }
        }
      }
    },
  },
  "components": {
    "schemas": {
      "Employee": {
        "type": "object",
        "properties": {
          "emp_no": {
            "type": "integer",
            "format": "int64"
          },
          "first_name": {
            "type": "string",
            "example": "Armin"
          },
          "last_name": {
            "type": "string",
            "example": "Arlert"
          },
          "position": {
            "type": "string",
          },
          "education": {
            "type": "string",
            "example": "mocha"
          },
          "dept_no": {
            "type": "integer",
            "format": "int64"
          }
        }
      },
      "EmployeeResult": {
        "type": "object",
        "properties": {
          "emp_no": {
            "type": "integer",
            "format": "int64"
          },
          "first_name": {
            "type": "string",
            "example": "Armin"
          },
          "last_name": {
            "type": "string",
            "example": "Arlert"
          },
          "position": {
            "type": "string",
          },
          "education": {
            "type": "string",
            "example": "mocha"
          },
          "dept_no": {
            "type": "integer",
            "format": "int64"
          }
        }
      },
      "DeleteEmp": {
        "type": "object",
        "properties": {
          "emp_no": {
            "type": "integer",
            "format": "int64"
          }
        }
      },
      "Bird": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "format": "int64"
          },
          "color": {
            "type": "string",
            "example": "blue"
          },
          "typee": {
            "type": "string",
            "example": "RET"
          },
          "steel_id": {
            "type": "integer",
            "format": "int64"
          },
          "sink_id": {
            "type": "integer",
            "format": "int64"
          }
        }
      },
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Party's Applicant",
	Description:      "MVC Example 01",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
