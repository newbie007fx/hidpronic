package mqtt

import (
	"context"
	"encoding/json"
	"hidroponic/internal/constants"
	"hidroponic/internal/errors"
	"hidroponic/internal/helpers"
	"hidroponic/internal/models/sensor"
	automationHelpers "hidroponic/internal/module/automation/helpers"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	plantHelpers "hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/platform/mqtt"
	"log/slog"
	"math"
	"time"
)

func (ml MqttListener) readSensorValue() {
	const topic = "read_sensor"

	var handler mqtt.CallbackFunc = func(payloadRaw []byte) {
		var payloadData []ReadSensorPayload
		err := json.Unmarshal(payloadRaw, &payloadData)
		if err != nil {
			slog.Warn("cannot read payload data")
			return
		}
		plantID := plantHelpers.GetActivePlantIDInstance().Get()
		if plantID == 0 {
			helpers.GetDeviceStateInstance().SetState(constants.StateOff)

			return
		}

		ml.processPayload(payloadData, plantID)
	}

	ml.mqttService.Subscribe(topic, handler)
}

func (ml MqttListener) processPayload(payload []ReadSensorPayload, plantID uint) {
	for _, readSensor := range payload {
		wsPayload := sensor.DataValueWs{
			DataType:   sensor.MapSensorToDataType[readSensor.Sensor],
			ActionType: sensor.ActionReadValue,
			Value:      float32(math.Round(float64(readSensor.Value)*100) / 100),
			CreatedAt:  time.Now().UnixMilli(),
		}

		if readSensor.Sensor == sensor.SensorRawWaterDistance || readSensor.Sensor == sensor.SensorNutritionWaterDistance {
			var err *errors.BaseError
			wsPayload.Value, err = ml.installationUsecase.CalculateContainerVolume(context.Background(), string(wsPayload.DataType), readSensor.Value)
			if err != nil {
				slog.Warn(err.Error())
			}
		}

		if readSensor.Sensor == sensor.SensorNutitionWaterLevel {
			req := models.CreateNutritionWaterLevel{
				Value:   readSensor.Value,
				PlantID: plantID,
			}

			ml.waterNutritionUsecase.InsertNutritionWaterLevelTemp(context.Background(), req)
		}

		ml.ws.Broadcast(wsPayload.DataType, wsPayload)
		automationHelpers.GetLastSensorDataInstance().Set(wsPayload.DataType, wsPayload.Value)
	}
}
