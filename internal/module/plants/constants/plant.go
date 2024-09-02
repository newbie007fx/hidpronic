package constants

import "hidroponic/internal/module/plants/types"

const (
	StatusCreated     types.Status = "CREATED"
	StatusActivated   types.Status = "ACTIVATED"
	StatusHarvested   types.Status = "HARVESTED"
	StatusDeactivated types.Status = "DEACTIVATED"

	TypeLeafCrop  types.PlantType = "LEAF_CROP"
	TypeFruitCrop types.PlantType = "FRUIT_CROP"

	GrowthVegetative types.Growth = "VEGETATIVE"
	GrowthGenerative types.Growth = "GENERATIVE"
)

func PlantTypeMap() map[types.PlantType]string {
	return map[types.PlantType]string{
		TypeFruitCrop: "Fruit Crop",
		TypeLeafCrop:  "Leaf Crop",
	}
}

func GrowthTypeMap() map[types.Growth]string {
	return map[types.Growth]string{
		GrowthGenerative: "Generative",
		GrowthVegetative: "Vegetative",
	}
}
