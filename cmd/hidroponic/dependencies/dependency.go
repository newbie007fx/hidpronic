package dependencies

import (
	"hidroponic/cmd/hidroponic/http/helpers/authentication"
	automationPorts "hidroponic/internal/module/automation/ports"
	automationRepository "hidroponic/internal/module/automation/repositories"
	automationUsecase "hidroponic/internal/module/automation/usecase"
	installationConfigPorts "hidroponic/internal/module/installationconfig/ports"
	installationConfigRepository "hidroponic/internal/module/installationconfig/repositories"
	installationConfigUsecase "hidroponic/internal/module/installationconfig/usecase"
	nutritionWaterLevelPorts "hidroponic/internal/module/nutritionwaterlevel/ports"
	nutritionWaterLevelRepository "hidroponic/internal/module/nutritionwaterlevel/repositories"
	nutritionWaterLevelUsecase "hidroponic/internal/module/nutritionwaterlevel/usecase"
	plantPorts "hidroponic/internal/module/plants/ports"
	plantRepository "hidroponic/internal/module/plants/repositories"
	plantUsecase "hidroponic/internal/module/plants/usecase"
	userPorts "hidroponic/internal/module/users/ports"
	userRepository "hidroponic/internal/module/users/repositories"
	userUsecase "hidroponic/internal/module/users/usecase"
	"hidroponic/internal/platform/configuration"
	"hidroponic/internal/platform/database"
)

type Dependency struct {
	DatabaseService *database.DatabaseService
	ConfigService   *configuration.ConfigService

	AutomationUsecase          automationPorts.Usecase
	InstallationConfigUsecase  installationConfigPorts.Usecase
	PlantUsecase               plantPorts.Usecase
	UserUsecase                userPorts.Usecase
	NutritionWaterLevelUsecase nutritionWaterLevelPorts.Usecase

	AuthToken *authentication.TokenAuth
}

func New(db *database.DatabaseService, cs *configuration.ConfigService) *Dependency {
	return &Dependency{
		DatabaseService: db,
		ConfigService:   cs,
	}
}

func (dp *Dependency) Init() error {
	conf := dp.ConfigService.GetConfig()

	installationConfigRepo := installationConfigRepository.New(dp.DatabaseService)
	dp.InstallationConfigUsecase = installationConfigUsecase.New(installationConfigRepo)

	plantRepo := plantRepository.New(dp.DatabaseService)
	dp.PlantUsecase = plantUsecase.New(plantRepo)

	userRepo := userRepository.New(dp.DatabaseService)
	dp.UserUsecase = userUsecase.New(userRepo)

	nutritionWaterLevelRepo := nutritionWaterLevelRepository.New(dp.DatabaseService)
	dp.NutritionWaterLevelUsecase = nutritionWaterLevelUsecase.New(nutritionWaterLevelRepo, plantRepo)

	automationRepo := automationRepository.New(dp.DatabaseService)
	dp.AutomationUsecase = automationUsecase.New(automationRepo, installationConfigRepo, plantRepo)

	dp.AuthToken = authentication.New(conf.JWT.Secret)

	return nil
}
