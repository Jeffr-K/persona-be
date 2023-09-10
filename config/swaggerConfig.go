package config

import (
	"github.com/swaggo/swag"
)

const DocumentInfo = `{
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

func MakeSwaggerSpec() swag.Spec {
	return swag.Spec{
		Version:          "1.0",
		Schemes:          []string{},
		Title:            "Sunnyside API Specification",
		Description:      "The Sunnyside Backend Application API Specification",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  DocumentInfo,
		LeftDelim:        "{{",
		RightDelim:       "}}",
		//Host:             "localhost:8080",
		//BasePath:         "/api/docs/*",
	}
}
