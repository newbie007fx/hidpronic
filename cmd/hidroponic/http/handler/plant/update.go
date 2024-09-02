package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/request"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (ph PlantHandlers) HarvestPlant(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.HarvestPlant
	baseErr := request.ReadRequestBody(req, &requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	baseErr = ph.plantUsecase.HarvestPlant(req.Context(), requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}

func (ph PlantHandlers) UpdatePlantGrowth(rw http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, errPars := strconv.Atoi(idString)
	if errPars != nil {
		errPars = errors.ErrorPathNotFound.New("data not found")
		return
	}

	err := ph.plantUsecase.UpdatePlantGrowth(req.Context(), uint(id))
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}
