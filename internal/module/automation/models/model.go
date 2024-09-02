package models

import (
	"hidroponic/internal/module/automation/types"
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
	FinishedAt  *time.Time   `json:"finished_at"`
	Plant       SimplePlant  `json:"plant"`
}

type AutomationBasic struct {
	ID          uint         `json:"id"`
	PlantID     uint         `json:"plant_id"`
	TargetPPM   float32      `json:"target_ppm"`
	Accuration  float32      `json:"accuration"`
	Duration    int          `json:"duration"`
	Status      types.Status `json:"status"`
	TriggeredAt time.Time    `json:"triggered_at"`
	FinishedAt  *time.Time   `json:"finished_at"`
	Plant       SimplePlant  `json:"plant"`
}

type SimplePlant struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Varieties string `json:"varieties"`
}

type InitAutomationResponse struct {
	AutomationID          uint
	TargetNutritionVolume float32
	TargetWaterDistance   float32
	TargetWaterVolume     float32
}

type CalculateNutritionNeeded struct {
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
