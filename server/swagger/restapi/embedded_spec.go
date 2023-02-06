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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is RESTful API specification for quocbang system.\nThe default response of each API endpoints are based on HTTP Status Code Standardization, e.g:\n  - 400 (Bad Request)\n  - 401 (Unauthorized)\n  - 403 (Forbidden)\n  - 408 (Request Timeout)\n  - 500 (Internal Server Error)\n",
    "title": "API Proposal for go-swagger-api",
    "contact": {
      "email": "quocbang@kenda.com.tw"
    },
    "license": {
      "name": "Copyright © 2023 by quocbang, All Rights Reserved"
    },
    "version": "1.1.1"
  },
  "basePath": "/api",
  "paths": {
    "/account/authorization": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "- 不允許授權系統管理員(ADMINISTRATOR)與主管(LEADER)角色.\n- 若想要授權系統管理員或主管角色時請再提登入異動單.\n",
        "tags": [
          "account"
        ],
        "summary": "新增帳號權限",
        "operationId": "CreateAccountAuthorization",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "employeeID",
                "roles"
              ],
              "properties": {
                "employeeID": {
                  "description": "人員工號",
                  "type": "string"
                },
                "roles": {
                  "$ref": "#/definitions/Roles"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/account/authorization/{employeeID}": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "- 不允許授權系統管理員(ADMINISTRATOR)與主管(LEADER)角色.\n- 若想要授權系統管理員或主管角色時請再提登入異動單.\n- 允許修改角色清單或者重置密碼需求\n",
        "tags": [
          "account"
        ],
        "summary": "修改帳號角色",
        "operationId": "UpdateAccountAuthorization",
        "parameters": [
          {
            "type": "string",
            "description": "人員工號",
            "name": "employeeID",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "roles",
                "resetPassword"
              ],
              "properties": {
                "resetPassword": {
                  "description": "重置密碼",
                  "type": "boolean"
                },
                "roles": {
                  "$ref": "#/definitions/Roles"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "刪除帳號",
        "operationId": "DeleteAccount",
        "parameters": [
          {
            "type": "string",
            "description": "人員工號",
            "name": "employeeID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/account/authorized/department-oid/{departmentOID}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "查詢帳號權限清單",
        "operationId": "ListAuthorizedAccount",
        "parameters": [
          {
            "$ref": "#/parameters/DepartmentOID"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/AccountData"
                }
              }
            }
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/account/role-list": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "取得角色清單",
        "operationId": "GetRoleList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "ID": {
                        "$ref": "#/definitions/Role"
                      },
                      "name": {
                        "description": "角色名稱",
                        "type": "string",
                        "example": "SCHEDULER"
                      }
                    }
                  }
                }
              }
            }
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/account/unauthorized/department-oid/{departmentOID}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "查詢可新增角色授權帳號清單",
        "operationId": "ListUnauthorizedAccount",
        "parameters": [
          {
            "$ref": "#/parameters/DepartmentOID"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "employeeID": {
                        "description": "人員工號",
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/limitary/{productType}": {
      "get": {
        "security": [],
        "tags": [
          "production"
        ],
        "summary": "get limitary hour from database",
        "operationId": "Getlimitary",
        "parameters": [
          {
            "type": "string",
            "name": "productType",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/Limitary"
                }
              }
            }
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/server/status": {
      "get": {
        "security": [],
        "summary": "檢查server狀態用",
        "operationId": "CheckServerStatus",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/user/change-password": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "修改密碼(限MES帳號)",
        "operationId": "ChangePassword",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "currentPassword",
                "newPassword"
              ],
              "properties": {
                "currentPassword": {
                  "description": "原密碼",
                  "type": "string",
                  "maxLength": 10,
                  "minLength": 6
                },
                "newPassword": {
                  "description": "新密碼",
                  "type": "string",
                  "maxLength": 10,
                  "minLength": 6
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/Default"
          }
        }
      }
    },
    "/user/login": {
      "post": {
        "security": [],
        "description": "支援登入類別為:\n- MES 帳號登入\n- Windows 帳號登入\n",
        "tags": [
          "account"
        ],
        "summary": "使用者登入",
        "operationId": "Login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/LoginResponse"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/logout": {
      "post": {
        "security": [],
        "tags": [
          "account"
        ],
        "summary": "使用者登出",
        "operationId": "Logout",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  },
  "definitions": {
    "AccountData": {
      "type": "array",
      "items": {
        "type": "object",
        "required": [
          "employeeID",
          "roles"
        ],
        "properties": {
          "employeeID": {
            "description": "人員工號",
            "type": "string",
            "example": 0
          },
          "roles": {
            "$ref": "#/definitions/Roles"
          }
        }
      }
    },
    "Department": {
      "description": "部門資訊",
      "type": "object",
      "properties": {
        "ID": {
          "description": "部門代號",
          "type": "string",
          "example": "M2110"
        },
        "OID": {
          "description": "部門OID",
          "type": "string"
        }
      }
    },
    "Departments": {
      "description": "部門清單",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Department"
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "自定義錯誤碼\n定義來源參考:\nhttps://gitlab.kenda.com.tw/kenda/mcom/-/blob/${xxx}/errors/code.proto ` + "`" + `Code` + "`" + ` enum.\nin which ` + "`" + `${xxx}` + "`" + ` inside the URL reference is the specified branch name of the corresponding features.\n",
          "type": "integer",
          "x-omitempty": false
        },
        "details": {
          "description": "補充訊息",
          "type": "string"
        }
      }
    },
    "Limitary": {
      "type": "object",
      "properties": {
        "max": {
          "type": "string",
          "format": "date-time",
          "x-order": 1,
          "readOnly": true
        },
        "min": {
          "type": "string",
          "format": "date-time",
          "x-order": 0,
          "readOnly": true
        }
      }
    },
    "LoginRequest": {
      "required": [
        "ID",
        "password",
        "loginType"
      ],
      "properties": {
        "ID": {
          "description": "使用者工號/帳號",
          "type": "string"
        },
        "loginType": {
          "$ref": "#/definitions/LoginType"
        },
        "password": {
          "description": "使用者密碼",
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "properties": {
        "authorizedDepartments": {
          "$ref": "#/definitions/Departments"
        },
        "roles": {
          "$ref": "#/definitions/Roles"
        },
        "token": {
          "description": "使用者令牌(token)",
          "type": "string"
        },
        "tokenExpiry": {
          "description": "token 過期時間",
          "type": "string",
          "format": "date-time",
          "example": "2017-07-21T17:32:28Z"
        }
      }
    },
    "LoginType": {
      "description": "登入類別:\n  * 0 - MES\n  * 1 - Windows\n",
      "type": "integer",
      "enum": [
        0,
        1
      ],
      "x-omitempty": false
    },
    "Principal": {
      "properties": {
        "ID": {
          "description": "使用者工號/帳號",
          "type": "string",
          "example": "userid"
        },
        "roles": {
          "$ref": "#/definitions/Roles"
        }
      }
    },
    "Role": {
      "description": "授權角色\n定義來源參考: https://gitlab.kenda.com.tw/kenda/mcom/-/blob/${xxx}/utils/roles/roles.proto ` + "`" + `Role` + "`" + ` type.\nin which ` + "`" + `${xxx}` + "`" + ` inside the URL reference is the specified branch name of the corresponding features.\n",
      "type": "integer",
      "x-omitempty": false
    },
    "Roles": {
      "description": "使用者角色",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Role"
      }
    }
  },
  "parameters": {
    "DepartmentOID": {
      "type": "string",
      "description": "部門OID",
      "name": "departmentOID",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "Default": {
      "description": "Unexpected error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "x-api-auth-key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": []
    }
  ],
  "tags": [
    {
      "description": "帳號管理相關",
      "name": "account"
    },
    {
      "name": "product"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is RESTful API specification for quocbang system.\nThe default response of each API endpoints are based on HTTP Status Code Standardization, e.g:\n  - 400 (Bad Request)\n  - 401 (Unauthorized)\n  - 403 (Forbidden)\n  - 408 (Request Timeout)\n  - 500 (Internal Server Error)\n",
    "title": "API Proposal for go-swagger-api",
    "contact": {
      "email": "quocbang@kenda.com.tw"
    },
    "license": {
      "name": "Copyright © 2023 by quocbang, All Rights Reserved"
    },
    "version": "1.1.1"
  },
  "basePath": "/api",
  "paths": {
    "/account/authorization": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "- 不允許授權系統管理員(ADMINISTRATOR)與主管(LEADER)角色.\n- 若想要授權系統管理員或主管角色時請再提登入異動單.\n",
        "tags": [
          "account"
        ],
        "summary": "新增帳號權限",
        "operationId": "CreateAccountAuthorization",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "employeeID",
                "roles"
              ],
              "properties": {
                "employeeID": {
                  "description": "人員工號",
                  "type": "string"
                },
                "roles": {
                  "$ref": "#/definitions/Roles"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/account/authorization/{employeeID}": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "- 不允許授權系統管理員(ADMINISTRATOR)與主管(LEADER)角色.\n- 若想要授權系統管理員或主管角色時請再提登入異動單.\n- 允許修改角色清單或者重置密碼需求\n",
        "tags": [
          "account"
        ],
        "summary": "修改帳號角色",
        "operationId": "UpdateAccountAuthorization",
        "parameters": [
          {
            "type": "string",
            "description": "人員工號",
            "name": "employeeID",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "roles",
                "resetPassword"
              ],
              "properties": {
                "resetPassword": {
                  "description": "重置密碼",
                  "type": "boolean"
                },
                "roles": {
                  "$ref": "#/definitions/Roles"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "刪除帳號",
        "operationId": "DeleteAccount",
        "parameters": [
          {
            "type": "string",
            "description": "人員工號",
            "name": "employeeID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/account/authorized/department-oid/{departmentOID}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "查詢帳號權限清單",
        "operationId": "ListAuthorizedAccount",
        "parameters": [
          {
            "type": "string",
            "description": "部門OID",
            "name": "departmentOID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/AccountData"
                }
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/account/role-list": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "取得角色清單",
        "operationId": "GetRoleList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/DataItems0"
                  }
                }
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/account/unauthorized/department-oid/{departmentOID}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "查詢可新增角色授權帳號清單",
        "operationId": "ListUnauthorizedAccount",
        "parameters": [
          {
            "type": "string",
            "description": "部門OID",
            "name": "departmentOID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/DataItems0"
                  }
                }
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/limitary/{productType}": {
      "get": {
        "security": [],
        "tags": [
          "production"
        ],
        "summary": "get limitary hour from database",
        "operationId": "Getlimitary",
        "parameters": [
          {
            "type": "string",
            "name": "productType",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/Limitary"
                }
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/server/status": {
      "get": {
        "security": [],
        "summary": "檢查server狀態用",
        "operationId": "CheckServerStatus",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/user/change-password": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "account"
        ],
        "summary": "修改密碼(限MES帳號)",
        "operationId": "ChangePassword",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "currentPassword",
                "newPassword"
              ],
              "properties": {
                "currentPassword": {
                  "description": "原密碼",
                  "type": "string",
                  "maxLength": 10,
                  "minLength": 6
                },
                "newPassword": {
                  "description": "新密碼",
                  "type": "string",
                  "maxLength": 10,
                  "minLength": 6
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/login": {
      "post": {
        "security": [],
        "description": "支援登入類別為:\n- MES 帳號登入\n- Windows 帳號登入\n",
        "tags": [
          "account"
        ],
        "summary": "使用者登入",
        "operationId": "Login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/LoginResponse"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/logout": {
      "post": {
        "security": [],
        "tags": [
          "account"
        ],
        "summary": "使用者登出",
        "operationId": "Logout",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  },
  "definitions": {
    "AccountData": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/AccountDataItems0"
      }
    },
    "AccountDataItems0": {
      "type": "object",
      "required": [
        "employeeID",
        "roles"
      ],
      "properties": {
        "employeeID": {
          "description": "人員工號",
          "type": "string",
          "example": 0
        },
        "roles": {
          "$ref": "#/definitions/Roles"
        }
      }
    },
    "DataItems0": {
      "type": "object",
      "properties": {
        "employeeID": {
          "description": "人員工號",
          "type": "string"
        }
      }
    },
    "Department": {
      "description": "部門資訊",
      "type": "object",
      "properties": {
        "ID": {
          "description": "部門代號",
          "type": "string",
          "example": "M2110"
        },
        "OID": {
          "description": "部門OID",
          "type": "string"
        }
      }
    },
    "Departments": {
      "description": "部門清單",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Department"
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "自定義錯誤碼\n定義來源參考:\nhttps://gitlab.kenda.com.tw/kenda/mcom/-/blob/${xxx}/errors/code.proto ` + "`" + `Code` + "`" + ` enum.\nin which ` + "`" + `${xxx}` + "`" + ` inside the URL reference is the specified branch name of the corresponding features.\n",
          "type": "integer",
          "x-omitempty": false
        },
        "details": {
          "description": "補充訊息",
          "type": "string"
        }
      }
    },
    "Limitary": {
      "type": "object",
      "properties": {
        "max": {
          "type": "string",
          "format": "date-time",
          "x-order": 1,
          "readOnly": true
        },
        "min": {
          "type": "string",
          "format": "date-time",
          "x-order": 0,
          "readOnly": true
        }
      }
    },
    "LoginRequest": {
      "required": [
        "ID",
        "password",
        "loginType"
      ],
      "properties": {
        "ID": {
          "description": "使用者工號/帳號",
          "type": "string"
        },
        "loginType": {
          "$ref": "#/definitions/LoginType"
        },
        "password": {
          "description": "使用者密碼",
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "properties": {
        "authorizedDepartments": {
          "$ref": "#/definitions/Departments"
        },
        "roles": {
          "$ref": "#/definitions/Roles"
        },
        "token": {
          "description": "使用者令牌(token)",
          "type": "string"
        },
        "tokenExpiry": {
          "description": "token 過期時間",
          "type": "string",
          "format": "date-time",
          "example": "2017-07-21T17:32:28Z"
        }
      }
    },
    "LoginType": {
      "description": "登入類別:\n  * 0 - MES\n  * 1 - Windows\n",
      "type": "integer",
      "enum": [
        0,
        1
      ],
      "x-omitempty": false
    },
    "Principal": {
      "properties": {
        "ID": {
          "description": "使用者工號/帳號",
          "type": "string",
          "example": "userid"
        },
        "roles": {
          "$ref": "#/definitions/Roles"
        }
      }
    },
    "Role": {
      "description": "授權角色\n定義來源參考: https://gitlab.kenda.com.tw/kenda/mcom/-/blob/${xxx}/utils/roles/roles.proto ` + "`" + `Role` + "`" + ` type.\nin which ` + "`" + `${xxx}` + "`" + ` inside the URL reference is the specified branch name of the corresponding features.\n",
      "type": "integer",
      "x-omitempty": false
    },
    "Roles": {
      "description": "使用者角色",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Role"
      }
    }
  },
  "parameters": {
    "DepartmentOID": {
      "type": "string",
      "description": "部門OID",
      "name": "departmentOID",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "Default": {
      "description": "Unexpected error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "x-api-auth-key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": []
    }
  ],
  "tags": [
    {
      "description": "帳號管理相關",
      "name": "account"
    },
    {
      "name": "product"
    }
  ]
}`))
}
