package wire

import (
	usecase2 "persona/internal/user/application/usecase"
	"persona/internal/user/presentor/controller"
)

func RegisterUserDependency() controller.UserController {
	userUseCase := usecase2.NewUserUseCase()
	userController := controller.NewUserController(userUseCase)

	return *userController
}

func RegisterAuthDependency() controller.AuthController {
	authUseCase := usecase2.NewAuthUseCase()
	authController := controller.NewAuthController(authUseCase)

	return *authController
}

func RegisterTokenDependency() controller.TokenController {
	tokenUseCase := usecase2.NewTokenUseCase()
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
