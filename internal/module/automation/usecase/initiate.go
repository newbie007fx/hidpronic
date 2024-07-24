package usecase

import (
	"context"
	"errors"
	"hidroponic/internal/models/sensor"
	"hidroponic/internal/module/automation/constants"
	"hidroponic/internal/module/automation/entities"
	"hidroponic/internal/module/automation/helpers"
	"hidroponic/internal/module/automation/helpers/fuzzy"
	"hidroponic/internal/module/automation/models"
	installationConfEntities "hidroponic/internal/module/installationconfig/entities"
	plantConstants "hidroponic/internal/module/plants/constants"
	plantEntities "hidroponic/internal/module/plants/entities"
	plantHelpers "hidroponic/internal/module/plants/helpers"
	"log"
	"time"
)

func (u Usecase) InitiateAutomation(ctx context.Context) (resp models.InitAutomationResponse, er error) {
	plant, err := u.plantRepo.GetActivePlant(ctx)
	if err != nil {
		return resp, errors.New(err.Error())
	}

	plantAge := plantHelpers.GetPlantAgeInstance().CalculateAgeInDays(plant.ActivatedAt.Time) + plant.PlantAge
	if plantAge > plant.HarvestAge {
		plantAge = plant.HarvestAge
	}
	nutritionTarget := plant.NutritionTargets[plantAge]
	targetPPM := nutritionTarget.TargetPPM

	if plant.PlantType == plantConstants.TypeFruitCrop && plant.CurrentGrowth == plantConstants.GrowthGenerative {
		if nutritionTarget.AdditionalPPM == 0 {
			nutritionTarget.AdditionalPPM = plant.NutritionAdjustment
			plant.NutritionTargets[plantAge] = nutritionTarget
			u.plantRepo.UpdatePlant(ctx, plant)
		}
		targetPPM = targetPPM + nutritionTarget.AdditionalPPM
	} else {
		if nutritionTarget.AdditionalPPM != 0 {
			nutritionTarget.AdditionalPPM = 0
			plant.NutritionTargets[plantAge] = nutritionTarget
			u.plantRepo.UpdatePlant(ctx, plant)
		}
	}

	installationConf, err := u.installationConfRepo.GetInstallationConfig(ctx, 1)
	if err != nil {
		return resp, errors.New(err.Error())
	}

	containerConf, err := u.installationConfRepo.FindContainerConfigByID(ctx, string(sensor.TypeNutritionWaterVolume))
	if err != nil {
		return resp, errors.New(err.Error())
	}

	lastSensorData := helpers.GetLastSensorDataInstance().Get()
	if lastSensorData == nil {
		return resp, errors.New("no sensor data found")
	}

	fuzzyValue := u.getFuzzyValue(targetPPM, plant, installationConf, lastSensorData)

	log.Println("fuzzy value result is ", fuzzyValue)

	if fuzzyValue <= 0.5 {
		return resp, errors.New("fuzzy value is under threshold")
	}

	entity := &entities.Automation{
		PlantID: plant.ID,
		BofereData: entities.SensorData{
			SensorData: *lastSensorData,
		},
		TargetPPM:   targetPPM,
		TriggeredAt: time.Now(),
		Status:      constants.StatusInitiate,
	}

	err = u.repo.InsertAutomation(ctx, entity)
	if err != nil {
		return resp, errors.New(err.Error())
	}

	targetWaterVolume, nutritionTargetVolume := u.calcualteNutritionNeeded(models.CalcualteNutritionNeeded{
		CurrentNutritionWaterVolume: lastSensorData.NutritionWaterVolume,
		TargetNutritionWaterVolume:  containerConf.Volume,
		CurrentNutritionWaterPPM:    lastSensorData.NutritionWaterPPM,
		TargetNutritionWaterPPM:     targetPPM,
		RawWaterPPM:                 installationConf.RawWaterPPM,
		NutritionPPM:                installationConf.NutritionPPM,
	})

	resp.AutomationID = entity.ID
	resp.TargetNutritionVolume = nutritionTargetVolume
	resp.TargetWaterDistance = containerConf.SensorGap
	resp.TargetWaterVolume = targetWaterVolume

	return
}

func (u *Usecase) getFuzzyValue(targetPPM float32, plant *plantEntities.Plant, installationConf *installationConfEntities.InstallationConfig, lastSensorData *models.SensorData) float32 {
	valueRange := installationConf.FuzzyWaterTemperaturePercent * plant.Temperature / 100
	lowTempValue := plant.Temperature - valueRange
	HighTempValue := plant.Temperature + valueRange

	nutritionValueRange := installationConf.FuzzyNutritionWaterLevelPercent * targetPPM / 100
	lowNutritionTarget := targetPPM - nutritionValueRange
	highNutritionTarget := targetPPM + nutritionValueRange

	fis := fuzzy.TsukamotoFIS{
		LowNutritionTarget:      lowNutritionTarget,
		OptimalNutritionTarget:  targetPPM,
		HighNutritionTarget:     highNutritionTarget,
		LowTemperatureValue:     lowTempValue,
		OptimalTemperatureValue: plant.Temperature,
		HighTemperatureValue:    HighTempValue,
		LowWaterVolume:          installationConf.FuzzyNutritionWaterVolumeLow,
		MediumWaterVolume:       installationConf.FuzzyNutritionWaterVolumeMedium,
		HighWaterVolume:         installationConf.FuzzyNutritionWaterVolumeHigh,
	}

	result := fis.Inference(lastSensorData.WaterTemperature, lastSensorData.NutritionWaterPPM, lastSensorData.NutritionWaterVolume)

	return result
}

func (u Usecase) calcualteNutritionNeeded(data models.CalcualteNutritionNeeded) (rawWaterVolume, nutritionVolume float32) {
	neededVolume := data.TargetNutritionWaterVolume - data.CurrentNutritionWaterVolume
	neededSolution := data.TargetNutritionWaterVolume*data.TargetNutritionWaterPPM - data.CurrentNutritionWaterVolume*data.CurrentNutritionWaterPPM

	matrixA := [2][2]float32{
		{1, 1},
		{data.RawWaterPPM, data.NutritionPPM},
	}

	matrixB := [2]float32{neededVolume, neededSolution}

	detA := matrixA[0][0]*matrixA[1][1] - matrixA[0][1]*matrixA[1][0]

	invA := [2][2]float32{
		{matrixA[1][1] / detA, -matrixA[0][1] / detA},
		{-matrixA[1][0] / detA, matrixA[0][0] / detA},
	}

	rawWaterVolume = invA[0][0]*matrixB[0] + invA[0][1]*matrixB[1]
	nutritionVolume = invA[1][0]*matrixB[0] + invA[1][1]*matrixB[1]

	return
}
