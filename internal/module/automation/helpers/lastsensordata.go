package helpers

import (
	"hidroponic/internal/models/sensor"
	"hidroponic/internal/module/automation/models"
)

var lastSensorDataInstance *lastSensorData

type lastSensorData struct {
	lastData *models.SensorData
}

func (ap *lastSensorData) Get() *models.SensorData {
	data := ap.lastData

	ap.lastData = nil

	return data
}

func (ap *lastSensorData) Set(sensorType sensor.DataType, value float32) {
	if ap.lastData == nil {
		ap.lastData = &models.SensorData{}
	}

	switch sensorType {
	case sensor.TypeNutritionWaterLevel:
		ap.lastData.NutritionWaterPPM = value
	case sensor.TypeNutritionWaterVolume:
		ap.lastData.NutritionWaterVolume = value
	case sensor.TypeWaterTemperature:
		ap.lastData.WaterTemperature = value
	case sensor.TypeWaterPH:
		ap.lastData.WaterPH = value
	}
}

func GetLastSensorDataInstance() *lastSensorData {
	return lastSensorDataInstance
}
