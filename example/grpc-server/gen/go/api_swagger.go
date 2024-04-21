package api

var SwaggerSpec = `{
  "swagger": "2.0",
  "info": {
    "title": "GRPC-Service",
    "description": "GRPC Service",
    "version": "1.0"
  },
  "tags": [
    {
      "name": "GrpcService"
    }
  ],
  "schemes": [
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v1/todo": {
      "get": {
        "operationId": "GrpcService_GetAllTodo",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/grpc_serverGetAllTodoResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "tags": [
          "GrpcService"
        ]
      },
      "post": {
        "operationId": "GrpcService_CreateTodo",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/grpc_serverTodo"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpc_serverCreateTodoRequest"
            }
          }
        ],
        "tags": [
          "GrpcService"
        ]
      }
    },
    "/api/v1/todo/{id}": {
      "get": {
        "operationId": "GrpcService_GetTodoById",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/grpc_serverTodo"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "GrpcService"
        ]
      }
    },
    "/api/v2/todo": {
      "post": {
        "operationId": "GrpcService_CreateTodoV2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/grpc_serverTodo"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/rpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "todo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpc_serverCreateTodoRequest"
            }
          }
        ],
        "tags": [
          "GrpcService"
        ]
      }
    }
  },
  "definitions": {
    "grpc_serverCreateTodoRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "name"
        },
        "type": {
          "$ref": "#/definitions/grpc_serverTodoType",
          "description": "type"
        }
      },
      "title": "CreateTodo",
      "required": [
        "name",
        "type"
      ]
    },
    "grpc_serverGetAllTodoResponse": {
      "type": "object",
      "properties": {
        "todos": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/grpc_serverTodo"
          }
        }
      },
      "title": "GetAllTodo"
    },
    "grpc_serverTodo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "name"
        },
        "type": {
          "$ref": "#/definitions/grpc_serverTodoType",
          "description": "type"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "description": "creation time"
        }
      }
    },
    "grpc_serverTodoType": {
      "type": "string",
      "enum": [
        "Task",
        "Event"
      ],
      "default": "Task"
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        }
      },
      "additionalProperties": {}
    },
    "rpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    }
  },
  "externalDocs": {
    "description": "Description"
  }
}
`
