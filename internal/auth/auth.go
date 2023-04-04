package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct {
	privateKey []byte
}

func NewAuth(privateKey []byte) *Auth {
	return &Auth{
		privateKey: privateKey,
	}
}

func (a Auth) SignIn(_ context.Context, username, password string) (string, error) {
	if username == "" || password == "" {
		return "", ErrInvalidCredentials
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	return token.SignedString(a.privateKey)
}

func (a Auth) Validate(_ context.Context, jws string) error {
	token, err := jwt.Parse(jws, func(token *jwt.Token) (interface{}, error) {
		return a.privateKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return ErrInvalidToken
	}
	if !token.Valid {
		return ErrInvalidToken
	}
	return nil
}
