package nutritionwaterlevel

import (
	"hidroponic/internal/module/nutritionwaterlevel/ports"
)

type NutritionWaterLevelHandlers struct {
	nutritionWaterLevelUsecase ports.Usecase
}

func New(nutritionWaterLevelUsecase ports.Usecase) *NutritionWaterLevelHandlers {
	return &NutritionWaterLevelHandlers{
		nutritionWaterLevelUsecase: nutritionWaterLevelUsecase,
	}
}
