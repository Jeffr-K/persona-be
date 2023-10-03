package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	user "persona/internal/user/wire"
	"persona/libs/database"
	"persona/libs/database/redis"
)

type Application struct{}

func (app Application) BootStrap(echo *echo.Echo) {
	app.registerModules(echo)
	app.registerMiddleware(echo)
	database.InitializeDatabase()
	redis.InitializeRedis()
	//kafka.InitializeKafka()
}

func (app Application) registerModules(server *echo.Echo) {
	user.RegisterUserModule(server)
}

func (app Application) registerMiddleware(server *echo.Echo) {
	server.Use(middleware.CORSWithConfig(InitializeCorsConfig()))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
}
