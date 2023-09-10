package service

import (
	"fmt"
	"iam/internal/user/entity"
	"iam/internal/user/repository/query"
	"iam/libs/database"
)

type UserExternalService struct{}

func (ex UserExternalService) FindOneBy(userId int) (entity.User, error) {
	repository := query.NewQueryRepository(database.Client)

	user, err := repository.GetTo(userId)
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return entity.User{}, err
	}

	return user, nil
}

func (ex UserExternalService) FindAll() ([]entity.User, error) {
	repository := query.NewQueryRepository(database.Client)

	list, err := repository.GetAllList()
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return []entity.User{}, err
	}

	return list, nil
}

func (ex UserExternalService) FindAllBy() []entity.User {
	return []entity.User{}
}
