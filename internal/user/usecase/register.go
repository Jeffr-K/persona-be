package usecase

import (
	"fmt"
	adapter "iam/internal/user/adapter/in"
	"iam/internal/user/entity"
	"iam/internal/user/port"
	account "iam/internal/user/repository/command"
	"iam/libs/database"
)

type IUserRegisterUseCase interface {
	RegisterMembership(dto *adapter.UserRegisterCommand) error
}

type UserRegisterUseCase struct{}

func (uc UserRegisterUseCase) RegisterMembership(command *adapter.UserRegisterCommand) error {
	aggregate := entity.User{}

	port := port.UserRegisterInBoundPort{
		Username: command.Username,
		Password: command.Password,
		Email:    command.Email,
		Address:  command.Address,
	}

	user, err := aggregate.Register(&port)
	if err != nil {
		fmt.Println("Error: UserRegisterUseCase.UserRegisterAsMember", err)
		return err
	}

	repository := account.NewCommandRepository(database.Client)
	err = repository.SaveTo(&user)
	if err != nil {
		fmt.Println("Error: UserCommandRepository.SaveTo", err)
	}

	return nil
}
