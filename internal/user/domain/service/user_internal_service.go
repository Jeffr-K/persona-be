package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"persona/internal/user/domain/entity"
	"persona/internal/user/infrastructure/repository/query"
	"persona/libs/database"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (ex userService) FindOneBy(userId int) (entity.User, error) {
	repository := query.NewQueryRepository(database.Client)

	user, err := repository.GetTo(userId)
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return entity.User{}, err
	}

	return user, nil
}

func (ex userService) FindOneByEmail(email string) (entity.User, error) {
	repository := query.NewQueryRepository(database.Client)
	user, err := repository.GetByEmail(email)
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return entity.User{}, err
	}

	return user, nil
}

func (ex userService) FindAll() ([]entity.User, error) {
	repository := query.NewQueryRepository(database.Client)

	list, err := repository.GetAllList()
	if err != nil {
		fmt.Println("ERROR: Not found a user")
		return []entity.User{}, err
	}

	return list, nil
}

func (ex userService) FindAllBy() []entity.User {
	return []entity.User{}
}

func Match(origin, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte(input))
	return err == nil
}
