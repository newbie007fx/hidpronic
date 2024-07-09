package usecase

import (
	"context"
	"database/sql"
	"hidroponic/internal/module/automation/constants"
	"hidroponic/internal/module/automation/entities"
	"hidroponic/internal/module/automation/models"
)

func (u Usecase) CompleteAutomation(ctx context.Context, automationID uint, data models.CompleteAutomation) error {
	automation, err := u.repo.GetAutomationByID(ctx, automationID)
	if err != nil {
		return err.ToError()
	}

	diff := data.FinishedAt.Sub(automation.TriggeredAt)
	automation.Duration = int(diff.Seconds())
	automation.Accuration = calculateAccuration(automation.TargetPPM, data.AfterData.NutritionWaterPPM)
	automation.AfterData = entities.SensorData{SensorData: data.AfterData}
	automation.FinishedAt = sql.NullTime{
		Valid: true,
		Time:  data.FinishedAt,
	}
	automation.Status = constants.StatusComplete

	err = u.repo.UpdateAutomation(ctx, automation)

	return err
}

func calculateAccuration(target, result float32) float32 {
	diff := target - result
	if diff < 0 {
		diff = diff * -1
	}

	return (1 - (diff / target)) * 100
}
