package service

import (
	"github.com/brandonrubio/twauter/service/api"
	"github.com/brandonrubio/twauter/service/api/handler"
	"github.com/brandonrubio/twauter/service/config"
	"github.com/brandonrubio/twauter/service/env"
	"github.com/brandonrubio/twauter/service/errorhandler"
	"github.com/brandonrubio/twauter/service/logger"
	"go.uber.org/dig"
)

type Dependency struct {
	Constructor any
	Interface   any
	Token       string
}

func GetDependencies() []Dependency {
	return []Dependency{
		{
			Constructor: env.CreateEnvService,
			Interface:   new(env.IEnvService),
			Token:       "EnvService",
		},
		{
			Constructor: config.CreateAppConfigService,
			Interface:   new(config.IAppConfigService),
			Token:       "AppConfigService",
		},
		{
			Constructor: logger.CreateLoggerService,
			Interface:   new(logger.ILoggerService),
			Token:       "LoggerService",
		},
		{
			Constructor: errorhandler.CreateErrorHandlerService,
			Interface:   new(errorhandler.IErrorHandlerService),
			Token:       "ErrorHandlerService",
		},
		{
			Constructor: handler.CreateTwautHandlerService,
			Interface:   new(handler.ITwautHandlerService),
			Token:       "TwautHandlerService",
		},
		{
			Constructor: api.CreateApiService,
			Interface:   new(api.IApiService),
			Token:       "ApiService",
		},
		{
			Constructor: CreateServiceCatalog,
			Interface:   new(IServiceCatalog),
			Token:       "ServiceCatalog",
		},
	}
}

func InitContainer(dependencies []Dependency) (*dig.Container, *ServiceCatalog) {
	container := dig.New()

	for _, dependency := range dependencies {
		err := container.Provide(dependency.Constructor, dig.As(dependency.Interface), dig.Name(dependency.Token))
		if err != nil {
			panic("failed to provide dig dependency: " + err.Error())
		}
		err = container.Provide(dependency.Constructor)
		if err != nil {
			panic("failed to provide dig dependency: " + err.Error())
		}
	}

	var serviceCatalog *ServiceCatalog

	err := container.Invoke(func(sc *ServiceCatalog) {
		serviceCatalog = sc
	})

	if err != nil {
		panic("failed to initialize service catalog: " + err.Error())
	}

	return container, serviceCatalog
}
