package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/users/entities"
)

type Repository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, *errors.BaseError)
}
