package helpers

import (
	"hidroponic/internal/constants"
	"hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/platform/mqtt"
	"hidroponic/internal/platform/websocket"
)

func InitHelpers(mqttService *mqtt.MqttService, wsService *websocket.WebSocketService) {
	state := constants.StateOff
	plantID := helpers.GetActivePlantIDInstance().Get()
	if plantID > 0 {
		state = constants.StateOn
	}
	deviceStateInstance = &deviceState{
		mqttService: mqttService,
		wsService:   wsService,
		state:       state,
	}
}
