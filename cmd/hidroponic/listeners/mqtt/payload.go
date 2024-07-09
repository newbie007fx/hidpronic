package mqtt

import "hidroponic/internal/models/sensor"

type ReadSensorPayload struct {
	Sensor sensor.SensorName `json:"sensor"`
	Value  float32           `json:"value"`
}

type StateClient struct {
	Action string `json:"action"`
	Value  *uint  `json:"value,omitempty"`
}
