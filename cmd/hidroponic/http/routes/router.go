package routes

import (
	"hidroponic/cmd/hidroponic/dependencies"
	"hidroponic/cmd/hidroponic/http/handler/authentication"
	"hidroponic/cmd/hidroponic/http/handler/installationconfig"
	"hidroponic/cmd/hidroponic/http/handler/nutritionwaterlevel"
	"hidroponic/cmd/hidroponic/http/handler/plant"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/cmd/hidroponic/http/middleware"
	"hidroponic/internal/errors"
	"hidroponic/internal/platform/httpserver"
	"hidroponic/internal/platform/websocket"
	"net/http"

	"github.com/gorilla/mux"
)

type apiRouter struct {
	dep                           *dependencies.Dependency
	baseRoute                     *mux.Router
	wss                           *websocket.WebSocketService
	authController                *authentication.AuthHandlers
	nutritionWaterLevelController *nutritionwaterlevel.NutritionWaterLevelHandlers
	plantController               *plant.PlantHandlers
	instalatinoConfigController   *installationconfig.InstallationConfigHandlers
}

func Init(httpService *httpserver.HttpService, wss *websocket.WebSocketService, dep *dependencies.Dependency) {
	router := &apiRouter{
		baseRoute:                     httpService.GetRoute(),
		wss:                           wss,
		dep:                           dep,
		authController:                authentication.New(dep.UserUsecase, dep.AuthToken),
		nutritionWaterLevelController: nutritionwaterlevel.New(dep.NutritionWaterLevelUsecase),
		plantController:               plant.New(dep.PlantUsecase),
		instalatinoConfigController:   installationconfig.New(dep.InstallationConfigUsecase),
	}

	router.baseRoute.HandleFunc("/ping", router.ping).Methods(http.MethodGet)
	router.baseRoute.NotFoundHandler = http.HandlerFunc(router.notFound)

	router.initApiv1()

}

func (ar apiRouter) initApiv1() {
	loginRoute := ar.baseRoute.Path("/api/v1/auth/login").Subrouter()
	loginRoute.HandleFunc("", ar.authController.Login).Methods(http.MethodPost)

	ar.baseRoute.HandleFunc("/ws", ar.wss.ServeHTTP(func(messageData []byte) *string { return nil }))

	apiv1Route := ar.baseRoute.PathPrefix("/api/v1").Subrouter()

	apiv1Route.Use(middleware.Auth(ar.dep.AuthToken))
	apiv1Route.HandleFunc("/nutrition-water-levels", ar.nutritionWaterLevelController.GetActivePlantNutritionWaterLevel).Methods(http.MethodGet)

	apiv1Route.HandleFunc("/plants", ar.plantController.InsertPlant).Methods(http.MethodPost)
	apiv1Route.HandleFunc("/plants", ar.plantController.UpdatePlant).Methods(http.MethodPut)
	apiv1Route.HandleFunc("/plants/status", ar.plantController.UpdatePlantStatus).Methods(http.MethodPut)
	apiv1Route.HandleFunc("/plants", ar.plantController.GetAllPlant).Methods(http.MethodGet)
	apiv1Route.HandleFunc("/plants/active", ar.plantController.GetActivePlant).Methods(http.MethodGet)
	apiv1Route.HandleFunc("/plants/{id}", ar.plantController.GetPlantByID).Methods(http.MethodGet)

	apiv1Route.HandleFunc("/installation-configs", ar.instalatinoConfigController.GetInstallationConfig).Methods(http.MethodGet)
	apiv1Route.HandleFunc("/installation-configs", ar.instalatinoConfigController.UpdatePlant).Methods(http.MethodPut)
}

func (apiRouter) notFound(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[map[string]string]{
		Error: errors.ErrorPathNotFound.New("path not found"),
	}

	resp.Send(w)
}

func (apiRouter) ping(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[map[string]string]{
		IsSuccess: true,
		Data:      map[string]string{"ping": "Pong!!!"},
	}

	resp.Send(w)
}
