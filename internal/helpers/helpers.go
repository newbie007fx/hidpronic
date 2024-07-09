package helpers

import (
	"hidroponic/internal/constants"
	"hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/platform/mqtt"
)

func InitHelpers(mqttService *mqtt.MqttService) {
	state := constants.StateOff
	plantID := helpers.GetActivePlantIDInstance().Get()
	if plantID > 0 {
		state = constants.StateOn
	}
	deviceStateInstance = &deviceState{
		mqttService: mqttService,
		state:       state,
	}
}
