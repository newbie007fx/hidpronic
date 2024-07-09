package entities

type InstallationConfig struct {
	ID                              uint    `db:"id"`
	NutritionPPM                    float32 `db:"nutrition_ppm"`
	RawWaterPPM                     float32 `db:"raw_water_ppm"`
	FuzzyNutritionWaterLevelPercent float32 `db:"fuzzy_nutrition_water_level_percent"`
	FuzzyWaterTemperaturePercent    float32 `db:"fuzzy_water_temperature_percent"`
	FuzzyNutritionWaterVolumeLow    float32 `db:"fuzzy_nutrition_water_volume_low"`
	FuzzyNutritionWaterVolumeMedium float32 `db:"fuzzy_nutrition_water_volume_medium"`
	FuzzyNutritionWaterVolumeHigh   float32 `db:"fuzzy_nutrition_water_volume_high"`
}
