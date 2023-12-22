package presentor

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/presentor/controller"
	"persona/libs/security/token"
)

type AuthRoutes struct {
	authController  controller.AuthController
	tokenController controller.TokenController
}

func NewAuthRoutes(
	userController controller.AuthController,
	tokenController controller.TokenController,
) AuthRoutes {
	return AuthRoutes{
		authController:  userController,
		tokenController: tokenController,
	}
}

func (a AuthRoutes) InitializeRoutes(route *echo.Echo) {
	auth := route.Group("/auth")

	auth.POST("/login", token.BearerTokenParser(a.authController.Login))
	//auth.POST("/login", a.authController.Login)
	auth.DELETE("/logout", a.authController.Logout)

	auth.POST("/token/reissue", a.tokenController.ReIssue)
}
