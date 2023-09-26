package wire

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/controller"
	"persona/internal/user/route"
)

func RegisterUserModule(server *echo.Echo) {
	userController := RegisterUserDependency()
	authController := controller.AuthController{}
	oauthController := controller.OAuthController{}
	tokenController := controller.TokenController{}

	userRouter := presentor.NewUserRoutes(userController)
	authRouter := presentor.NewAuthRoutes(authController)
	oauthRouter := presentor.NewOAuthRoutes(oauthController)
	tokenRouter := presentor.NewTokenRoutes(tokenController)

	userRouter.InitializeRoutes(server)
	authRouter.InitializeRoutes(server)
	oauthRouter.InitializeRoutes(server)
	tokenRouter.InitializeRoutes(server)
}
