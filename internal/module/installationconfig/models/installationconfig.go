package models

import "hidroponic/internal/platform/validation"

type InstallationConfig struct {
	ID                              uint    `json:"id"`
	NutritionPPM                    float32 `json:"nutrition_ppm"`
	RawWaterPPM                     float32 `json:"raw_water_ppm"`
	FuzzyNutritionWaterLevelPercent float32 `json:"fuzzy_nutrition_water_level_percent"`
	FuzzyWaterTemperaturePercent    float32 `json:"fuzzy_water_temperature_percent"`
	FuzzyNutritionWaterVolumeLow    float32 `json:"fuzzy_nutrition_water_volume_low"`
	FuzzyNutritionWaterVolumeMedium float32 `json:"fuzzy_nutrition_water_volume_medium"`
	FuzzyNutritionWaterVolumeHigh   float32 `json:"fuzzy_nutrition_water_volume_high"`
}

type InstallationConfigResponse struct {
	ID                              uint            `json:"id"`
	NutritionPPM                    float32         `json:"nutrition_ppm"`
	RawWaterPPM                     float32         `json:"raw_water_ppm"`
	RawWaterContainer               ContainerConfig `json:"raw_water_container"`
	NutritionWaterContainer         ContainerConfig `json:"nutrition_water_container"`
	FuzzyNutritionWaterLevelPercent float32         `json:"fuzzy_nutrition_water_level_percent"`
	FuzzyWaterTemperaturePercent    float32         `json:"fuzzy_water_temperature_percent"`
	FuzzyNutritionWaterVolumeLow    float32         `json:"fuzzy_nutrition_water_volume_low"`
	FuzzyNutritionWaterVolumeMedium float32         `json:"fuzzy_nutrition_water_volume_medium"`
	FuzzyNutritionWaterVolumeHigh   float32         `json:"fuzzy_nutrition_water_volume_high"`
}

type UpdateInstallationConfig struct {
	ID                              uint                  `validate:"required" json:"id"`
	NutritionPPM                    float32               `validate:"required" json:"nutrition_ppm"`
	RawWaterPPM                     float32               `validate:"required" json:"raw_water_ppm"`
	RawWaterContainer               UpdateContainerConfig `validate:"required" json:"raw_water_container"`
	NutritionWaterContainer         UpdateContainerConfig `validate:"required" json:"nutrition_water_container"`
	FuzzyNutritionWaterLevelPercent float32               `validate:"required" json:"fuzzy_nutrition_water_level_percent"`
	FuzzyWaterTemperaturePercent    float32               `validate:"required" json:"fuzzy_water_temperature_percent"`
	FuzzyNutritionWaterVolumeLow    float32               `validate:"required" json:"fuzzy_nutrition_water_volume_low"`
	FuzzyNutritionWaterVolumeMedium float32               `validate:"required" json:"fuzzy_nutrition_water_volume_medium"`
	FuzzyNutritionWaterVolumeHigh   float32               `validate:"required" json:"fuzzy_nutrition_water_volume_high"`
}

func (uic *UpdateInstallationConfig) Validate() error {
	return validation.Validate(uic)
}
