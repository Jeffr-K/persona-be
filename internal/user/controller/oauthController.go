package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type OAuthController struct{}

func (c OAuthController) Google(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}

func (c OAuthController) Facebook(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}

func (c OAuthController) Kakao(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}

func (c OAuthController) Naver(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}
