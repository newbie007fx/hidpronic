package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/entities"
	"time"
)

type Repository interface {
	GetNutritionWaterLevelByPlantIDWithLimit(ctx context.Context, plantID uint, limit int) ([]entities.NutritionWaterLevel, *errors.BaseError)
	InsertNutritionWaterLevel(ctx context.Context, data *entities.NutritionWaterLevel) *errors.BaseError

	GetNutritionWaterLevelTempByPlantIDWithByRange(ctx context.Context, plantID uint, startDate, endDate time.Time) ([]entities.NutritionWaterLevel, *errors.BaseError)
	InsertNutritionWaterLevelTemp(ctx context.Context, data *entities.NutritionWaterLevel) *errors.BaseError
	DeleteNutritionWaterLevelTempUntilDate(ctx context.Context, date time.Time) *errors.BaseError
}
