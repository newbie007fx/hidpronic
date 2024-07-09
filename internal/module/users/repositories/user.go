package repositories

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/users/entities"
	"hidroponic/internal/module/users/ports"
	"hidroponic/internal/platform/database"
)

type Repository struct {
	DB *database.DatabaseService
}

func New(db *database.DatabaseService) ports.Repository {
	return &Repository{
		DB: db,
	}
}

func (r Repository) FindByUsername(ctx context.Context, username string) (*entities.User, *errors.BaseError) {
	query := `SELECT "id", "name", "username", "email", "password", "email", "created_at" FROM users WHERE "username" = $1`

	var user entities.User
	err := r.DB.GetContext(ctx, &user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("user not found")
		}

		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &user, nil
}
