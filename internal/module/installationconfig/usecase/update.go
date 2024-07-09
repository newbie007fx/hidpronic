package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/entities"
	"hidroponic/internal/module/installationconfig/models"
)

func (u Usecase) UpdateInstallationConfig(ctx context.Context, data models.UpdateInstallationConfig) *errors.BaseError {
	installation, err := u.repo.GetInstallationConfig(ctx, DefaultID)
	if err != nil {
		return err
	}
	installation.NutritionPPM = data.NutritionPPM
	installation.RawWaterPPM = data.RawWaterPPM
	err = u.repo.UpdateInstallationConfig(ctx, installation)
	if err != nil {
		return err
	}

	nutritionWaterContainer, err := u.repo.FindContainerConfigByID(ctx, data.NutritionWaterContainer.ID)
	if err != nil {
		return err
	}
	u.mapContainerConfigRequestToEntity(data.NutritionWaterContainer, nutritionWaterContainer)
	err = u.repo.UpdateContainerConfig(ctx, nutritionWaterContainer)
	if err != nil {
		return err
	}

	rawWaterContainer, err := u.repo.FindContainerConfigByID(ctx, data.RawWaterContainer.ID)
	if err != nil {
		return err
	}
	u.mapContainerConfigRequestToEntity(data.RawWaterContainer, rawWaterContainer)
	err = u.repo.UpdateContainerConfig(ctx, rawWaterContainer)

	return err
}

func (u Usecase) mapContainerConfigRequestToEntity(data models.UpdateContainerConfig, entity *entities.ContainerConfig) {
	entity.SensorGap = data.SensorGap
	entity.Height = data.Height
	entity.BottomArea = data.BottomArea
	entity.TopArea = data.TopArea
	entity.Volume, _ = u.calculateContainerVolumeByConfig(*entity, 0)
}
