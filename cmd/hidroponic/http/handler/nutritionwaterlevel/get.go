package nutritionwaterlevel

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/nutritionwaterlevel/models"
	"net/http"
)

func (wnl NutritionWaterLevelHandlers) GetActivePlantNutritionWaterLevel(rw http.ResponseWriter, req *http.Request) {
	result, err := wnl.nutritionWaterLevelUsecase.GetActivePlantNutritionWaterLevel(req.Context())
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[[]models.NutritionWaterLevel]{
		IsSuccess: true,
		Data:      result,
	}

	resp.Send(rw)
}
