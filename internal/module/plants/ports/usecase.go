package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/models"
)

type Usecase interface {
	GetAllPlant(ctx context.Context) (resp []models.BasicPlant, err *errors.BaseError)
	GetPlantByID(ctx context.Context, id uint) (resp *models.Plant, err *errors.BaseError)
	GetActivePlant(ctx context.Context) (resp *models.Plant, err *errors.BaseError)

	InsertPlant(ctx context.Context, createPlant models.CreatePlant) (resp *models.Plant, err *errors.BaseError)
	UpdatePlantStatus(ctx context.Context, data models.UpdatePlantStatus) *errors.BaseError
	UpdatePlant(ctx context.Context, data models.UpdatePlant) *errors.BaseError
}
