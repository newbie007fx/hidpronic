package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"hidroponic/internal/errors"
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

func (r Repository) GetAllAutomation(ctx context.Context, limit, offset int, filter map[string]string) ([]entities.Automation, *errors.BaseError) {
	query := `SELECT "plants"."name" as "plants.name", "plants"."varieties" as "plants.varieties", "automation"."id", "automation"."plant_id", "automation"."target_ppm", "automation"."accuration", "automation"."duration", "automation"."status", "automation"."triggered_at", "automation"."finished_at" 
	FROM automation JOIN plants ON "automation"."plant_id" = "plants"."id"`

	args := []any{}
	for key, val := range filter {
		operation := "AND"
		if len(args) == 0 {
			operation = "WHERE"
		}

		if key == "status" {
			key = `"automation"."status"`
		}

		args = append(args, val)
		query = fmt.Sprintf(`%s %s %s = $%d`, query, operation, key, len(args))
	}

	args = append(args, limit, offset)
	query = fmt.Sprintf(`%s ORDER BY "automation"."id" DESC LIMIT $%d OFFSET $%d`, query, len(args)-1, len(args))

	var result []entities.Automation
	err := r.DB.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return result, errors.ErrorQueryDatabase.New(err.Error())
	}

	return result, nil
}

func (r Repository) GetAutomationByID(ctx context.Context, id uint) (*entities.Automation, *errors.BaseError) {
	query := `SELECT "plants"."name" as "plants.name", "plants"."varieties" as "plants.varieties", "automation"."id", "automation"."plant_id", "automation"."after_data", "automation"."before_data", "automation"."accuration", "automation"."target_ppm", "automation"."duration", "automation"."status", "automation"."triggered_at", "automation"."finished_at" 
	FROM automation JOIN plants ON "automation"."plant_id" = "plants"."id" WHERE "automation"."id" = $1`

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

func (r Repository) DeleteAutomationByPlantID(ctx context.Context, plantID uint) *errors.BaseError {
	query := `DELETE FROM automation "plant_id" = $1`

	_, err := r.DB.ExecContext(ctx, query, plantID)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}
