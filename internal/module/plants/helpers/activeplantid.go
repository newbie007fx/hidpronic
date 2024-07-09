package helpers

import (
	"context"
	"fmt"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/ports"
	"log/slog"

	"github.com/patrickmn/go-cache"
)

var activePlantIdInstance *activePlantID

const cacheKey = "active-plant-id"

type activePlantID struct {
	plantUsecase ports.Usecase
	caching      *cache.Cache
}

func (ap *activePlantID) Get() uint {
	if value, ok := ap.caching.Get(cacheKey); ok {
		return value.(uint)
	}

	var plantID uint = 0
	plant, err := ap.plantUsecase.GetActivePlant(context.Background())
	if err != nil {
		if err.Code != errors.ErrorQueryNoRow {
			slog.Error(fmt.Sprintf("errror get active plant with message: %s", err.Message))
			return 0
		}
	} else {
		plantID = plant.ID
	}

	ap.caching.Set(cacheKey, plantID, cache.DefaultExpiration)

	return plantID
}

func (ap *activePlantID) Set(ID uint) {
	ap.caching.Set(cacheKey, ID, cache.DefaultExpiration)
}

func GetActivePlantIDInstance() *activePlantID {
	return activePlantIdInstance
}
