package token

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/domain/service"
	"persona/libs/database/redis"
)

func BearerTokenParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header.Get("Authorization")

		if len(header) < 7 || header[:7] != "Bearer " {
			return echo.ErrUnauthorized
		}

		token := header[7:]

		redisRepository := redis.NewRedisRepository()
		data, err := redisRepository.Get("wjdrlrkdl3@gmail.com")
		if err != nil {
			return err
		}

		storedAccessToken := data.AccessToken
		//refreshToken := data.RefreshToken

		if storedAccessToken != token {
			return echo.ErrUnauthorized
		}

		tokenFactory := NewToken()
		claim, err := tokenFactory.DecodeJwtToken(token)
		if err != nil {
			return err
		}

		email := claim.Email

		userService := service.NewUserService()
		user, err := userService.FindOneByEmail(email)
		if err != nil {
			return err
		}

		context.Set("user", user)

		return next(context)
	}
}
