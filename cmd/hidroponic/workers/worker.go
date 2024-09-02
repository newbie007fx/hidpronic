package workers

import (
	"context"
	automationPorts "hidroponic/internal/module/automation/ports"
	nutritionWaterLevelPorts "hidroponic/internal/module/nutritionwaterlevel/ports"
	plantPorts "hidroponic/internal/module/plants/ports"
	"hidroponic/internal/platform/configuration"
	"hidroponic/internal/platform/websocket"
	"time"
)

type Worker struct {
	configService     *configuration.ConfigService
	plantUsecase      plantPorts.Usecase
	waterLevelUsecase nutritionWaterLevelPorts.Usecase
	automationUsecase automationPorts.Usecase

	websocketService *websocket.WebSocketService
}

func New(plantUsecase plantPorts.Usecase, waterLevelUsecase nutritionWaterLevelPorts.Usecase, automationUsecase automationPorts.Usecase, websocketService *websocket.WebSocketService, configService *configuration.ConfigService) *Worker {
	return &Worker{
		automationUsecase: automationUsecase,
		configService:     configService,
		plantUsecase:      plantUsecase,
		waterLevelUsecase: waterLevelUsecase,

		websocketService: websocketService,
	}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			currentTime := time.Now()

			w.StoreNutritionWaterLevelFromTemp(currentTime)
			w.AutomationTrigger(currentTime)

			sleep := 60 - time.Now().Second()
			time.Sleep(time.Duration(sleep) * time.Second)
		}

	}
}
