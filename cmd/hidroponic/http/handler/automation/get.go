package automation

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/automation/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ah AutomationHandlers) GetAllAutomation(rw http.ResponseWriter, req *http.Request) {
	filter := extractFilter(req)
	limit, offset := extractLimitOffset(req)
	automationData, err := ah.automationUsecase.GetAllAutomation(req.Context(), limit, offset, filter)
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[[]models.AutomationBasic]{
		IsSuccess: true,
		Data:      automationData,
	}

	resp.Send(rw)
}

func (ah AutomationHandlers) GetAutomationByID(rw http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, errPars := strconv.Atoi(idString)
	if errPars != nil {
		errPars = errors.ErrorPathNotFound.New("data not found")
		return
	}
	automationData, err := ah.automationUsecase.GetAutomationByID(req.Context(), uint(id))
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[models.Automation]{
		IsSuccess: true,
		Data:      automationData,
	}

	resp.Send(rw)
}

func extractLimitOffset(req *http.Request) (limit, offset int) {
	limit = 20

	stringLimit := req.URL.Query().Get("limit")
	if stringLimit != "" {
		tmpLimit, _ := strconv.Atoi(stringLimit)
		if tmpLimit > 0 {
			limit = tmpLimit
		}
	}

	stringOffset := req.URL.Query().Get("offset")
	if stringOffset != "" {
		tmpOffset, _ := strconv.Atoi(stringOffset)
		if tmpOffset > 0 {
			offset = tmpOffset
		}
	}

	return
}

func extractFilter(req *http.Request) map[string]string {
	allowedFilters := []string{"plant_id", "status"}

	filterMap := map[string]string{}
	for _, filterKey := range allowedFilters {
		if req.URL.Query().Has(filterKey) {
			filterMap[filterKey] = req.URL.Query().Get(filterKey)
		}
	}

	return filterMap
}
