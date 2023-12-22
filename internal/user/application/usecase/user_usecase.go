package usecase

import (
	"fmt"
	"persona/internal/user/application/port"
	user "persona/internal/user/domain/entity"
	"persona/internal/user/domain/service"
	commander "persona/internal/user/infrastructure/repository/command"
	queryier "persona/internal/user/infrastructure/repository/query"
	inbound "persona/internal/user/presentor/adapter/in"
	"persona/libs/database"
)

type UserUseCase struct{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (uc UserUseCase) RegisterMembership(command *inbound.UserRegisterCommand) error {
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

	userRepository := commander.NewCommandRepository(database.Client)
	if err = userRepository.SaveTo(schema); err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) DropdownMembership(command *inbound.UserDropdownCommand) error {
	userService := service.NewUserService()

	aggregate, err := userService.FindOneBy(command.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	userRepository := commander.NewCommandRepository(database.Client)
	if err = userRepository.DeleteTo(aggregate.ID); err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) FindUserByQuery(query *inbound.UserSearchQuery) (user.User, error) {
	userRepository := queryier.NewQueryRepository(database.Client)
	member, err := userRepository.GetBy(query.ID, "id")
	if err != nil {
		return user.User{}, err
	}

	return member, nil
}

func (uc UserUseCase) FindUsersByQuery() ([]user.User, error) {
	userRepository := queryier.NewQueryRepository(database.Client)
	users, err := userRepository.GetAllList()

	if err != nil {
		return []user.User{}, nil
	}

	return users, nil
}
