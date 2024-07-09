package models

type ContainerConfig struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	SensorGap  float32 `json:"sensor_gap"`
	Height     float32 `json:"height"`
	BottomArea float32 `json:"bottom_area"`
	TopArea    float32 `json:"top_area"`
	Volume     float32 `json:"volume"`
}

type UpdateContainerConfig struct {
	ID         string  `validate:"required" json:"id"`
	SensorGap  float32 `validate:"required" json:"sensor_gap"`
	Height     float32 `validate:"required" json:"height"`
	BottomArea float32 `validate:"required" json:"bottom_area"`
	TopArea    float32 `validate:"required" json:"top_area"`
}
