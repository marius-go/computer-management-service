// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "computer management",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/computer": {
      "get": {
        "tags": [
          "computer"
        ],
        "summary": "list computers",
        "operationId": "listComputers",
        "parameters": [
          {
            "type": "string",
            "description": "filter computer by employee abbreviation",
            "name": "employeeAbbreviation",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/computer"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "computer"
        ],
        "summary": "update an existing computer",
        "operationId": "updateComputer",
        "parameters": [
          {
            "description": "desired new state of Computer",
            "name": "computer",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/computer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "updated",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "400": {
            "$ref": "#/responses/badRequest"
          },
          "404": {
            "$ref": "#/responses/notFound"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "computer"
        ],
        "summary": "create a new computer entry",
        "operationId": "createComputer",
        "parameters": [
          {
            "description": "computer object to be added to the service",
            "name": "computer",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newComputer"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "400": {
            "$ref": "#/responses/badRequest"
          },
          "409": {
            "description": "already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/computer/{computerName}": {
      "get": {
        "tags": [
          "computer"
        ],
        "summary": "find computer by name",
        "operationId": "getComputer",
        "parameters": [
          {
            "type": "string",
            "description": "name of the computer",
            "name": "computerName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "404": {
            "$ref": "#/responses/notFound"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "computer"
        ],
        "summary": "remove computer from the service",
        "operationId": "deleteComputer",
        "parameters": [
          {
            "type": "string",
            "description": "name of the computer to be removed from the service",
            "name": "computerName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "removed"
          },
          "404": {
            "$ref": "#/responses/notFound"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "computer": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "employeeAbbreviation": {
          "type": "string"
        },
        "ip": {
          "type": "string"
        },
        "mac": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "newComputer": {
      "allOf": [
        {
          "$ref": "#/definitions/computer"
        },
        {
          "type": "object",
          "required": [
            "mac",
            "ip"
          ]
        }
      ]
    }
  },
  "responses": {
    "badRequest": {
      "description": "bad request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "notFound": {
      "description": "the specified resource was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "computer management",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/computer": {
      "get": {
        "tags": [
          "computer"
        ],
        "summary": "list computers",
        "operationId": "listComputers",
        "parameters": [
          {
            "type": "string",
            "description": "filter computer by employee abbreviation",
            "name": "employeeAbbreviation",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/computer"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "computer"
        ],
        "summary": "update an existing computer",
        "operationId": "updateComputer",
        "parameters": [
          {
            "description": "desired new state of Computer",
            "name": "computer",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/computer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "updated",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "the specified resource was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "computer"
        ],
        "summary": "create a new computer entry",
        "operationId": "createComputer",
        "parameters": [
          {
            "description": "computer object to be added to the service",
            "name": "computer",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newComputer"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/computer/{computerName}": {
      "get": {
        "tags": [
          "computer"
        ],
        "summary": "find computer by name",
        "operationId": "getComputer",
        "parameters": [
          {
            "type": "string",
            "description": "name of the computer",
            "name": "computerName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/computer"
            }
          },
          "404": {
            "description": "the specified resource was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "computer"
        ],
        "summary": "remove computer from the service",
        "operationId": "deleteComputer",
        "parameters": [
          {
            "type": "string",
            "description": "name of the computer to be removed from the service",
            "name": "computerName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "removed"
          },
          "404": {
            "description": "the specified resource was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "NewComputerAllOf1": {
      "type": "object",
      "required": [
        "mac",
        "ip"
      ]
    },
    "computer": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "employeeAbbreviation": {
          "type": "string"
        },
        "ip": {
          "type": "string"
        },
        "mac": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "newComputer": {
      "allOf": [
        {
          "$ref": "#/definitions/computer"
        },
        {
          "$ref": "#/definitions/NewComputerAllOf1"
        }
      ]
    }
  },
  "responses": {
    "badRequest": {
      "description": "bad request",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "notFound": {
      "description": "the specified resource was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
}
