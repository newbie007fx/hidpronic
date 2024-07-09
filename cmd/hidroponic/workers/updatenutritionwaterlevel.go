package workers

import (
	"context"
	"fmt"
	"hidroponic/internal/models/sensor"
	"hidroponic/internal/module/plants/helpers"
	"log/slog"
	"time"
)

func (w *Worker) StoreNutritionWaterLevelFromTemp(currentTime time.Time) {
	if currentTime.Minute()%10 == 0 && currentTime.Second() < 30 {
		plantID := helpers.GetActivePlantIDInstance().Get()
		if plantID == 0 {
			slog.Warn("cannot store nutrition water level: no active plant")
			return
		}

		startDate := currentTime.Add(-10 * time.Minute)

		resp, err := w.waterLevelUsecase.InsertNutritionWaterLevelFromTempRange(context.Background(), plantID, startDate, currentTime)
		if err != nil {
			slog.Warn(fmt.Sprintf("cannot insert with message: %s", err.Error()))
			return
		}

		dataValue := sensor.DataValueWs{
			DataType:   sensor.TypeNutritionWaterLevel,
			Value:      resp.Value,
			ActionType: sensor.ActionNewValue,
			CreatedAt:  resp.CreatedAt,
		}

		w.websocketService.Broadcast(dataValue.DataType, dataValue)
	}
}
