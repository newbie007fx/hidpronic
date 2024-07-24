package usecase

import (
	"context"
	"database/sql"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/entities"
	"hidroponic/internal/module/plants/models"
	"time"
)

func (u *Usecase) InsertPlant(ctx context.Context, createPlant models.CreatePlant) (resp *models.Plant, err *errors.BaseError) {
	nutritionTargets := u.calculateNutritionTarget(&createPlant)

	entity := &entities.Plant{
		Name: createPlant.Name,
		Description: sql.NullString{
			String: createPlant.Description,
			Valid:  true,
		},
		Varieties:           createPlant.Varieties,
		PlantType:           createPlant.PlantType,
		GenerativeAge:       createPlant.GenerativeAge,
		HarvestAge:          createPlant.HarvestAge,
		NutritionMin:        createPlant.NutritionMin,
		NutritionMax:        createPlant.NutritionMax,
		NutritionAdjustment: createPlant.NutritionAdjustment,
		NutritionTargets:    u.getNutritionTargetEntity(nutritionTargets),
		PHLevel:             createPlant.PHLevel,
		Temperature:         createPlant.Temperature,
		PlantAge:            createPlant.PlantAge,
		CurrentGrowth:       constants.GrowthVegetative,
		Status:              constants.StatusCreated,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	err = u.repo.InsertPlant(ctx, entity)
	if err != nil {
		return
	}

	return u.mapEntityToModel(entity), nil
}

func (u *Usecase) calculateNutritionTarget(createPlant *models.CreatePlant) []models.NutritionTarget {
	x := createPlant.NutritionMax - createPlant.NutritionMin
	um := createPlant.HarvestAge - createPlant.PlantAge

	var nutritionTargets []models.NutritionTarget = []models.NutritionTarget{}
	for i := 0; i <= um; i++ {
		target := models.NutritionTarget{
			PlantAge:  createPlant.PlantAge + i,
			TargetPPM: createPlant.NutritionMin + x*u.getGrowthPercentage(i, um)/100,
		}

		nutritionTargets = append(nutritionTargets, target)
	}

	return nutritionTargets
}

func (Usecase) getGrowthPercentage(age, maxAge int) float32 {
	percent := float32(age) / float32(maxAge) * 100

	if percent <= 35 {
		return percent + (percent / 35 * 5)
	} else if percent <= 45 {
		return 45
	} else if percent <= 65 {
		return percent + ((percent - 45) / 20 * 5)
	} else if percent < 75 {
		return 75
	} else if percent <= 90 {
		return percent + ((percent - 75) / 15 * 5)
	} else {
		return 100
	}
}
