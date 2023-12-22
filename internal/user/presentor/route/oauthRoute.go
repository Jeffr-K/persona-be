package presentor

import (
	"github.com/labstack/echo/v4"
	"persona/internal/user/presentor/controller"
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
	oauth := route.Group("/oauth")

	oauth.POST("/google/registration", o.oauthController.Google)
	oauth.POST("/facebook/registration", o.oauthController.Facebook)
	oauth.POST("/kakao/registration", o.oauthController.Naver)
	oauth.POST("/naver/registration", o.oauthController.Kakao)

}
