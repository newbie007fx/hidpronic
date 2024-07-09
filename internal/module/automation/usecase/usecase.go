package usecase

import (
	"hidroponic/internal/module/automation/ports"
	installationConfPorts "hidroponic/internal/module/installationconfig/ports"
	plantPorts "hidroponic/internal/module/plants/ports"
)

type Usecase struct {
	installationConfRepo installationConfPorts.Repository
	plantRepo            plantPorts.Repository
	repo                 ports.Repository
}

func New(repo ports.Repository, installationConfRepo installationConfPorts.Repository, plantRepo plantPorts.Repository) ports.Usecase {
	return &Usecase{
		installationConfRepo: installationConfRepo,
		plantRepo:            plantRepo,
		repo:                 repo,
	}
}
