package usecase

import (
	"context"
	"database/sql"
	commonConstants "hidroponic/internal/constants"
	"hidroponic/internal/errors"
	"hidroponic/internal/helpers"
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/entities"
	plantHelpers "hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/module/plants/models"
	"time"
)

func (u *Usecase) UpdatePlantStatus(ctx context.Context, data models.UpdatePlantStatus) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, data.ID)
	if err != nil {
		return err
	}

	previousStatus := plant.Status
	if previousStatus == data.Status {
		return nil
	}

	if err := plant.ValidateStatus(data.Status); err != nil {
		return errors.ErrorValidation.New("status not allowed to update")
	}

	plant.Status = data.Status
	if plant.Status == constants.StatusActived {
		plant.ActivedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	plant.UpdatedAt = time.Now()
	err = u.repo.UpdatePlant(ctx, plant)
	if err != nil {
		return err
	}

	var id uint = 0
	state := commonConstants.StateOff
	if previousStatus == constants.StatusActived {
		id = plant.ID
		state = commonConstants.StateOn
	}

	plantHelpers.GetActivePlantIDInstance().Set(id)
	helpers.GetDeviceStateInstance().SetState(state).PublishState(nil)

	return nil
}

func (u *Usecase) UpdatePlant(ctx context.Context, data models.UpdatePlant) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if plant.Status == constants.StatusActived {
		return errors.ErrorActionFobidden.New("not allowed to update")
	}

	plant.Name = data.Name
	plant.Varieties = data.Varieties
	plant.PlantType = data.PlantType
	plant.GenerativeAge = data.GenerativeAge
	plant.HarvestAge = data.HarvestAge
	plant.NutritionMin = data.NutritionMin
	plant.NutritionMax = data.NutritionMax
	plant.NutritionAdjustment = data.NutritionAdjustment
	u.updateNutritionTarget(plant.NutritionTargets, data.NutritionTargets)
	plant.PHLevel = data.PHLevel
	plant.Temperature = data.Temperature
	plant.PlantAge = data.PlantAge
	plant.UpdatedAt = time.Now()

	return u.repo.UpdatePlant(ctx, plant)
}

func (Usecase) updateNutritionTarget(nutritionMap entities.NutritionTargetMap, nutritionTargets []models.NutritionTarget) {
	for _, target := range nutritionTargets {
		if val, ok := nutritionMap[target.PlantAge]; ok {
			val.TargetPPM = target.TargetPPM
			nutritionMap[target.PlantAge] = val
		}
	}
}
