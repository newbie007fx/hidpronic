package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/request"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/plants/models"
	"net/http"
)

func (ph PlantHandlers) InsertPlant(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.CreatePlant
	err := request.ReadRequestBody(req, &requestBody)
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	plantData, err := ph.plantUsecase.InsertPlant(req.Context(), requestBody)
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
