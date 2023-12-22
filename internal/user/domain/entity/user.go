package entity

import (
	"github.com/google/uuid"
	"log"
	port2 "persona/internal/user/application/port"
	"persona/libs/security"
	"time"
)

type UserSpecification interface {
	Register(command *port2.UserRegisterInBoundPort) error
}

type User struct {
	ID        int
	UUID      uuid.UUID
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(username, password, email string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (root User) Register(command *port2.UserRegisterInBoundPort) (*User, error) {
	password, err := encodeHashPassword(command.Password)
	if err != nil {
		return &User{}, err
	}

	user := &User{
		Username:  command.Username,
		Password:  password,
		Email:     command.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}

func encodeHashPassword(password string) (string, error) {
	hashedPassword, err := security.NewBCrypt().HashingPassword(password)
	if err != nil {
		log.Fatalln("Password Hashing error: ", err)
		return "", err
	}
	return hashedPassword, nil
}

func (root User) Dropdown(command *port2.UserDropdownInBoundPort) (User, error) {
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
