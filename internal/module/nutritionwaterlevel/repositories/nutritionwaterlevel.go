package repositories

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/entities"
)

func (r Repository) GetNutritionWaterLevelByPlantIDWithLimit(ctx context.Context, plantID uint, limit int) ([]entities.NutritionWaterLevel, *errors.BaseError) {
	query := `SELECT "id", "value", "plant_id", "created_at" FROM nutrition_water_levels WHERE "plant_id" = $1 ORDER BY "created_at" ASC LIMIT $2`

	var result []entities.NutritionWaterLevel
	err := r.DB.SelectContext(ctx, &result, query, plantID, limit)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) InsertNutritionWaterLevel(ctx context.Context, data *entities.NutritionWaterLevel) *errors.BaseError {
	query := `INSERT INTO nutrition_water_levels ("value", "plant_id", "created_at") VALUES (:value, :plant_id, :created_at)`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
