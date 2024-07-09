package plant

import (
	"hidroponic/internal/module/plants/ports"
)

type PlantHandlers struct {
	plantUsecase ports.Usecase
}

func New(plantUsecase ports.Usecase) *PlantHandlers {
	return &PlantHandlers{
		plantUsecase: plantUsecase,
	}
}
