package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/models"
)

func (u Usecase) GetAllPlant(ctx context.Context) (resp []models.Plant, err *errors.BaseError) {
	result, err := u.repo.GetAllPlant(ctx)
	if err != nil {
		return
	}

	resp = []models.Plant{}
	for _, res := range result {
		resp = append(resp, *u.mapEntityToModel(&res))
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
