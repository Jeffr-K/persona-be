package usecase

import (
	in2 "persona/internal/user/adapter/in"
	user "persona/internal/user/entity"
	"persona/internal/user/port"
	domain "persona/internal/user/repository/command"
	"persona/libs/database"
)

type UserUseCase struct{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (usecase UserUseCase) RegisterMembership(command *in2.UserRegisterCommand) error {
	aggregate := user.New(command.Username, command.Password, command.Email)
	port := port.UserRegisterInBoundPort{
		Username: command.Username,
		Password: command.Password,
		Email:    command.Email,
	}

	schema, err := aggregate.Register(&port)
	if err != nil {
		return err
	}

	userRepository := domain.NewCommandRepository(database.Client)
	if err = userRepository.SaveTo(&schema); err != nil {
		return err
	}

	return nil
}

func (usecase UserUseCase) DropdownMembership(command *in2.UserDropdownCommand) error {
	return nil
}
