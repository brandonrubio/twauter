package logger

import (
	"fmt"
	"os"

	"github.com/brandonrubio/twauter/service/config"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILoggerService interface {
	Init() error
	Log(loglevel string, message string, source string)
}

type LoggerService struct {
	logger           *zap.Logger
	sugar            *zap.SugaredLogger
	appConfigService config.IAppConfigService
}

func (ls *LoggerService) Init() error {
	env := ls.appConfigService.GetEnv()
	loggerConfig := ls.appConfigService.GetLoggerConfig()

	var level zap.AtomicLevel
	switch loggerConfig.Level {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return fmt.Errorf("invalid log level in logger config")
	}

	stdout := zapcore.AddSync(os.Stdout)
	var consoleLoggerConfig zapcore.EncoderConfig
	if env == "production" {
		consoleLoggerConfig = zap.NewProductionEncoderConfig()
	} else {
		consoleLoggerConfig = zap.NewDevelopmentEncoderConfig()
	}
	consoleLoggerConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleLoggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleLoggerConfig)

	var fileLoggerConfig zapcore.EncoderConfig
	if env == "production" {
		fileLoggerConfig = zap.NewProductionEncoderConfig()
	} else {
		fileLoggerConfig = zap.NewDevelopmentEncoderConfig()
	}
	fileLoggerConfig.TimeKey = "timestamp"
	fileLoggerConfig.LevelKey = "level"
	fileLoggerConfig.MessageKey = "msg"
	logFile, err := os.Create("log.txt")
	if err != nil {
		return fmt.Errorf("failed to create log file: %s", err)
	}
	fileEncoder := zapcore.NewJSONEncoder(fileLoggerConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, logFile, level),
	)

	logger := zap.New(core)
	defer logger.Sync() //nolint:errcheck

	sugar := logger.Sugar()
	ls.logger = logger
	ls.sugar = sugar

	return nil
}

func (ls *LoggerService) Log(loglevel string, message string, source string) {
	switch loglevel {
	case "debug":
		ls.sugar.Debugw(message, "source", source)
	case "info":
		ls.sugar.Infow(message, "source", source)
	case "warn":
		ls.sugar.Warnw(message, "source", source)
	case "error":
		ls.sugar.Errorw(message, "source", source)
	case "dpanic":
		ls.sugar.DPanicw(message, "source", source)
	case "panic":
		ls.sugar.Panicw(message, "source", source)
	case "fatal":
		ls.sugar.Fatalw(message, "source", source)
	default:
		ls.sugar.Errorw(fmt.Sprintf("invalid loglevel %s", loglevel), "source", source)
	}
}

type LoggerServiceDependencies struct {
	dig.In
	AppConfigService config.IAppConfigService `name:"AppConfigService"`
}

func CreateLoggerService(dependencies LoggerServiceDependencies) *LoggerService {
	return &LoggerService{
		appConfigService: dependencies.AppConfigService,
	}
}
