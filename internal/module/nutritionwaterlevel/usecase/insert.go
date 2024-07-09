package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/entities"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	"time"
)

func (u *Usecase) InsertNutritionWaterLevelTemp(ctx context.Context, req models.CreateNutritionWaterLevel) *errors.BaseError {
	entity := &entities.NutritionWaterLevel{
		Value:     req.Value,
		PlantID:   req.PlantID,
		CreatedAt: time.Now(),
	}

	return u.repo.InsertNutritionWaterLevelTemp(ctx, entity)
}

func (u *Usecase) InsertNutritionWaterLevelFromTempRange(ctx context.Context, plantID uint, startDate, endDate time.Time) (resp *models.NutritionWaterLevel, err *errors.BaseError) {
	result, err := u.repo.GetNutritionWaterLevelTempByPlantIDWithByRange(ctx, plantID, startDate, endDate)
	if err != nil {
		return
	}

	if len(result) == 0 {
		return resp, errors.ErrorInternalServer.New("empty result")
	}

	var totalValue float32 = 0
	for _, row := range result {
		totalValue += row.Value
	}

	entity := &entities.NutritionWaterLevel{
		Value:     totalValue / float32(len(result)),
		CreatedAt: time.Now(),
	}

	err = u.repo.InsertNutritionWaterLevel(ctx, entity)
	if err != nil {
		return
	}

	err = u.repo.DeleteNutritionWaterLevelTempUntilDate(ctx, endDate)
	if err != nil {
		return
	}

	return u.mapEntityToModel(entity), nil
}
