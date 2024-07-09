package sensor

type SensorName string
type DataType string
type ActionType string

const ActionReadValue ActionType = "read_value"
const ActionNewValue ActionType = "new_value"

const SensorWaterPH SensorName = "water_ph"
const SensorNutitionWaterLevel SensorName = "nutrition_water_level"
const SensorNutritionWaterDistance SensorName = "nutrition_water_distance"
const SensorRawWaterDistance SensorName = "raw_water_distance"
const SensorWaterTemperature SensorName = "water_temperature"

const TypeWaterPH DataType = "water_ph"
const TypeNutritionWaterLevel DataType = "nutrition_water_level"
const TypeNutritionWaterVolume DataType = "nutrition_water_volume"
const TypeRawWaterVolume DataType = "raw_water_volume"
const TypeWaterTemperature DataType = "water_temperature"

type DataValueWs struct {
	DataType   DataType   `json:"data_type"`
	Value      float32    `json:"value"`
	ActionType ActionType `json:"action_type"`
	CreatedAt  int64      `json:"created_at"`
}

var MapSensorToDataType = map[SensorName]DataType{
	SensorWaterPH:                TypeWaterPH,
	SensorNutitionWaterLevel:     TypeNutritionWaterLevel,
	SensorNutritionWaterDistance: TypeNutritionWaterVolume,
	SensorRawWaterDistance:       TypeRawWaterVolume,
	SensorWaterTemperature:       TypeWaterTemperature,
}
