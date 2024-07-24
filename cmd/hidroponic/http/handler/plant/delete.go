package plant

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ph PlantHandlers) DeletePlant(rw http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, errPars := strconv.Atoi(idString)
	if errPars != nil {
		errPars = errors.ErrorPathNotFound.New("data not found")
		return
	}
	err := ph.plantUsecase.DeletePlant(req.Context(), uint(id))
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}
