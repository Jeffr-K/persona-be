package token

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JwtMiddlewareConfig = middleware.JWTConfig{
	Skipper:                 nil,
	BeforeFunc:              nil,
	SuccessHandler:          nil,
	ErrorHandler:            nil,
	ErrorHandlerWithContext: nil,
	ContinueOnIgnoredError:  false,
	SigningKey:              nil,
	SigningKeys:             nil,
	SigningMethod:           "",
	ContextKey:              "",
	Claims:                  nil,
	TokenLookup:             "",
	TokenLookupFuncs:        nil,
	AuthScheme:              "",
	KeyFunc:                 nil,
	ParseTokenFunc:          nil,
}

//func JwtParseErrorHandler(context echo.Context, err error) middleware.JWTErrorHandler {
//	return func(c echo.Context, err error) error {
//		return err
//	}
//}

func JwtTokenParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		//user := context.Get("user").(*jwt.Token)
		//claims := user.Claims.(jwt.MapClaims)
		//isAdmin := claims["admin"].(bool)
		//
		//if isAdmin == false {
		//	return echo.ErrUnauthorized
		//}
		fmt.Println("z=================================")

		return next(context)
	}
}

func BearerTokenParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header.Get("Authorization")

		if len(header) < 7 || header[:7] != "Bearer " {
			return echo.ErrUnauthorized
		}

		token := header[7:]

		// Redis 에 저장되어 있는 애랑 매치해서 맞으면 next 아니면 Err
		fmt.Println("token: ", token)

		return next(context)
	}
}
