// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/api/createEvent": {
            "post": {
                "description": "Endpoint used to create an entry of the event in db.",
                "tags": [
                    "Events"
                ],
                "summary": "Create Event",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/deleteEvent": {
            "delete": {
                "description": "Endpoint used to delete an event from db based on id.",
                "tags": [
                    "Events"
                ],
                "summary": "Delete Event",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/getAttendees": {
            "get": {
                "description": "Endpoint used to get the list of all the attendees in an event.",
                "tags": [
                    "Events"
                ],
                "summary": "Get Attendees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/getEvent": {
            "get": {
                "description": "Endpoint used to get an events from db based on id.",
                "tags": [
                    "Events"
                ],
                "summary": "Get Event",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/getEvents": {
            "get": {
                "description": "Endpoint used to get all events from db.",
                "tags": [
                    "Events"
                ],
                "summary": "Get Events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/getUser": {
            "get": {
                "description": "Endpoint used to get an user details from the db.",
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/joinevent": {
            "get": {
                "description": "Endpoint used to add a user to an Event.",
                "tags": [
                    "Events"
                ],
                "summary": "Join an Event",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Endpoint used to login an user by validating credentials and generating JWT cookie for session.",
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/logout": {
            "get": {
                "description": "Endpoint used to logout an user by deleting the JWT cookie used for mainting session.",
                "tags": [
                    "User"
                ],
                "summary": "Logout User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "Endpoint used to create an user in db.",
                "tags": [
                    "User"
                ],
                "summary": "Register User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/searchEvent": {
            "get": {
                "description": "Endpoint used to get all events based on the filter from db.",
                "tags": [
                    "Events"
                ],
                "summary": "Get Filtered Events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "attendes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "interests": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
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
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:9010",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "SocialBytes API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
