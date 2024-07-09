package repositories

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/entities"
	"hidroponic/internal/module/plants/ports"
	"hidroponic/internal/platform/database"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *database.DatabaseService
}

func New(db *database.DatabaseService) ports.Repository {
	return &Repository{
		DB: db,
	}
}

func (r Repository) GetAllPlant(ctx context.Context) ([]entities.Plant, *errors.BaseError) {
	query := `SELECT "id", "name", "varieties", "plant_type", "generative_age", "harvest_age", "nutrition_min", "nutrition_max", "nutrition_adjustment", "nutrition_targets", "ph_level", "temperature", "plant_age", "current_growth", "status", "created_at", "updated_at", "actived_at" FROM plants ORDER BY "status" ASC`

	var result []entities.Plant
	err := r.DB.SelectContext(ctx, &result, query)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) GetPlantByID(ctx context.Context, id uint) (*entities.Plant, *errors.BaseError) {
	query := `SELECT "id", "name", "varieties", "plant_type", "generative_age", "harvest_age", "nutrition_min", "nutrition_max", "nutrition_adjustment", "nutrition_targets", "ph_level", "temperature", "plant_age", "current_growth", "status", "created_at", "updated_at", "actived_at" FROM plants WHERE "id" = $1`

	var result entities.Plant
	err := r.DB.GetContext(ctx, &result, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("plant not found")
		}
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &result, nil
}

func (r Repository) GetActivePlant(ctx context.Context) (*entities.Plant, *errors.BaseError) {
	query := `SELECT "id", "name", "varieties", "plant_type", "generative_age", "harvest_age", "nutrition_min", "nutrition_max", "nutrition_adjustment", "nutrition_targets", "ph_level", "temperature", "plant_age", "current_growth", "status", "created_at", "updated_at", "actived_at" FROM plants WHERE "status" = $1`

	var result entities.Plant
	err := r.DB.GetContext(ctx, &result, query, constants.StatusActived)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("active plant not found")
		}
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &result, nil
}

func (r Repository) InsertPlant(ctx context.Context, data *entities.Plant) *errors.BaseError {
	plainQuery := `INSERT INTO plants ("name", "varieties", "plant_type", "generative_age", "harvest_age", "nutrition_min", "nutrition_max",
	"nutrition_adjustment", "nutrition_targets", "ph_level", "temperature", "plant_age", "current_growth", "status", "created_at", "updated_at",
	"actived_at") 
	VALUES 
	(:name, :varieties, :plant_type, :generative_age, :harvest_age, :nutrition_min, :nutrition_max, :nutrition_adjustment, :nutrition_targets, :ph_level,
	:temperature, :plant_age, :current_growth, :status, :created_at, :updated_at, :actived_at) RETURNING id`

	query, args, err := sqlx.Named(plainQuery, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}
	query = r.DB.Rebind(query)

	var id uint
	err = r.DB.Get(&id, query, args...)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	data.ID = id

	return nil
}

func (r Repository) UpdatePlant(ctx context.Context, data *entities.Plant) *errors.BaseError {
	query := `UPDATE plants SET "name" = :name, "varieties" = :varieties, "plant_type" = :plant_type, "generative_age" = :generative_age, 
	"harvest_age" = :harvest_age, "nutrition_min" = :nutrition_min, "nutrition_max" = :nutrition_max, "nutrition_adjustment" = :nutrition_adjustment,
	"nutrition_targets" = :nutrition_targets, "ph_level" = :ph_level, "temperature" = :temperature, "plant_age" = :plant_age, 
	"current_growth" = :current_growth, "status" = :status, "created_at" = :created_at, "updated_at" = :updated_at, "actived_at" = :actived_at 
	WHERE "id" = :id`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
