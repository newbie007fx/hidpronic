package repositories

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/entities"
)

func (r Repository) FindContainerConfigByID(ctx context.Context, ContainerID string) (*entities.ContainerConfig, *errors.BaseError) {
	query := `SELECT "id", "name", "sensor_gap", "height", "bottom_area", "top_area", "volume" FROM container_configs WHERE "id" = $1`

	var containerConfig entities.ContainerConfig
	err := r.DB.GetContext(ctx, &containerConfig, query, ContainerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("container config not found")
		}

		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &containerConfig, nil
}

func (r Repository) GetAllContainerConfig(ctx context.Context) ([]entities.ContainerConfig, *errors.BaseError) {
	query := `SELECT "id", "name", "sensor_gap", "height", "bottom_area", "top_area", "volume" FROM container_configs`

	var result []entities.ContainerConfig
	err := r.DB.SelectContext(ctx, &result, query)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) UpdateContainerConfig(ctx context.Context, data *entities.ContainerConfig) *errors.BaseError {
	query := `UPDATE container_configs SET "name" = :name, "sensor_gap" = :sensor_gap, "height" = :height, "bottom_area" = :bottom_area, "top_area" = :top_area, "volume" = :volume WHERE "id" = :id`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
