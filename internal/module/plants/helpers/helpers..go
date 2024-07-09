package helpers

import (
	"hidroponic/internal/module/plants/ports"
	"time"

	"github.com/patrickmn/go-cache"
)

func InitHelpers(plantUsecase ports.Usecase) {
	activePlantIdInstance = &activePlantID{
		caching:      cache.New(7*time.Minute, 10*time.Minute),
		plantUsecase: plantUsecase,
	}
}
