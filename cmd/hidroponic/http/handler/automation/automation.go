package automation

import (
	"hidroponic/internal/module/automation/ports"
)

type AutomationHandlers struct {
	automationUsecase ports.Usecase
}

func New(automationUsecase ports.Usecase) *AutomationHandlers {
	return &AutomationHandlers{
		automationUsecase: automationUsecase,
	}
}
