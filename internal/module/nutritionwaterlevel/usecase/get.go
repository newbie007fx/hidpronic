package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	"hidroponic/internal/module/plants/helpers"
)

func (u Usecase) GetActivePlantNutritionWaterLevel(ctx context.Context) (resp []models.NutritionWaterLevel, err *errors.BaseError) {
	plantID := helpers.GetActivePlantIDInstance().Get()
	if plantID == 0 {
		return resp, nil
	}

	limit := 144
	result, err := u.repo.GetNutritionWaterLevelByPlantIDWithLimit(ctx, plantID, limit)
	if err != nil {
		return
	}

	resp = []models.NutritionWaterLevel{}
	for _, res := range result {
		resp = append(resp, *u.mapEntityToModel(&res))
	}

	return
}
