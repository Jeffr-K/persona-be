package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"persona/internal/user/controller"
	presentor "persona/internal/user/route"
	"persona/libs/database"
	"persona/libs/database/redis"
)

type Application struct{}

func (app Application) BootStrap(echo *echo.Echo) {
	app.registerRouter(echo)
	app.registerMiddleware(echo)
	database.InitializeDatabase()
	redis.InitializeRedis()
	//queue.InitializeKafka()
}

func (app Application) registerRouter(server *echo.Echo) {
	userController := controller.UserController{}
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

func (app Application) registerMiddleware(server *echo.Echo) {
	server.Use(middleware.CORSWithConfig(InitializeCorsConfig()))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
}
