package usecase

import (
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/entities"
	"hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/module/plants/models"
	"hidroponic/internal/module/plants/ports"
)

type Usecase struct {
	repo ports.Repository
}

func New(repo ports.Repository) ports.Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u Usecase) mapEntityToModel(entity *entities.Plant) *models.Plant {
	model := &models.Plant{
		ID:                  entity.ID,
		Name:                entity.Name,
		Varieties:           entity.Varieties,
		PlantType:           entity.PlantType,
		GenerativeAge:       entity.GenerativeAge,
		HarvestAge:          entity.HarvestAge,
		NutritionMin:        entity.NutritionMin,
		NutritionMax:        entity.NutritionMax,
		NutritionAdjustment: entity.NutritionAdjustment,
		NutritionTargets:    u.getNutritionTargetModel(entity.NutritionTargets),
		PHLevel:             entity.PHLevel,
		Temperature:         entity.Temperature,
		PlantAge:            entity.PlantAge,
		CurrentGrowth:       entity.CurrentGrowth,
		Status:              entity.Status,
		CreatedAt:           entity.CreatedAt,
		UpdatedAt:           entity.UpdatedAt,
	}

	if entity.ActivedAt.Valid && entity.Status == constants.StatusActived {
		model.ActivedAt = &entity.ActivedAt.Time
		model.CurrentAge = model.PlantAge + helpers.GetPlantAgeInstance().CalculateAgeInDays(entity.ActivedAt.Time)
	}

	return model
}

func (Usecase) getNutritionTargetModel(data entities.NutritionTargetMap) []models.NutritionTarget {
	nutritionTargets := []models.NutritionTarget{}
	for key, value := range data {
		target := models.NutritionTarget{
			PlantAge:      key,
			TargetPPM:     value.TargetPPM,
			AdditionalPPM: value.AdditionalPPM,
		}

		nutritionTargets = append(nutritionTargets, target)
	}
	return nutritionTargets
}

func (Usecase) getNutritionTargetEntity(data []models.NutritionTarget) entities.NutritionTargetMap {
	nutritionTargetMap := entities.NutritionTargetMap{}
	for _, value := range data {
		target := entities.NutritionTarget{
			TargetPPM:     value.TargetPPM,
			AdditionalPPM: value.AdditionalPPM,
		}

		nutritionTargetMap[value.PlantAge] = target
	}
	return nutritionTargetMap
}
