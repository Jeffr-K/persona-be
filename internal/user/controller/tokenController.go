package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type TokenController struct{}

func (t TokenController) ReIssue(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}
