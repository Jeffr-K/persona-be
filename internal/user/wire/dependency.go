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

func RegisterAuthDependency() controller.AuthController {
	authUseCase := usecase.NewAuthUseCase()
	authController := controller.NewAuthController(authUseCase)

	return *authController
}

func RegisterTokenDependency() controller.TokenController {
	tokenUseCase := usecase.NewTokenUseCase()
	tokenController := controller.NewTokenController(tokenUseCase)

	return *tokenController
}

//func RegisterAuthDependency() controller.AuthController {}
//func RegisterOAuthDependency() controller.OAuthController {}
//func RegisterNamecardDependency() controller.NamecardController {}
//func RegisterPersonalizationDependency() controller.PersonalizationController {}
//func RegisterPhoneDependency() controller.PhoneController {}
//func RegisterReferrerDependency() controller.ReferrerController {}
//func RegisterFollowDependency() controller.FollowController {}
//func RegisterProfileDependency() controller.ProfileController {}
