package usecase

import (
	"fmt"
	inboundPort "persona/internal/user/application/port"
	inbound "persona/internal/user/presentor/adapter/in"
	"persona/internal/user/presentor/adapter/out"

	"persona/internal/user/domain/service"
	"persona/libs/database/redis"
	"persona/libs/security/token"
	"strconv"
)

type AuthUseCase struct {
}

func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

func (au AuthUseCase) Login(command *inbound.UserLoginCommand) (interface{}, error) {
	userService := service.NewUserService()

	user, err := userService.FindOneByEmail(command.Email)
	if err != nil {
		return out.BusinessResponse{}, fmt.Errorf("UserNotFound %v", err)
	}

	userPassword := user.Password
	loginPassword := command.Password

	if err := service.Match(userPassword, loginPassword); err != true {
		return out.BusinessResponse{}, fmt.Errorf("user password is wrong. %v", err)
	}

	tokenFactory := token.NewToken()

	accessToken, err := tokenFactory.GenerateJwtAccessToken(user.Email)
	if err != nil {
		return out.BusinessResponse{}, fmt.Errorf("cannot make user access token. %v", err)
	}

	refreshToken, err := tokenFactory.GenerateJwtRefreshToken(user.Email)
	if err != nil {
		return out.BusinessResponse{}, fmt.Errorf("cannot make user refresh token. %v", err)
	}

	tokens := make(map[string]interface{})
	tokens["accessToken"] = accessToken
	tokens["refreshToken"] = refreshToken

	redisRepository := redis.NewRedisRepository()
	if err = redisRepository.Insert(user.Email, tokens); err != nil {
		fmt.Println("cannot insert into redis.")
		return out.BusinessResponse{}, fmt.Errorf("cannot make user refresh token. %v", err)
	}

	return tokens, nil
}

func (au AuthUseCase) Logout(port *inboundPort.UserLogoutInBoundPort) error {
	userService := service.NewUserService()

	user, err := userService.FindOneBy(port.ID)
	if err != nil {
		return err
	}

	redisRepository := redis.NewRedisRepository()
	if err = redisRepository.Delete(strconv.Itoa(user.ID)); err != nil {
		return err
	}

	return nil
}
