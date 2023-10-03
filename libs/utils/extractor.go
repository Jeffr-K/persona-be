package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"persona/internal/user/entity"
)

func ExtractUserFromContext(context echo.Context) (entity.User, error) {
	userContext := context.Get("user")

	user, ok := userContext.(entity.User)
	if !ok {
		return entity.User{}, fmt.Errorf("user not found in context")
	}

	return user, nil
}
