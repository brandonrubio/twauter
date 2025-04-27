package errorhandler

import (
	"go.uber.org/dig"

	"github.com/brandonrubio/twauter/service/logger"
)

type IErrorHandlerService interface {
	HandleError(err error, source string)
}

type ErrorHandlerService struct {
	loggerService logger.ILoggerService
}

func (ehs *ErrorHandlerService) HandleError(err error, source string) {
	ehs.loggerService.Log("error", err.Error(), source)
}

type ErrorHandlerServiceDependencies struct {
	dig.In
	LoggerService logger.ILoggerService `name:"LoggerService"`
}

func CreateErrorHandlerService(dependencies ErrorHandlerServiceDependencies) *ErrorHandlerService {
	return &ErrorHandlerService{
		loggerService: dependencies.LoggerService,
	}
}
