package usecases

import (
	repo "BIOTRACKERSERVICE/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// var _ struct{} = (*repo.Queries{})(nil)
type (
	DbManager interface {
		AddPlot(ctx context.Context, arg repo.AddPlotParams) error
		AddUserCredentials(ctx context.Context, arg repo.AddUserCredentialsParams) (int64, error)
		AddUserInfo(ctx context.Context, arg repo.AddUserInfoParams) error
		CheckUserExists(ctx context.Context, email string) (bool, error)
		GetPlotsByIds(ctx context.Context, plotIds []int32) ([]repo.Plot, error)
		GetUserCredentials(ctx context.Context, email string) (string, error)
		GetUserPlotsInfo(ctx context.Context, userID int32) ([]repo.GetUserPlotsInfoRow, error)
		WithTx(tx pgx.Tx) *repo.Queries
	}
	txBuilder interface {
		Begin(ctx context.Context) (pgx.Tx, error)
	}
	db interface {
		repo.DBTX
		txBuilder
	}
)
type Deps struct {
	Repository DbManager
	TxBuilder  db
}
type Repo struct {
	Deps
}

func NewRepo(deps Deps) *Repo {
	return &Repo{
		Deps: deps,
	}
}

func (r *Repo) RegisterUser(ctx context.Context, user UserDTO) error {
	tx, err := r.TxBuilder.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := r.Repository.WithTx(tx)
	exists, err := qtx.CheckUserExists(ctx, user.Email)

	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExists
	}
	id, err := qtx.AddUserCredentials(ctx, repo.AddUserCredentialsParams{
		UserEmail:    user.Email,
		UserPassword: user.Password,
	})
	if err != nil {
		return err
	}
	err = qtx.AddUserInfo(ctx, repo.AddUserInfoParams{
		ID:        int32(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		//TODO: not timestamp
		Dateofbirhday: pgtype.Timestamp{},
	})
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
