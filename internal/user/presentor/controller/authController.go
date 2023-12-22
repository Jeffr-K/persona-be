package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	inboundPort "persona/internal/user/application/port"
	"persona/internal/user/application/usecase"
	inbound "persona/internal/user/presentor/adapter/in"
	"persona/libs/utils"
	"strconv"
)

type AuthController struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthController(usecase *usecase.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: usecase}
}

func (c AuthController) Login(context echo.Context) error {
	user, err := utils.ExtractUserFromContext(context)
	if err != nil {
		return err
	}
	fmt.Println("user: >>>", user)

	command := inbound.UserLoginCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	response, err := c.authUseCase.Login(&command)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다")
	}

	return context.JSON(http.StatusOK, response)
}

func (c AuthController) Logout(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Path Parameter 요청 값이 잘못되었습니다.")
	}

	port := inboundPort.UserLogoutInBoundPort{ID: id}
	if err := c.authUseCase.Logout(&port); err != nil {
		return context.JSON(http.StatusBadRequest, "Logout Failed.")
	}

	return context.JSON(http.StatusOK, "Logout Succeed.")
}
