package usecase

import (
	"hidroponic/internal/module/installationconfig/entities"
	"hidroponic/internal/module/installationconfig/models"
	"hidroponic/internal/module/installationconfig/ports"
)

type Usecase struct {
	repo ports.Repository
}

func New(repo ports.Repository) ports.Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u Usecase) mapContainerConfigEntityToModel(entity *entities.ContainerConfig) *models.ContainerConfig {
	return &models.ContainerConfig{
		ID:         entity.ID,
		Name:       entity.Name,
		SensorGap:  entity.SensorGap,
		Height:     entity.Height,
		BottomArea: entity.BottomArea,
		TopArea:    entity.TopArea,
		Volume:     entity.Volume,
	}
}
