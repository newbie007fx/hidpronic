package authentication

import (
	"hidroponic/cmd/hidroponic/http/helpers/authentication"
	"hidroponic/internal/module/users/ports"
)

type AuthHandlers struct {
	authToken   *authentication.TokenAuth
	userUsecase ports.Usecase
}

func New(userUsecase ports.Usecase, authToken *authentication.TokenAuth) *AuthHandlers {
	return &AuthHandlers{
		userUsecase: userUsecase,
		authToken:   authToken,
	}
}
