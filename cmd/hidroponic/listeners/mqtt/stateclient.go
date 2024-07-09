package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"hidroponic/internal/helpers"
	automationHelpers "hidroponic/internal/module/automation/helpers"
	"hidroponic/internal/module/automation/models"
	"hidroponic/internal/platform/mqtt"
	"log/slog"
	"time"
)

func (ml MqttListener) listenClientState() {
	const topic = "device_state_client"

	var handler mqtt.CallbackFunc = func(payloadRaw []byte) {
		var payloadData StateClient
		err := json.Unmarshal(payloadRaw, &payloadData)
		if err != nil {
			slog.Warn("cannot read payload data")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()

		switch payloadData.Action {
		case "get":
			helpers.GetDeviceStateInstance().PublishState(nil)
		case "complete":
			lastData := automationHelpers.GetLastSensorDataInstance().Get()
			err := ml.automationUsecase.CompleteAutomation(ctx, *payloadData.Value, models.CompleteAutomation{
				AfterData:  *lastData,
				FinishedAt: time.Now(),
			})
			if err != nil {
				slog.Error(fmt.Sprintf("error complete automation with message: %s", err.Error()))
			}
		default:
			slog.Error(fmt.Sprintf("error, unidentified action: %s", payloadData.Action))
		}

	}

	ml.mqttService.Subscribe(topic, handler)
}
