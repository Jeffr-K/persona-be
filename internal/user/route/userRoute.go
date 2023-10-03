package presentor

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/controller"
)

type UserRoutes struct {
	userController controller.UserController
}

func NewUserRoutes(controller controller.UserController) UserRoutes {
	return UserRoutes{
		userController: controller,
	}
}

func (u UserRoutes) InitializeRoutes(route *echo.Echo) {
	user := route.Group("/user")

	user.POST("/registration", u.userController.Register)
	user.DELETE("/dropdown/:userId", u.userController.Dropdown)
	user.GET("/search/:userId", u.userController.GetUserById)
	user.GET("/search/list", u.userController.GetUsers)
}
