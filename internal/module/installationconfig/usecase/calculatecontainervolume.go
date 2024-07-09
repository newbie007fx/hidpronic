package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/installationconfig/entities"
)

func (u *Usecase) CalculateContainerVolume(ctx context.Context, ID string, distance float32) (float32, *errors.BaseError) {
	containerConfig, err := u.repo.FindContainerConfigByID(ctx, ID)
	if err != nil {
		return 0, err
	}

	return u.calculateContainerVolumeByConfig(*containerConfig, distance)
}

func (u *Usecase) calculateContainerVolumeByConfig(entity entities.ContainerConfig, distance float32) (volume float32, err *errors.BaseError) {
	finalHeight := entity.Height
	if distance > 0 {
		finalDistance := distance - entity.SensorGap
		finalHeight = finalHeight - finalDistance
	}
	area := entity.BottomArea + ((entity.TopArea-entity.BottomArea)/entity.Height)*finalHeight
	volume = 0.5 * (entity.BottomArea + area) * finalHeight

	return
}
