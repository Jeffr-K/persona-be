package service

import (
	"fmt"
	"persona/internal/user/entity"
	"persona/internal/user/repository/query"
	"persona/libs/database"
)

type UserExternalService struct{}

func (ex UserExternalService) FindOneBy(userId int) (domain.User, error) {
	repository := query.NewQueryRepository(database.Client)

	user, err := repository.GetTo(userId)
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return domain.User{}, err
	}

	return user, nil
}

func (ex UserExternalService) FindAll() ([]domain.User, error) {
	repository := query.NewQueryRepository(database.Client)

	list, err := repository.GetAllList()
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return []domain.User{}, err
	}

	return list, nil
}

func (ex UserExternalService) FindAllBy() []domain.User {
	return []domain.User{}
}
