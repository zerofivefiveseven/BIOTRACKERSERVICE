package main

import (
	"BIOTRACKERSERVICE/internal/auth"
	"BIOTRACKERSERVICE/internal/config"
	"BIOTRACKERSERVICE/internal/handlers"
	repo "BIOTRACKERSERVICE/internal/repository"
	"BIOTRACKERSERVICE/internal/usecases"
	"BIOTRACKERSERVICE/internal/usecases/authentication"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	pxgConfig, err := pgxpool.ParseConfig(cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, pxgConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	repository := repo.New(dbPool)
	dbManager := usecases.NewRepo(usecases.Deps{
		Repository: repository,
		TxBuilder:  dbPool,
	})

	authenticator := auth.New(cfg.Auth)
	authSystem := authentication.NewAuthenticationSystem(authentication.Deps{
		Authenticator: authenticator,
		Repo:          repository,
	})

	UC := handlers.Usecases{
		DbManager: dbManager,
		Auth:      authSystem,
	}

	controller := handlers.NewController(UC)
	server := controller.NewServer(cfg.HTTPServer)
	log.Printf("server is listening at %s", cfg.HTTPServer.Address)
	log.Fatal(server.ListenAndServe())
}
