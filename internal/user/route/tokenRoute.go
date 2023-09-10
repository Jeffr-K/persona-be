package presentor

import (
	"github.com/labstack/echo/v4"
	"iam/internal/user/controller"
)

type TokenRoute struct {
	tokenController controller.TokenController
}

func NewTokenRoutes(controller controller.TokenController) TokenRoute {
	return TokenRoute{
		tokenController: controller,
	}
}

func (tr TokenRoute) InitializeRoutes(route *echo.Echo) {
	user := route.Group("/token")

	user.POST("/", tr.tokenController.ReIssue)
}
