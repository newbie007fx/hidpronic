package devicestate

import (
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/helpers"
	"hidroponic/internal/types"
	"net/http"
)

func (ah DeviceStateHnalders) GetDeviceState(rw http.ResponseWriter, req *http.Request) {
	stateHelper := helpers.GetDeviceStateInstance()
	resp := response.Response[any]{
		IsSuccess: true,
		Data: map[string]types.DeviceState{
			"state": stateHelper.GetState(),
		},
	}

	resp.Send(rw)
}
