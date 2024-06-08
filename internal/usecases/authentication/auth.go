package authentication

import (
	"BIOTRACKERSERVICE/internal/auth"
	"BIOTRACKERSERVICE/internal/usecases"
	"context"
	"database/sql"
	"errors"
)

type (
	Authenticator interface {
		UserToken(ctx context.Context, credentials usecases.CredentialsDTO) (string, error)
		UserAuth(ctx context.Context, token string) error
	}
	Repository interface {
		UserCredentials(ctx context.Context, username string) (string, error)
	}
)

type Deps struct {
	Authenticator Authenticator
	Repo          Repository
}
type AuthSystem struct {
	Deps
}

func NewAuthenticationSystem(deps Deps) *AuthSystem {
	return &AuthSystem{
		Deps: deps,
	}
}

func (s *AuthSystem) UserToken(ctx context.Context, credentials usecases.CredentialsDTO) (string, error) {
	expectedPassword, err := s.Repo.UserCredentials(ctx, credentials.Username)
	if err != nil {
		return "", auth.ErrUnauthorized
	}
	if expectedPassword != credentials.Password {
		return "", auth.ErrUnauthorized
	}

	token, err := s.Deps.Authenticator.UserToken(ctx, credentials)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *AuthSystem) UserAuth(ctx context.Context, token string) error {
	return s.Deps.Authenticator.UserAuth(ctx, token)
}
func (s *AuthSystem) UserCredentials(ctx context.Context, username string) (string, error) {
	password, err := s.Repo.UserCredentials(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", auth.ErrUnauthorized
	}
	if err != nil {
		return "", err
	}
	return password, nil
}
