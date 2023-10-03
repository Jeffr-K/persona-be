package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	inbound "persona/internal/user/adapter/in"
	"persona/internal/user/entity"
	"persona/internal/user/usecase"
	"strconv"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c UserController) Register(context echo.Context) error {
	command := inbound.UserRegisterCommand{}

	if err := context.Bind(&command); err != nil {
		return context.JSON(http.StatusBadRequest, "입력값이 잘못되었습니다.")
	}

	if result := c.userUseCase.RegisterMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusOK, "회원가입이 완료되었습니다.")
}

func (c UserController) Dropdown(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Path Parameter 요청 값이 잘못되었습니다.")
	}

	command := inbound.UserDropdownCommand{ID: id}

	if result := c.userUseCase.DropdownMembership(&command); result != nil {
		return context.JSON(http.StatusInternalServerError, "서버 내부에서 에러가 발생하였습니다.")
	}

	return context.JSON(http.StatusCreated, context.JSON(http.StatusOK, "회원탈퇴가 정상적으로 완료되었습니다."))
}

func (c UserController) GetUserById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("userId"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Path Parameter 요청 값이 잘못되었습니다.")
	}

	query := inbound.UserSearchQuery{ID: id}

	user, err := c.userUseCase.FindUserByQuery(&query)
	if err != nil {
		return context.JSON(http.StatusNotFound, "찾는 유저가 없습니다.")
	}

	return context.JSON(http.StatusOK, user)
}

func (c UserController) GetUsers(context echo.Context) error {
	users, err := c.userUseCase.FindUsersByQuery()

	if err != nil {
		return context.JSON(http.StatusOK, []entity.User{})
	}

	return context.JSON(http.StatusOK, users)
}
