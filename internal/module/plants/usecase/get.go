package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/models"
)

func (u Usecase) GetAllPlant(ctx context.Context) (resp []models.BasicPlant, err *errors.BaseError) {
	result, err := u.repo.GetAllPlant(ctx)
	if err != nil {
		return
	}

	resp = []models.BasicPlant{}
	for _, res := range result {
		basicPlant := models.BasicPlant{
			ID:         res.ID,
			Name:       res.Name,
			Varieties:  res.Varieties,
			PlantType:  res.PlantType,
			HarvestAge: res.HarvestAge,
			Status:     res.Status,
			CreatedAt:  res.CreatedAt,
		}
		resp = append(resp, basicPlant)
	}

	return
}

func (u Usecase) GetActivePlant(ctx context.Context) (resp *models.Plant, err *errors.BaseError) {
	result, err := u.repo.GetActivePlant(ctx)
	if err != nil {
		return
	}

	resp = u.mapEntityToModel(result)

	return
}

func (u Usecase) GetPlantByID(ctx context.Context, id uint) (resp *models.Plant, err *errors.BaseError) {
	result, err := u.repo.GetPlantByID(ctx, id)
	if err != nil {
		return
	}

	resp = u.mapEntityToModel(result)

	return
}
