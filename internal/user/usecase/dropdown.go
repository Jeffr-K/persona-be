package usecase

import (
	"fmt"
	adapter "iam/internal/user/adapter/in"
	"iam/internal/user/entity"
	"iam/internal/user/port"
	account "iam/internal/user/repository/command"
	"iam/libs/database"
)

type IUserDropdownUseCase interface {
}

type UserDropdownUseCase struct {
}

func (ud UserDropdownUseCase) DropdownMembership(command *adapter.UserDropdownCommand) error {
	aggregate := entity.User{}

	port := port.UserDropdownInBoundPort{
		ID: command.ID,
	}

	_, err := aggregate.Dropdown(&port)
	if err != nil {
		fmt.Println("ERROR: Failed to dropdown a user", err)
		return err
	}
	fmt.Println("0", port.ID)
	repository := account.NewCommandRepository(database.Client)
	err = repository.DeleteTo(port.ID)

	return nil
}
