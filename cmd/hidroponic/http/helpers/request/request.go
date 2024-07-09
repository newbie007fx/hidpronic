package request

import (
	"encoding/json"
	"hidroponic/internal/errors"
	"net/http"
)

func ReadRequestBody(req *http.Request, reqModel interface{ Validate() error }) *errors.BaseError {
	if err := json.NewDecoder(req.Body).Decode(reqModel); err != nil {
		return errors.ErrorInvalidRequestBody.New("Error parse request body")
	}

	err := reqModel.Validate()
	if err != nil {
		return errors.ErrorValidation.New(err.Error())
	}

	return nil
}
