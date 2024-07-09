package models

import (
	"hidroponic/internal/module/plants/types"
	"time"
)

type Automation struct {
	ID          uint         `json:"id"`
	PlantID     uint         `json:"plant_id"`
	BofereData  SensorData   `json:"before_data"`
	AfterData   SensorData   `json:"after_data"`
	Accuration  float32      `json:"accuration"`
	TargetPPM   float32      `json:"target_ppm"`
	Duration    int          `json:"duration"`
	Status      types.Status `json:"status"`
	TriggeredAt time.Time    `json:"triggered_at"`
	FinishedAt  time.Time    `json:"finished_at"`
}

type InitAutomationResponse struct {
	AutomationID          uint
	TargetNutritionVolume float32
	TargetWaterDistance   float32
	TargetWaterVolume     float32
}

type CalcualteNutritionNeeded struct {
	CurrentNutritionWaterVolume float32
	TargetNutritionWaterVolume  float32
	CurrentNutritionWaterPPM    float32
	TargetNutritionWaterPPM     float32
	RawWaterPPM                 float32
	NutritionPPM                float32
}

type SensorData struct {
	NutritionWaterPPM    float32 `json:"nutrition_water_ppm"`
	NutritionWaterVolume float32 `json:"nutrition_water_volume"`
	WaterTemperature     float32 `json:"water_temperature"`
	WaterPH              float32 `json:"water_ph"`
}

type CompleteAutomation struct {
	AfterData  SensorData `json:"after_data"`
	FinishedAt time.Time  `json:"finished_at"`
}
