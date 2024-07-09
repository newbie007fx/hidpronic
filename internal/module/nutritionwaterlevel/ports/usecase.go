package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	"time"
)

type Usecase interface {
	InsertNutritionWaterLevelTemp(ctx context.Context, req models.CreateNutritionWaterLevel) *errors.BaseError

	GetActivePlantNutritionWaterLevel(ctx context.Context) (resp []models.NutritionWaterLevel, err *errors.BaseError)
	InsertNutritionWaterLevelFromTempRange(ctx context.Context, plantID uint, startDate, endDate time.Time) (resp *models.NutritionWaterLevel, err *errors.BaseError)
}
