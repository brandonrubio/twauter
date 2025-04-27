package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/brandonrubio/twauter/service/api/handler"
	"github.com/brandonrubio/twauter/service/config"
	"github.com/brandonrubio/twauter/service/errorhandler"
	"github.com/brandonrubio/twauter/service/logger"
	"go.uber.org/dig"
)

const version = "v0"

type IApiService interface {
	Init() error
	Start()
}

type ApiService struct {
	router              *http.ServeMux
	appConfigService    config.IAppConfigService
	loggerService       logger.ILoggerService
	errorHandlerService errorhandler.IErrorHandlerService
	twautHandlerService handler.ITwautHandlerService
}

func (as *ApiService) Init() error {
	as.router = http.NewServeMux()
	base := fmt.Sprintf("/api/%s", version)

	as.router.HandleFunc(fmt.Sprintf("GET %s/twaut", base), as.twautHandlerService.HandleGetAll)
	as.loggerService.Log("info", fmt.Sprintf("registered route GET %s/twaut", base), "ApiService.Init")

	as.router.HandleFunc(fmt.Sprintf("GET %s/twaut/{id}", base), as.twautHandlerService.HandleGetOne)
	as.loggerService.Log("info", fmt.Sprintf("registered route GET %s/twaut/{id}", base), "ApiService.Init")

	as.router.HandleFunc(fmt.Sprintf("POST %s/twaut", base), as.twautHandlerService.HandlePost)
	as.loggerService.Log("info", fmt.Sprintf("registered route POST %s/twaut", base), "ApiService.Init")

	as.router.HandleFunc(fmt.Sprintf("PUT %s/twaut/{id}", base), as.twautHandlerService.HandlePut)
	as.loggerService.Log("info", fmt.Sprintf("registered route PUT %s/twaut/{id}", base), "ApiService.Init")

	as.router.HandleFunc(fmt.Sprintf("DELETE %s/twaut/{id}", base), as.twautHandlerService.HandleDelete)
	as.loggerService.Log("info", fmt.Sprintf("registered route DELETE %s/twaut/{id}", base), "ApiService.Init")

	return nil
}

func (as *ApiService) Start() {
	apiConfig := as.appConfigService.GetApiConfig()
	as.loggerService.Log("info", fmt.Sprintf("listening on port %s", apiConfig.Port), "ApiService.Start")
	err := http.ListenAndServe(fmt.Sprintf(":%s", apiConfig.Port), as.router)
	if err != nil {
		as.errorHandlerService.HandleError(err, "failed to start http server")
		os.Exit(1)
	}
}

type ApiServiceDependencies struct {
	dig.In
	AppConfigService    config.IAppConfigService          `name:"AppConfigService"`
	LoggerService       logger.ILoggerService             `name:"LoggerService"`
	ErrorHandlerService errorhandler.IErrorHandlerService `name:"ErrorHandlerService"`
	TwautHandlerService handler.ITwautHandlerService      `name:"TwautHandlerService"`
}

func CreateApiService(dependencies ApiServiceDependencies) *ApiService {
	return &ApiService{
		appConfigService:    dependencies.AppConfigService,
		loggerService:       dependencies.LoggerService,
		errorHandlerService: dependencies.ErrorHandlerService,
		twautHandlerService: dependencies.TwautHandlerService,
	}
}
