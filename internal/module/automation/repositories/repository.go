package repositories

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/automation/constants"
	"hidroponic/internal/module/automation/entities"
	"hidroponic/internal/module/automation/ports"
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

func (r Repository) GetAllAutomationWithLimit(ctx context.Context, limit int) ([]entities.Automation, *errors.BaseError) {
	query := `SELECT "id", "plant_id", "after_data", "before_data", "accuration", "target_ppm", "duration", "status", "triggered_at", "finished_at" FROM automation WHERE "status" = $1 ORDER BY "id" DESC LIMIT $2`

	var result []entities.Automation
	err := r.DB.SelectContext(ctx, &result, query, constants.StatusComplete, limit)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) GetAutomationByID(ctx context.Context, id uint) (*entities.Automation, *errors.BaseError) {
	query := `SELECT "id", "plant_id", "after_data", "before_data", "accuration", "target_ppm", "duration", "status", "triggered_at", "finished_at" FROM automation WHERE "id" = $1`

	var result entities.Automation
	err := r.DB.GetContext(ctx, &result, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("automation not found")
		}
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &result, nil
}

func (r Repository) InsertAutomation(ctx context.Context, data *entities.Automation) *errors.BaseError {
	plainQuery := `INSERT INTO automation ("plant_id", "after_data", "before_data", "accuration", "target_ppm", "duration", "status", "triggered_at", "finished_at") 
	VALUES 
	(:plant_id, :after_data, :before_data, :accuration, :target_ppm", :duration, status, :triggered_at, :finished_at) RETURNING id`

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

func (r Repository) UpdateAutomation(ctx context.Context, data *entities.Automation) *errors.BaseError {
	query := `UPDATE automation SET "plant_id" = :plant_id, "after_data" = :aftar_data, "before_data" = :berfore_data, "accuration" = :accuration, "target_ppm" = :target_ppm, "duration" = :duration, "status" = :status, "triggered_at"  :triggered_at, "finished_at" = :finished_at 
	WHERE "id" = :id`

	_, err := r.DB.NamedExecContext(ctx, query, data)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
