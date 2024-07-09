package repositories

import (
	"hidroponic/internal/module/nutritionwaterlevel/ports"
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
