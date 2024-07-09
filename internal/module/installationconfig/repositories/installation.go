package repositories

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/entities"
)

func (r Repository) UpdateInstallationConfig(ctx context.Context, data *entities.InstallationConfig) *errors.BaseError {
	query := `UPDATE installation_configs SET "nutrition_ppm" = :nutrition_ppm, "raw_water_ppm" = :raw_water_ppm, "fuzzy_nutrition_water_level_percent" = :fuzzy_nutrition_water_level_percent, "fuzzy_water_temperature_percent" = :fuzzy_water_temperature_percent, "fuzzy_nutrition_water_volume_low" = :fuzzy_nutrition_water_volume_low, "fuzzy_nutrition_water_volume_medium" = :fuzzy_nutrition_water_volume_medium, "fuzzy_nutrition_water_volume_high" = :fuzzy_nutrition_water_volume_high WHERE "id" = :id`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}

func (r Repository) GetInstallationConfig(ctx context.Context, ID uint) (*entities.InstallationConfig, *errors.BaseError) {
	query := `SELECT "id", "nutrition_ppm", "raw_water_ppm", "fuzzy_nutrition_water_level_percent", "fuzzy_water_temperature_percent", "fuzzy_nutrition_water_volume_low", "fuzzy_nutrition_water_volume_medium", "fuzzy_nutrition_water_volume_high" FROM installation_configs WHERE "id" = $1`

	var installationConfig entities.InstallationConfig
	err := r.DB.GetContext(ctx, &installationConfig, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("container config not found")
		}

		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &installationConfig, nil
}
