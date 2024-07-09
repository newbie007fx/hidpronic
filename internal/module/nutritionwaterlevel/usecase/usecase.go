package usecase

import (
	"hidroponic/internal/module/nutritionwaterlevel/entities"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	"hidroponic/internal/module/nutritionwaterlevel/ports"
	plantPorts "hidroponic/internal/module/plants/ports"
)

type Usecase struct {
	repo      ports.Repository
	plantRepo plantPorts.Repository
}

func New(repo ports.Repository, plantRepo plantPorts.Repository) ports.Usecase {
	return &Usecase{
		repo:      repo,
		plantRepo: plantRepo,
	}
}

func (Usecase) mapEntityToModel(entity *entities.NutritionWaterLevel) *models.NutritionWaterLevel {
	return &models.NutritionWaterLevel{
		Value:     entity.Value,
		PlantID:   entity.PlantID,
		CreatedAt: entity.CreatedAt.UnixMilli(),
	}
}
