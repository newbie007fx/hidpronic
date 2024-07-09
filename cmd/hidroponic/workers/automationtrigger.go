package workers

import (
	"context"
	"fmt"
	"hidroponic/internal/constants"
	"hidroponic/internal/helpers"
	"log/slog"
	"time"
)

func (w *Worker) AutomationTrigger(currentTime time.Time) {
	state := helpers.GetDeviceStateInstance().GetState()
	if currentTime.Minute()%15 == 0 && currentTime.Second() < 30 && state == constants.StateOn {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()

		resp, err := w.automationUsecase.InitiateAutomation(ctx)
		if err != nil {
			slog.Error(fmt.Sprintf("cannot trigger automation with message: %s", err.Error()))
			return
		}

		helpers.GetDeviceStateInstance().SetState(constants.StateRunAutomation).PublishState(&map[string]interface{}{
			"automation_id":           resp.AutomationID,
			"target_nutrition_volume": resp.TargetNutritionVolume,
			"target_water_distance":   resp.TargetWaterDistance,
		})
	}
}
