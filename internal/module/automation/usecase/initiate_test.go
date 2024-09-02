package usecase_test

import (
	"hidroponic/internal/module/automation/models"
	"testing"
)

func TestCalculateNutritionNeeded(t *testing.T) {
	data := models.CalculateNutritionNeeded{
		CurrentNutritionWaterVolume: 200,
		TargetNutritionWaterVolume:  6500,
		CurrentNutritionWaterPPM:    100,
		TargetNutritionWaterPPM:     200,
		RawWaterPPM:                 190,
		NutritionPPM:                500,
	}

	waterVolume, nutritionTargetVolume := automationUsc.CalculateNutritionNeeded(data)

	t.Log("nutrition ", nutritionTargetVolume)
	t.Log("water ", waterVolume)
}
