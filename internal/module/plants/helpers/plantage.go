package helpers

import (
	"time"
)

var plantAgeInstance *plantAge

type plantAge struct {
}

func (ap *plantAge) CalculateAgeInDays(plantTime time.Time) int {
	currentTime := time.Now()
	currentTime = currentTime.Truncate(24 * time.Hour)

	plantTime = plantTime.Truncate(24 * time.Hour)

	diff := currentTime.Sub(plantTime)

	return int(diff.Hours() / 24)
}

func GetPlantAgeInstance() *plantAge {
	return plantAgeInstance
}
