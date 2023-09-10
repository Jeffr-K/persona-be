package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct{}

func (c AuthController) Login(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, auth Controller")
}

func (c AuthController) Logout(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, auth Controller")
}
