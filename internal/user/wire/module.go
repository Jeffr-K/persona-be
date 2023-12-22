package wire

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/presentor/controller"
	"persona/internal/user/presentor/route"
)

func RegisterUserModule(server *echo.Echo) {
	userController := RegisterUserDependency()
	authController := RegisterAuthDependency()
	tokenController := RegisterTokenDependency()
	oauthController := controller.OAuthController{}

	userRouter := presentor.NewUserRoutes(userController)
	authRouter := presentor.NewAuthRoutes(authController, tokenController)
	oauthRouter := presentor.NewOAuthRoutes(oauthController)

	userRouter.InitializeRoutes(server)
	authRouter.InitializeRoutes(server)
	oauthRouter.InitializeRoutes(server)

}
