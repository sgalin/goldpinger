// Copyright 2018 Bloomberg Finance L.P.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
  "swagger": "2.0",
  "info": {
    "title": "Goldpinger",
    "version": "1.0.0"
  },
  "paths": {
    "/check": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint",
        "produces": [
          "application/json"
        ],
        "operationId": "checkServicePods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckResults"
            }
          }
        }
      }
    },
    "/check_all": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.",
        "produces": [
          "application/json"
        ],
        "operationId": "checkAllPods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckAllResults"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "return query stats",
        "produces": [
          "application/json"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "return success",
            "schema": {
              "$ref": "#/definitions/PingResults"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CallStats": {
      "properties": {
        "check": {
          "type": "integer"
        },
        "check_all": {
          "type": "integer"
        },
        "ping": {
          "type": "integer"
        }
      }
    },
    "CheckAllPodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/CheckResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "CheckAllResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "hosts": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "hostIP": {
                "type": "string",
                "format": "ipv4"
              },
              "podIP": {
                "type": "string",
                "format": "ipv4"
              }
            }
          }
        },
        "hosts-healthy": {
          "type": "integer",
          "format": "int32"
        },
        "hosts-number": {
          "type": "integer",
          "format": "int32"
        },
        "responses": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/CheckAllPodResult"
          }
        }
      }
    },
    "CheckResults": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/PodResult"
      }
    },
    "PingResults": {
      "type": "object",
      "properties": {
        "boot_time": {
          "type": "string",
          "format": "date-time"
        },
        "received": {
          "$ref": "#/definitions/CallStats"
        }
      }
    },
    "PodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/PingResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Goldpinger",
    "version": "1.0.0"
  },
  "paths": {
    "/check": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint",
        "produces": [
          "application/json"
        ],
        "operationId": "checkServicePods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckResults"
            }
          }
        }
      }
    },
    "/check_all": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.",
        "produces": [
          "application/json"
        ],
        "operationId": "checkAllPods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckAllResults"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "return query stats",
        "produces": [
          "application/json"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "return success",
            "schema": {
              "$ref": "#/definitions/PingResults"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CallStats": {
      "properties": {
        "check": {
          "type": "integer"
        },
        "check_all": {
          "type": "integer"
        },
        "ping": {
          "type": "integer"
        }
      }
    },
    "CheckAllPodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/CheckResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "CheckAllResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "hosts": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "hostIP": {
                "type": "string",
                "format": "ipv4"
              },
              "podIP": {
                "type": "string",
                "format": "ipv4"
              }
            }
          }
        },
        "hosts-healthy": {
          "type": "integer",
          "format": "int32"
        },
        "hosts-number": {
          "type": "integer",
          "format": "int32"
        },
        "responses": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/CheckAllPodResult"
          }
        }
      }
    },
    "CheckResults": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/PodResult"
      }
    },
    "PingResults": {
      "type": "object",
      "properties": {
        "boot_time": {
          "type": "string",
          "format": "date-time"
        },
        "received": {
          "$ref": "#/definitions/CallStats"
        }
      }
    },
    "PodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/PingResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
}