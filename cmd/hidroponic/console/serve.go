package console

import (
	"context"
	"hidroponic/cmd/hidroponic/dependencies"
	"hidroponic/cmd/hidroponic/http/routes"
	mqttListener "hidroponic/cmd/hidroponic/listeners/mqtt"
	"hidroponic/cmd/hidroponic/workers"
	"hidroponic/internal/helpers"
	automationHelper "hidroponic/internal/module/automation/helpers"
	plantHelper "hidroponic/internal/module/plants/helpers"
	"hidroponic/internal/platform/configuration"
	"hidroponic/internal/platform/database"
	"hidroponic/internal/platform/httpserver"
	"hidroponic/internal/platform/mqtt"
	"hidroponic/internal/platform/validation"
	"hidroponic/internal/platform/websocket"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func (cs Console) initServeCommand() {
	cmd := cs.ConsoleService.GetCommandInstance()

	cmd.Use = "serve"

	cmd.Short = "Run service"

	cmd.Run = func(_ *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		config := configuration.New(".", "config", "yaml")
		config.Setup()

		logLevel := &slog.LevelVar{}
		logLevel.UnmarshalText([]byte(config.GetConfig().App.LogLevel))
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
		slog.SetDefault(logger)

		wss := websocket.New()
		wss.Setup()

		mqttService := mqtt.New(config)
		mqttService.Setup()

		db := database.New(config)
		db.Setup()

		validator := validation.New()
		validator.Setup()

		server := httpserver.New(config)
		server.Setup()

		dep := dependencies.New(db, config)
		dep.Init()

		plantHelper.InitHelpers(dep.PlantUsecase)
		helpers.InitHelpers(mqttService, wss)
		automationHelper.InitHelpers()

		listenerMqtt := mqttListener.New(mqttService, wss, dep.InstallationConfigUsecase, dep.NutritionWaterLevelUsecase, dep.AutomationUsecase)
		listenerMqtt.Run()

		worker := workers.New(dep.PlantUsecase, dep.NutritionWaterLevelUsecase, dep.AutomationUsecase, wss, config)
		go worker.Run(ctx)

		routes.Init(server, wss, dep)

		server.Start()

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-done
		log.Println("shutting down")

		server.Shutdown(ctx)
		db.Shutdown()
		mqttService.Shutdown()

		log.Println("all server stopped!")
	}

	cs.ConsoleService.RegisterCommand(cmd)
}
