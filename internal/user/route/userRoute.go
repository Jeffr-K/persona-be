package presentor

import (
	"github.com/labstack/echo/v4"
	"iam/internal/user/controller"
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
	user.DELETE("/dropdown", u.userController.Dropdown)
	user.GET("/search/:id", u.userController.Search)
	user.GET("/search/list", u.userController.List)
	user.PUT("/profile", u.userController.Modify)
}
