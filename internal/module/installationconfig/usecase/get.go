package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/models/sensor"
	"hidroponic/internal/module/installationconfig/models"
)

const DefaultID uint = 1

func (u Usecase) GetInstallationConfig(ctx context.Context) (*models.InstallationConfigResponse, *errors.BaseError) {
	var resp models.InstallationConfigResponse

	installation, err := u.repo.GetInstallationConfig(ctx, DefaultID)
	if err != nil {
		return nil, err
	}

	resp.ID = installation.ID
	resp.NutritionPPM = installation.NutritionPPM
	resp.RawWaterPPM = installation.RawWaterPPM
	resp.FuzzyNutritionWaterLevelPercent = installation.FuzzyNutritionWaterLevelPercent
	resp.FuzzyWaterTemperaturePercent = installation.FuzzyWaterTemperaturePercent
	resp.FuzzyNutritionWaterVolumeLow = installation.FuzzyNutritionWaterVolumeLow
	resp.FuzzyNutritionWaterVolumeMedium = installation.FuzzyNutritionWaterVolumeMedium
	resp.FuzzyNutritionWaterVolumeHigh = installation.FuzzyNutritionWaterVolumeHigh

	containers, err := u.repo.GetAllContainerConfig(ctx)
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		if container.ID == string(sensor.TypeNutritionWaterVolume) {
			resp.NutritionWaterContainer = *u.mapContainerConfigEntityToModel(&container)
		} else {
			resp.RawWaterContainer = *u.mapContainerConfigEntityToModel(&container)
		}
	}

	return &resp, nil
}
