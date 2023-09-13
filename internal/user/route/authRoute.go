package presentor

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/controller"
)

type AuthRoutes struct {
	authController controller.AuthController
}

func NewAuthRoutes(controller controller.AuthController) AuthRoutes {
	return AuthRoutes{
		authController: controller,
	}
}

func (a AuthRoutes) InitializeRoutes(route *echo.Echo) {
	auth := route.Group("/auth")

	auth.POST("/", a.authController.Login)
	auth.DELETE("/", a.authController.Logout)
}
