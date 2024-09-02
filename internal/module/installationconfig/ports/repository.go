package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/entities"
)

//go:generate mockery --name Repository
type Repository interface {
	UpdateContainerConfig(ctx context.Context, data *entities.ContainerConfig) *errors.BaseError
	FindContainerConfigByID(ctx context.Context, ContainerID string) (*entities.ContainerConfig, *errors.BaseError)
	GetAllContainerConfig(ctx context.Context) ([]entities.ContainerConfig, *errors.BaseError)

	UpdateInstallationConfig(ctx context.Context, data *entities.InstallationConfig) *errors.BaseError
	GetInstallationConfig(ctx context.Context, ID uint) (*entities.InstallationConfig, *errors.BaseError)
}
