package usecase

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var googleOAuthConfig = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

type IGoogleOAuth2UseCase interface {
	GetGoogleOAuthToken()
	GetUserProfile()
	GetUserProfileFromRedirectionUrl()
}

type GoogleOAuth2UseCase struct {
}

func NewGoogleOAuth2UseCase() IGoogleOAuth2UseCase {
	return &GoogleOAuth2UseCase{}
}

func (uc GoogleOAuth2UseCase) GetGoogleOAuthToken() {
	const url = "https://oauth2.googleapis.com/token"

}

func (uc GoogleOAuth2UseCase) GetUserProfile() {}

func (uc GoogleOAuth2UseCase) GetUserProfileFromRedirectionUrl() {}
