package entity

import (
	port "iam/internal/user/port"
	"iam/libs/security"
	"log"
	"time"
)

type User struct {
	ID        uint
	Version   uint
	Username  string
	Password  string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (root User) Register(command *port.UserRegisterInBoundPort) (User, error) {
	hashedPassword, err := security.NewBCrypt().HashingPassword(command.Password)
	if err != nil {
		log.Fatalln("Password Hashing error: ", err)
		return User{}, err
	}

	user := User{
		Username: command.Username,
		Password: hashedPassword,
		Email:    command.Email,
		Address:  command.Address,
	}

	return user, nil
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
