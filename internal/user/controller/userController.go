package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	in2 "persona/internal/user/adapter/in"
	"persona/internal/user/usecase"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c UserController) Register(context echo.Context) error {
	command := in2.UserRegisterCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	if result := c.userUseCase.RegisterMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusOK, "회원가입이 완료되었습니다.")
}

func (c UserController) Dropdown(context echo.Context) error {
	command := in2.UserDropdownCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	if result := c.userUseCase.DropdownMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusOK, context.JSON(http.StatusOK, "회원탈퇴가 정상적으로 완료되었습니다."))
}
