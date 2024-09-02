package ports

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/automation/models"
)

type Usecase interface {
	InitiateAutomation(ctx context.Context) (resp models.InitAutomationResponse, er error)
	CompleteAutomation(ctx context.Context, automationID uint, data models.CompleteAutomation) error

	GetAllAutomation(ctx context.Context, limit, offset int, filter map[string]string) (resp []models.AutomationBasic, err *errors.BaseError)
	GetAutomationByID(ctx context.Context, id uint) (resp models.Automation, err *errors.BaseError)

	CalculateNutritionNeeded(data models.CalculateNutritionNeeded) (rawWaterVolume, nutritionVolume float32)
}
