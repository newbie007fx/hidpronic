package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/automation/entities"
)

type Repository interface {
	GetAllAutomationWithLimit(ctx context.Context, limit int) ([]entities.Automation, *errors.BaseError)
	GetAutomationByID(ctx context.Context, id uint) (*entities.Automation, *errors.BaseError)

	InsertAutomation(ctx context.Context, data *entities.Automation) *errors.BaseError
	UpdateAutomation(ctx context.Context, data *entities.Automation) *errors.BaseError
}
