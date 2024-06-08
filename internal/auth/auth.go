package auth

import (
	"BIOTRACKERSERVICE/internal/config"
	"BIOTRACKERSERVICE/internal/usecases"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct {
	PrivateKey     string
	ExpirationTime time.Duration
}

func New(cfg config.Auth) *Auth {
	return &Auth{
		PrivateKey:     cfg.PrivateKey,
		ExpirationTime: cfg.ExpirationTime,
	}
}

func (a *Auth) UserToken(ctx context.Context, credentials usecases.CredentialsDTO) (string, error) {
	expirationTime := time.Now().Add(a.ExpirationTime)
	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) UserAuth(ctx context.Context, token string) error {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(a.PrivateKey), nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return ErrUnauthorized
	}
	if err != nil {
		return ErrInvalidToken
	}
	if !tkn.Valid {
		return ErrUnauthorized
	}
	return nil
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
