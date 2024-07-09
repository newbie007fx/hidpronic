package entities

type ContainerConfig struct {
	ID         string  `db:"id"`
	Name       string  `db:"name"`
	SensorGap  float32 `db:"sensor_gap"`
	Height     float32 `db:"height"`
	BottomArea float32 `db:"bottom_area"`
	TopArea    float32 `db:"top_area"`
	Volume     float32 `db:"volume"`
}
