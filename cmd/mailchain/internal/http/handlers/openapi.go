// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

// nolint: lll
func spec() string {
  return `
{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "All the information needed to talk to the API.\n\nTo raise see anything wrong? Raise an [issue](https://github.com/mailchain/mailchain/issues)",
    "title": "Mailchain API",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/addresses": {
      "get": {
        "description": "Get all address that this user has access to. The addresses can be used to send or receive messages.",
        "tags": [
          "Addresses"
        ],
        "summary": "Get addresses.",
        "operationId": "GetAddresses",
        "responses": {
          "200": {
            "$ref": "#/responses/GetAddressesResponse"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      }
    },
    "/ethereum/{network}/address/{address}/messages": {
      "get": {
        "description": "Get mailchain messages.",
        "tags": [
          "Messages",
          "Ethereum"
        ],
        "summary": "Get Messages.",
        "operationId": "GetMessages",
        "parameters": [
          {
            "pattern": "0x[a-fA-F0-9]{40}",
            "type": "string",
            "example": "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
            "x-go-name": "Address",
            "description": "address to query",
            "name": "address",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "mainnet",
              "ropsten",
              "rinkeby",
              "local"
            ],
            "type": "string",
            "example": "ropsten",
            "x-go-name": "Network",
            "description": "Network",
            "name": "network",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/GetMessagesResponse"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      }
    },
    "/ethereum/{network}/address/{address}/public-key": {
      "get": {
        "description": "Get the public key.",
        "tags": [
          "PublicKey",
          "Ethereum"
        ],
        "summary": "Get public key from an address.",
        "operationId": "GetPublicKey",
        "parameters": [
          {
            "pattern": "0x[a-fA-F0-9]{40}",
            "type": "string",
            "example": "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
            "x-go-name": "Address",
            "description": "address to query to get public key for",
            "name": "address",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "mainnet",
              "ropsten",
              "rinkeby",
              "local"
            ],
            "type": "string",
            "example": "ropsten",
            "x-go-name": "Network",
            "description": "Network for the message to send",
            "name": "network",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/GetPublicKeyResponse"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      }
    },
    "/ethereum/{network}/messages/send": {
      "post": {
        "description": "Securely send message to ethereum address that can only be discovered and de-cryted by the private key holder.\n\nCreate mailchain message\nEncrypt content with public key\nStore message\nEncrypt location\nStore encrypted location on the blockchain.",
        "tags": [
          "Send",
          "Ethereum"
        ],
        "summary": "Send message.",
        "operationId": "SendMessage",
        "parameters": [
          {
            "enum": [
              "mainnet",
              "ropsten",
              "rinkeby",
              "local"
            ],
            "type": "string",
            "example": "ropsten",
            "x-go-name": "Network",
            "description": "Network",
            "name": "network",
            "in": "path",
            "required": true
          },
          {
            "description": "Message to send",
            "name": "PostRequestBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SendMessageRequestBody"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/StatusOK"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      }
    },
    "/messages/{message_id}/read": {
      "get": {
        "tags": [
          "Messages"
        ],
        "summary": "Get message read status.",
        "operationId": "GetRead",
        "responses": {
          "200": {
            "$ref": "#/responses/GetReadResponse"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      },
      "put": {
        "description": "Put encrypted input of all mailchain messages.",
        "tags": [
          "Messages"
        ],
        "summary": "Put inputs.",
        "operationId": "PutRead",
        "parameters": [
          {
            "type": "string",
            "x-go-name": "MessageID",
            "description": "Unique id of the message",
            "name": "message_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/StatusOK"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      },
      "delete": {
        "description": "Mark message as unread",
        "tags": [
          "Messages"
        ],
        "operationId": "DeleteRead",
        "parameters": [
          {
            "type": "string",
            "x-go-name": "MessageID",
            "description": "Unique id of the message",
            "name": "message_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/StatusOK"
          },
          "404": {
            "$ref": "#/responses/NotFoundError"
          },
          "422": {
            "$ref": "#/responses/ValidationError"
          }
        }
      }
    }
  },
  "definitions": {
    "GetMessagesResponseHeaders": {
      "type": "object",
      "properties": {
        "date": {
          "description": "When the message was created, this can be different to the transaction data of the message.",
          "type": "string",
          "format": "date-time",
          "x-go-name": "Date",
          "readOnly": true,
          "example": "12 Mar 19 20:23 UTC"
        },
        "from": {
          "description": "The sender of the message",
          "type": "string",
          "x-go-name": "From",
          "readOnly": true,
          "example": "Charlotte \u003c5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761@ropsten.ethereum\u003e"
        },
        "message-id": {
          "description": "Unique identifier of the message",
          "type": "string",
          "x-go-name": "MessageID",
          "readOnly": true,
          "example": "002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471@mailchain"
        },
        "reply-to": {
          "description": "Reply to if the reply address is different to the from address.",
          "type": "string",
          "x-go-name": "ReplyTo",
          "readOnly": true
        },
        "to": {
          "description": "The recipient of the message",
          "type": "string",
          "x-go-name": "To",
          "readOnly": true
        }
      },
      "x-go-name": "getHeaders",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "GetMessagesResponseMessage": {
      "type": "object",
      "properties": {
        "body": {
          "description": "Body of the mail message",
          "type": "string",
          "x-go-name": "Body",
          "readOnly": true,
          "example": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur maximus metus ante, sit amet ullamcorper dui hendrerit ac."
        },
        "headers": {
          "$ref": "#/definitions/GetMessagesResponseHeaders"
        },
        "read": {
          "description": "Read status of the message",
          "type": "boolean",
          "x-go-name": "Read",
          "readOnly": true,
          "example": true
        },
        "status": {
          "type": "string",
          "x-go-name": "Status",
          "readOnly": true
        },
        "status-code": {
          "type": "string",
          "x-go-name": "StatusCode",
          "readOnly": true
        },
        "subject": {
          "description": "Subject of the mail message",
          "type": "string",
          "x-go-name": "Subject",
          "readOnly": true,
          "example": "Hello world"
        }
      },
      "x-go-name": "getMessage",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "GetPublicKeyResponseBody": {
      "description": "GetBody body response",
      "type": "object",
      "required": [
        "public_key"
      ],
      "properties": {
        "public_key": {
          "description": "The public key",
          "type": "string",
          "x-go-name": "PublicKey",
          "example": "0x79964e63752465973b6b3c610d8ac773fc7ce04f5d1ba599ba8768fb44cef525176f81d3c7603d5a2e466bc96da7b2443bef01b78059a98f45d5c440ca379463"
        }
      },
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "GetReadResponseBody": {
      "type": "object",
      "required": [
        "read"
      ],
      "properties": {
        "read": {
          "description": "Read",
          "type": "boolean",
          "x-go-name": "Read",
          "example": true
        }
      },
      "x-go-name": "getBody",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "PostMessagesResponseHeaders": {
      "type": "object",
      "required": [
        "from",
        "to"
      ],
      "properties": {
        "from": {
          "description": "The sender of the message",
          "type": "string",
          "x-go-name": "From",
          "example": "Charlotte \u003c5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761@ropsten.ethereum\u003e"
        },
        "reply-to": {
          "description": "Reply to if the reply address is different to the from address.",
          "type": "string",
          "x-go-name": "ReplyTo"
        },
        "to": {
          "description": "The recipient of the message",
          "type": "string",
          "x-go-name": "To"
        }
      },
      "x-go-name": "PostHeaders",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "PostMessagesResponseMessage": {
      "type": "object",
      "required": [
        "headers",
        "body",
        "subject",
        "public-key"
      ],
      "properties": {
        "body": {
          "description": "Body of the mail message",
          "type": "string",
          "x-go-name": "Body",
          "example": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur maximus metus ante,"
        },
        "headers": {
          "$ref": "#/definitions/PostMessagesResponseHeaders"
        },
        "public-key": {
          "description": "Public key of the recipient to encrypt with",
          "type": "string",
          "x-go-name": "PublicKey"
        },
        "subject": {
          "description": "Subject of the mail message",
          "type": "string",
          "x-go-name": "Subject",
          "example": "Hello world"
        }
      },
      "x-go-name": "PostMessage",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    },
    "SendMessageRequestBody": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "$ref": "#/definitions/PostMessagesResponseMessage"
        }
      },
      "x-go-name": "PostRequestBody",
      "x-go-package": "github.com/mailchain/mailchain/cmd/mailchain/internal/http/handlers"
    }
  },
  "responses": {
    "GetAddressesResponse": {
      "description": "GetAddressesResponse Holds the response messages",
      "schema": {
        "type": "array",
        "items": {
          "type": "string"
        }
      }
    },
    "GetMessagesResponse": {
      "description": "GetResponse Holds the response messages",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/GetMessagesResponseMessage"
        }
      }
    },
    "GetPublicKeyResponse": {
      "description": "GetPublicKeyResponse public key from address response",
      "schema": {
        "$ref": "#/definitions/GetPublicKeyResponseBody"
      }
    },
    "GetReadResponse": {
      "schema": {
        "$ref": "#/definitions/GetReadResponseBody"
      }
    },
    "NotFoundError": {
      "description": "Description of an error.",
      "headers": {
        "code": {
          "type": "string",
          "example": "404",
          "description": "Code describing the error"
        },
        "message": {
          "type": "string",
          "example": "Not found.",
          "description": "Description of the error"
        }
      }
    },
    "StatusOK": {
      "description": "StatusOK Description of an StatusOK."
    },
    "ValidationError": {
      "description": "Description of an error.",
      "headers": {
        "code": {
          "type": "string",
          "example": "422",
          "description": "Code describing the error"
        },
        "message": {
          "type": "string",
          "example": "Response to invalid input",
          "description": "Description of the error"
        }
      }
    }
  }
}
`
}
