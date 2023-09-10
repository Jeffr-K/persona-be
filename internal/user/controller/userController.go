package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	adapter "iam/internal/user/adapter/in"
	"iam/internal/user/service"
	usecase "iam/internal/user/usecase"
	"net/http"
)

type UserController struct {
	userRegisterUseCase usecase.UserRegisterUseCase
	userDropdownUseCase usecase.UserDropdownUseCase
	userExternalService service.UserExternalService
}

// Search @Summary Get User
// @Success 200 {object} string
// @Router /user/search [get]
func (c UserController) Search(context echo.Context) error {
	query := adapter.UserSearchQuery{}

	if err := context.Bind(&query); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	user, err := c.userExternalService.FindOneBy(query.ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "찾는 유저가 없습니다.")
	}

	return context.JSON(http.StatusOK, user)
}

func (c UserController) List(context echo.Context) error {
	list, err := c.userExternalService.FindAll()
	if err != nil {
		fmt.Println("ERROR: Not found any user.")
		return context.JSON(http.StatusInternalServerError, list)
	}

	return context.JSON(http.StatusInternalServerError, list)
}

func (c UserController) Register(context echo.Context) error {
	command := adapter.UserRegisterCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	if result := c.userRegisterUseCase.RegisterMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusOK, "회원가입이 완료되었습니다.")
}

func (c UserController) Dropdown(context echo.Context) error {
	command := adapter.UserDropdownCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	if result := c.userDropdownUseCase.DropdownMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusOK, context.JSON(http.StatusOK, "회원탈퇴가 정상적으로 완료되었습니다."))
}

func (c UserController) Modify(context echo.Context) error {
	return context.JSON(http.StatusOK, context.JSON(http.StatusOK, "hello"))
}
