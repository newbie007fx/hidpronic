package ports

import (
	"context"
	"hidroponic/internal/module/automation/models"
)

type Usecase interface {
	InitiateAutomation(ctx context.Context) (resp models.InitAutomationResponse, er error)
	CompleteAutomation(ctx context.Context, automationID uint, data models.CompleteAutomation) error
}
