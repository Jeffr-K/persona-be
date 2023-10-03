package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type IToken interface {
	GenerateJwtAccessToken(email string) (string, error)
	GenerateJwtRefreshToken(email string) (string, error)
	IsValidToken(requestToken string) bool
	DecodeJwtToken(tokenString string) (Claims, error)
}

// Token https://developer.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
// Auth0 [auth0](https://auth0.com/docs/secure/tokens/id-tokens/validate-id-tokens)
type Token struct{}

type Claims struct {
	Email          string
	StandardClaims jwt.StandardClaims
}

func NewToken() IToken {
	return &Token{}
}

func (t Token) GenerateJwtAccessToken(email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t Token) GenerateJwtRefreshToken(email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t Token) IsValidToken(requestToken string) bool {
	_, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if err != nil {
		return false
	}

	return true
}

func (t Token) DecodeJwtToken(userToken string) (Claims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return Claims{}, errors.New("invalid user email")
	}

	return Claims{Email: email}, nil
}

func (t Token) VerifyJwtTokenValidation() {}
