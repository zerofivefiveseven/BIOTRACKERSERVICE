package handlers

import (
	repo "BIOTRACKERSERVICE/internal/repository"
	"BIOTRACKERSERVICE/internal/usecases"
	"context"
)

// var _ struct{} = (*usecases.Repo)(nil)
// var _ struct{} = (*authentication.AuthSystem)(nil)

type (
	DbManager interface {
		RegisterUser(ctx context.Context, user usecases.UserDTO) error
		AddPlot(ctx context.Context, plotInfo usecases.AddPlotDTO) error
		GetUserPlotsInfo(ctx context.Context, userID int32) ([]repo.GetUserPlotsInfoRow, error)
		GetPlotsByIds(ctx context.Context, ids []int32) ([]usecases.PlotDTO, error)
	}
	Auth interface {
		UserToken(ctx context.Context, credentials usecases.CredentialsDTO) (string, error)
		UserAuth(ctx context.Context, token string) error
		UserCredentials(ctx context.Context, username string) (string, error)
	}
)

type Usecases struct {
	DbManager
	Auth
}

// Controller - is controller/delivery layer
type Controller struct {
	Usecases
}

// NewController - returns Controller
func NewController(us Usecases) *Controller {
	return &Controller{
		Usecases: us,
	}
}
