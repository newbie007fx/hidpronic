package repositories

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/nutritionwaterlevel/entities"
	"time"
)

func (r Repository) GetNutritionWaterLevelTempByPlantIDWithByRange(ctx context.Context, plantID uint, startDate, endDate time.Time) ([]entities.NutritionWaterLevel, *errors.BaseError) {
	query := `SELECT "id", "plant_id", "value", "created_at" FROM nutrition_water_levels_temp WHERE "plant_id" = $1 AND "created_at" >= $2 AND "created_at" < $3`

	var result []entities.NutritionWaterLevel
	err := r.DB.SelectContext(ctx, &result, query, plantID, startDate, endDate)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) InsertNutritionWaterLevelTemp(ctx context.Context, data *entities.NutritionWaterLevel) *errors.BaseError {
	query := `INSERT INTO nutrition_water_levels_temp ("plant_id", "value", "created_at") VALUES (:plant_id, :value, :created_at)`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}

func (r Repository) DeleteNutritionWaterLevelTempUntilDate(ctx context.Context, date time.Time) *errors.BaseError {
	query := `DELETE FROM nutrition_water_levels_temp WHERE "created_at" <= $1`

	_, err := r.DB.ExecContext(ctx, query, date)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
