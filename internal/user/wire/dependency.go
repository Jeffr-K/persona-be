package wire

import (
	"persona/internal/user/controller"
	"persona/internal/user/usecase"
)

func RegisterUserDependency() controller.UserController {
	userUseCase := usecase.NewUserUseCase()
	userController := controller.NewUserController(userUseCase)

	return *userController
}

//func RegisterAuthDependency() controller.AuthController {}
//func RegisterOAuthDependency() controller.OAuthController {}
//func RegisterTokenDependency() controller.TokenController {}
//func RegisterNamecardDependency() controller.NamecardController {}
//func RegisterPersonalizationDependency() controller.PersonalizationController {}
//func RegisterPhoneDependency() controller.PhoneController {}
//func RegisterReferrerDependency() controller.ReferrerController {}
//func RegisterFollowDependency() controller.FollowController {}
//func RegisterProfileDependency() controller.ProfileController {}
