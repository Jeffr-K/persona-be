package config

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitializeCorsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:                             []string{"*"},
		AllowMethods:                             []string{http.MethodHead, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodPatch},
		AllowCredentials:                         false,
		UnsafeWildcardOriginWithAllowCredentials: false,
	}
}
