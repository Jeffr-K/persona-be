package entity

import (
	"log"
	"persona/internal/user/port"
	comm "persona/internal/user/repository/command"
	"persona/libs/database"
	"persona/libs/security"
)

type User struct {
	Username string
	Password string
	Email    string
}

func (root User) Register(command *port.UserRegisterInBoundPort) error {
	hashedPassword, err := security.NewBCrypt().HashingPassword(command.Password)
	if err != nil {
		log.Fatalln("Password Hashing error: ", err)
		return err
	}

	user := User{
		Username: command.Username,
		Password: hashedPassword,
		Email:    command.Email,
	}

	repo := comm.NewCommandRepository(database.Client)
	repo.SaveTo(&user)

	return nil
}

func (root User) Dropdown(command *port.UserDropdownInBoundPort) (User, error) {
	// TODO: 딱히 무언가 할게 없다면
	// TODO: 도메인 이벤트를 생성해서 전파할 준비

	return User{}, nil
}

func (root User) FindUserAccount() error {
	return nil
}

func (root User) FindUserAccountPassword() error {
	return nil
}

func (root User) Login() error {
	return nil
}

func (root User) Logout() error {
	return nil
}
