package mqtt

import (
	automationPorts "hidroponic/internal/module/automation/ports"
	installationPorts "hidroponic/internal/module/installationconfig/ports"
	waterNutritionPorts "hidroponic/internal/module/nutritionwaterlevel/ports"
	"hidroponic/internal/platform/mqtt"
	"hidroponic/internal/platform/websocket"
)

type MqttListener struct {
	ws          *websocket.WebSocketService
	mqttService *mqtt.MqttService

	automationUsecase     automationPorts.Usecase
	installationUsecase   installationPorts.Usecase
	waterNutritionUsecase waterNutritionPorts.Usecase
}

func New(mqttService *mqtt.MqttService, ws *websocket.WebSocketService, installationUsecase installationPorts.Usecase, waterNutritionUsecase waterNutritionPorts.Usecase, automationUsecase automationPorts.Usecase) *MqttListener {
	return &MqttListener{
		mqttService:           mqttService,
		ws:                    ws,
		automationUsecase:     automationUsecase,
		installationUsecase:   installationUsecase,
		waterNutritionUsecase: waterNutritionUsecase,
	}
}

func (ml MqttListener) Run() error {
	ml.readSensorValue()
	ml.listenClientState()

	return nil
}
