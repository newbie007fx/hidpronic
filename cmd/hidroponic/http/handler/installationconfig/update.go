package installationconfig

import (
	"hidroponic/cmd/hidroponic/http/helpers/request"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/installationconfig/models"
	"net/http"
)

func (ic InstallationConfigHandlers) UpdatePlant(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.UpdateInstallationConfig
	baseErr := request.ReadRequestBody(req, &requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	baseErr = ic.instalaltionConfigUsecase.UpdateInstallationConfig(req.Context(), requestBody)
	if baseErr != nil {
		response.WriterResponseError(rw, baseErr)
		return
	}

	resp := response.Response[any]{
		IsSuccess: true,
	}

	resp.Send(rw)
}
