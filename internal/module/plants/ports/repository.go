package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/entities"
)

type Repository interface {
	GetAllPlant(ctx context.Context) ([]entities.Plant, *errors.BaseError)
	GetPlantByID(ctx context.Context, id uint) (*entities.Plant, *errors.BaseError)
	GetActivePlant(ctx context.Context) (*entities.Plant, *errors.BaseError)

	InsertPlant(ctx context.Context, data *entities.Plant) *errors.BaseError
	UpdatePlant(ctx context.Context, data *entities.Plant) *errors.BaseError
	DeletePlant(ctx context.Context, id uint) *errors.BaseError
}
