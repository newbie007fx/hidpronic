package installationconfig

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/module/installationconfig/models"
	"net/http"
)

func (wnl InstallationConfigHandlers) GetInstallationConfig(rw http.ResponseWriter, req *http.Request) {
	result, err := wnl.instalaltionConfigUsecase.GetInstallationConfig(req.Context())
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[models.InstallationConfigResponse]{
		IsSuccess: true,
		Data:      *result,
	}

	resp.Send(rw)
}
