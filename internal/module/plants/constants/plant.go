package constants

import "hidroponic/internal/module/plants/types"

const (
	StatusCreated   types.Status = "CREATED"
	StatusActived   types.Status = "ACTIVED"
	StatusHarvested types.Status = "HARVESTED"
	StatusDeactived types.Status = "DEACTIVED"

	TypeLeafCrop  types.PlantType = "LEAF_CROP"
	TypeFruitCrop types.PlantType = "FRUIT_CROP"

	GrowthVegetative types.Growth = "VEGETATIVE"
	GrowthGenerative types.Growth = "GENERATIVE"
)
