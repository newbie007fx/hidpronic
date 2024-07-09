package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/request"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/plants/models"
	"net/http"
)

func (ph PlantHandlers) UpdatePlant(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.UpdatePlant
	baseErr := request.ReadRequestBody(req, &requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	baseErr = ph.plantUsecase.UpdatePlant(req.Context(), requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}

func (ph PlantHandlers) UpdatePlantStatus(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.UpdatePlantStatus
	baseErr := request.ReadRequestBody(req, &requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	baseErr = ph.plantUsecase.UpdatePlantStatus(req.Context(), requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}
