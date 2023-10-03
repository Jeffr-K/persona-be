package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"persona/config"
	_ "persona/docs"
)

// @title Sunnyside
// @version 0.1
func main() {
	server := echo.New()

	server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, sunnyside")
	})

	config.Application{}.BootStrap(server)

	//q := kafka.Queue{}
	//
	//if err := q.Produce("찌혜얌, 닥꼬기 맛이께 머겅 내사룽. "); err != nil {
	//	fmt.Println("에러얌얌")
	//}
	server.GET("/api/docs/*", echoSwagger.WrapHandler)

	server.Logger.Fatal(server.Start(":8080"))
}
