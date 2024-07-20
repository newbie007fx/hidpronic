package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ph PlantHandlers) GetAllPlant(rw http.ResponseWriter, req *http.Request) {
	plantData, err := ph.plantUsecase.GetAllPlant(req.Context())
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[[]models.BasicPlant]{
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

func (ph PlantHandlers) GetPlantByID(rw http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, errPars := strconv.Atoi(idString)
	if errPars != nil {
		errPars = errors.ErrorPathNotFound.New("data not found")
		return
	}
	plantData, err := ph.plantUsecase.GetPlantByID(req.Context(), uint(id))
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
