package service

import (
	"fmt"
	"os"

	"github.com/brandonrubio/twauter/service/api"
	"github.com/brandonrubio/twauter/service/config"
	"github.com/brandonrubio/twauter/service/env"
	"github.com/brandonrubio/twauter/service/logger"
	"go.uber.org/dig"
)

type IServiceCatalog interface {
	InitServices()
	Run()
}

type ServiceCatalog struct {
	envService      env.IEnvService
	appConfgService config.IAppConfigService
	loggerService   logger.ILoggerService
	apiService      api.IApiService
}

func (sc *ServiceCatalog) InitServices() {
	defaults := sc.envService.Init()

	err := sc.appConfgService.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = sc.loggerService.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for envName, envValue := range defaults {
		sc.loggerService.Log("warn", fmt.Sprintf("environment variable %s is not set, using default value %s", envName, envValue), "ServiceCatalog.InitServices")
	}

	err = sc.apiService.Init()
	if err != nil {
		panic("failed to init api")
	}
}

func (sc *ServiceCatalog) Run() {
	sc.loggerService.Log("info", "starting twautter...", "ServiceCatalog.Run")
	sc.apiService.Start()
}

type ServiceCatalogDependencies struct {
	dig.In
	EnvService       env.IEnvService          `name:"EnvService"`
	AppConfigService config.IAppConfigService `name:"AppConfigService"`
	LoggerService    logger.ILoggerService    `name:"LoggerService"`
	ApiService       api.IApiService          `name:"ApiService"`
}

func CreateServiceCatalog(dependencies ServiceCatalogDependencies) *ServiceCatalog {
	return &ServiceCatalog{
		envService:      dependencies.EnvService,
		appConfgService: dependencies.AppConfigService,
		loggerService:   dependencies.LoggerService,
		apiService:      dependencies.ApiService,
	}
}
