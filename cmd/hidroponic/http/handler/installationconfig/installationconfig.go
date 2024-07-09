package installationconfig

import (
	"hidroponic/internal/module/installationconfig/ports"
)

type InstallationConfigHandlers struct {
	instalaltionConfigUsecase ports.Usecase
}

func New(instalaltionConfigUsecase ports.Usecase) *InstallationConfigHandlers {
	return &InstallationConfigHandlers{
		instalaltionConfigUsecase: instalaltionConfigUsecase,
	}
}
