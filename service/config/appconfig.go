package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/brandonrubio/twauter/service/env"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type LoggerConfig struct {
	Level string `json:"Level"`
}

type ApiConfig struct {
	Port string `json:"port"`
}

type EnvironmentConfig struct {
	Logger LoggerConfig `json:"logger"`
	Api    ApiConfig    `json:"api"`
}

type AppConfig struct {
	Development EnvironmentConfig `json:"development"`
	Production  EnvironmentConfig `json:"production"`
}

type IAppConfigService interface {
	Init() error
	GetEnv() string
	GetLoggerConfig() LoggerConfig
	GetApiConfig() ApiConfig
}

type AppConfigService struct {
	env        string
	appConfig  AppConfig
	envService env.IEnvService
}

func (acs *AppConfigService) Init() error {
	viper.SetConfigName("appconfig")
	viper.SetConfigType("json")

	executablePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get binary executable path %s", err)
	}
	viper.AddConfigPath(filepath.Dir(executablePath))

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read appconfig: %s", err)
	}

	err = viper.Unmarshal(&acs.appConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal appconfig.json: %s", err)
	}

	acs.env = acs.envService.GetConfig().AppEnv

	return nil
}

func (acs *AppConfigService) GetEnv() string {
	return acs.env
}

func (acs *AppConfigService) GetLoggerConfig() LoggerConfig {
	if acs.env == "production" {
		return acs.appConfig.Production.Logger
	} else {
		return acs.appConfig.Development.Logger
	}
}

func (acs *AppConfigService) GetApiConfig() ApiConfig {
	if acs.env == "production" {
		return acs.appConfig.Production.Api
	} else {
		return acs.appConfig.Development.Api
	}
}

type AppConfigServiceDependencies struct {
	dig.In
	EnvService env.IEnvService `name:"EnvService"`
}

func CreateAppConfigService(dependencies AppConfigServiceDependencies) *AppConfigService {
	return &AppConfigService{
		envService: dependencies.EnvService,
	}
}
