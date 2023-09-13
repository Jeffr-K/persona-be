package presentor

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/controller"
)

type OAuthRoutes struct {
	oauthController controller.OAuthController
}

func NewOAuthRoutes(controller controller.OAuthController) OAuthRoutes {
	return OAuthRoutes{
		oauthController: controller,
	}
}

func (o OAuthRoutes) InitializeRoutes(route *echo.Echo) {
	oauth := route.Group("/oauth/google")

	oauth.POST("/", o.oauthController.Google)
	oauth.POST("/", o.oauthController.Facebook)
	oauth.POST("/", o.oauthController.Naver)
	oauth.POST("/", o.oauthController.Kakao)

}
