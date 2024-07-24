package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/constants"
)

func (u *Usecase) DeletePlant(ctx context.Context, id uint) *errors.BaseError {
	plant, err := u.repo.GetPlantByID(ctx, id)
	if err != nil {
		return err
	}

	if plant.Status == constants.StatusActivated {
		return errors.ErrorActionFobidden.New("active plant is not allowed to delete")
	}

	return u.repo.DeletePlant(ctx, id)
}
