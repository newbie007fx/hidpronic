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
	if previousStatus == constants.StatusCreated && plant.Status == constants.StatusActivated {
		plant.ActivatedAt = sql.NullTime{
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
	if plant.Status == constants.StatusActivated {
		id = plant.ID
		state = commonConstants.StateOn
	}

	plantHelpers.GetActivePlantIDInstance().Set(id)
	helpers.GetDeviceStateInstance().SetState(state).PublishState(nil)

	return nil
}

func (u *Usecase) HarvestPlant(ctx context.Context, data models.HarvestPlant) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, data.ID)
	if err != nil {
		return err
	}

	previousStatus := plant.Status
	if previousStatus == constants.StatusHarvested {
		return nil
	}

	if err := plant.ValidateStatus(constants.StatusHarvested); err != nil {
		return errors.ErrorValidation.New("status not allowed to harvest")
	}

	plant.Status = constants.StatusHarvested
	plant.Yields = data.Yields
	plant.HarvestedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	plant.UpdatedAt = time.Now()
	err = u.repo.UpdatePlant(ctx, plant)
	if err != nil {
		return err
	}

	var id uint = 0
	state := commonConstants.StateOff

	plantHelpers.GetActivePlantIDInstance().Set(id)
	helpers.GetDeviceStateInstance().SetState(state).PublishState(nil)

	return nil
}

func (u *Usecase) UpdatePlant(ctx context.Context, data models.UpdatePlant) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, data.ID)
	if err != nil {
		return err
	}

	plant.Name = data.Name
	plant.Description = sql.NullString{String: data.Description, Valid: true}
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

func (u *Usecase) UpdatePlantGrowth(ctx context.Context, id uint) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, id)
	if err != nil {
		return err
	}

	if plant.PlantType == constants.TypeLeafCrop {
		return errors.ErrorValidation.New("action only applicable for fruit crop")
	}

	if plant.Status != constants.StatusActivated {
		return errors.ErrorValidation.New("action only applicable for active plant")
	}

	if plant.CurrentGrowth == constants.GrowthGenerative {
		return nil
	}

	plant.CurrentGrowth = constants.GrowthGenerative
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
