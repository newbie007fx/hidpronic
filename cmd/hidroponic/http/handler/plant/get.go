package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/plants/models"
	"net/http"
)

func (ph PlantHandlers) GetAllPlant(rw http.ResponseWriter, req *http.Request) {
	plantData, err := ph.plantUsecase.GetAllPlant(req.Context())
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[[]models.Plant]{
		IsSuccess: true,
		Data:      plantData,
	}

	resp.Send(rw)
}

func (ph PlantHandlers) GetActivePlant(rw http.ResponseWriter, req *http.Request) {
	plantData, err := ph.plantUsecase.GetActivePlant(req.Context())
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[*models.Plant]{
		IsSuccess: true,
		Data:      plantData,
	}

	resp.Send(rw)
}
