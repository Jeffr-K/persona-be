package usecase

type OAuthUseCase struct{}

func NewOAuthUseCase() *OAuthUseCase {
	return &OAuthUseCase{}
}

func (ou OAuthUseCase) LoginWithGoogle() error {
	return nil
}

func (ou OAuthUseCase) LoginWithFacebook() error {
	return nil
}

func (ou OAuthUseCase) LoginWithGithub() error {
	return nil
}

func (ou OAuthUseCase) LoginWithKakao() error {
	return nil
}

func (ou OAuthUseCase) LoginWithNaver() error {
	return nil
}
