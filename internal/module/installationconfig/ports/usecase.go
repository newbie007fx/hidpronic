package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/models"
)

type Usecase interface {
	CalculateContainerVolume(ctx context.Context, ID string, distance float32) (float32, *errors.BaseError)

	GetInstallationConfig(ctx context.Context) (*models.InstallationConfigResponse, *errors.BaseError)
	UpdateInstallationConfig(ctx context.Context, data models.UpdateInstallationConfig) *errors.BaseError
}
