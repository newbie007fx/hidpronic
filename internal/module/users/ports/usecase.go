package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/users/models"
)

type Usecase interface {
	VerifyUsernamePassword(ctx context.Context, username, password string) (*models.User, *errors.BaseError)
}
