package usecase

import (
	adapter "persona/internal/user/adapter/in"
)

type UserUseCase struct{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (usecase UserUseCase) RegisterMembership(command *adapter.UserRegisterCommand) error {
	//aggregate := entity.User{}
	//
	//port := port.UserRegisterInBoundPort{
	//	Username: command.Username,
	//	Password: command.Password,
	//	Email:    command.Email,
	//}
	return nil

}

func (usecase UserUseCase) DropdownMembership(command *adapter.UserDropdownCommand) error {
	return nil
}
