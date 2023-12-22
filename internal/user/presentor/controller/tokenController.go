package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"persona/internal/user/application/usecase"
)

type TokenController struct {
	tokenUseCase *usecase.TokenUseCase
}

func NewTokenController(usecase *usecase.TokenUseCase) *TokenController {
	return &TokenController{tokenUseCase: usecase}
}

func (t TokenController) ReIssue(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, user Controller")
}
