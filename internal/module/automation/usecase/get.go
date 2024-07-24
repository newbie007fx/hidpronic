package usecase

import (
	"context"
	"hidroponic/internal/errors"
	"hidroponic/internal/module/automation/models"
)

func (u Usecase) GetAllAutomation(ctx context.Context, limit, offset int, filter map[string]string) (resp []models.AutomationBasic, err *errors.BaseError) {
	result, err := u.repo.GetAllAutomation(ctx, limit, offset, filter)
	if err != nil {
		return
	}

	resp = []models.AutomationBasic{}
	for _, res := range result {
		basicPlant := models.AutomationBasic{
			ID:          res.ID,
			PlantID:     res.PlantID,
			TargetPPM:   res.TargetPPM,
			Accuration:  res.Accuration,
			Duration:    res.Duration,
			Status:      res.Status,
			TriggeredAt: res.TriggeredAt,
			FinishedAt:  res.FinishedAt.Time,
			Plant: models.SimplePlant{
				ID:        res.PlantID,
				Name:      res.Plant.Name,
				Varieties: res.Plant.Varieties,
			},
		}
		resp = append(resp, basicPlant)
	}

	return
}

func (u Usecase) GetAutomationByID(ctx context.Context, id uint) (resp models.Automation, err *errors.BaseError) {
	result, err := u.repo.GetAutomationByID(ctx, id)
	if err != nil {
		return
	}

	resp = models.Automation{
		ID:          result.ID,
		PlantID:     result.PlantID,
		BofereData:  result.BofereData.SensorData,
		AfterData:   result.AfterData.SensorData,
		Accuration:  result.Accuration,
		TargetPPM:   result.TargetPPM,
		Duration:    result.Duration,
		Status:      result.Status,
		TriggeredAt: result.TriggeredAt,
		FinishedAt:  result.FinishedAt.Time,
		Plant: models.SimplePlant{
			ID:        result.PlantID,
			Name:      result.Plant.Name,
			Varieties: result.Plant.Varieties,
		},
	}

	return
}
