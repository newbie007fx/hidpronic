package helpers

import (
	"hidroponic/internal/constants"
	"hidroponic/internal/platform/mqtt"
	"hidroponic/internal/platform/websocket"
	"hidroponic/internal/types"
)

var deviceStateInstance *deviceState

const topic string = "device_state"

type deviceState struct {
	mqttService *mqtt.MqttService
	wsService   *websocket.WebSocketService
	state       types.DeviceState
}

func (ap *deviceState) SetState(state types.DeviceState) *deviceState {
	if state == constants.StateComplete {
		state = constants.StateOn
	}

	ap.state = state
	ap.wsService.Broadcast("deviceState", map[string]types.DeviceState{"stete": state})
	return ap
}

func (ap deviceState) GetState() types.DeviceState {
	return ap.state
}

func (ap deviceState) PublishState(additionalData *map[string]interface{}) {
	ap.mqttService.Publish(topic, publishPayload{
		State: ap.state,
		Data:  additionalData,
	})
}

func GetDeviceStateInstance() *deviceState {
	return deviceStateInstance
}

type publishPayload struct {
	State types.DeviceState       `json:"state"`
	Data  *map[string]interface{} `json:"data,omitempty"`
}
