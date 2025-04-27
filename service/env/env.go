package env

import "os"

const (
	DefaultAppEnv string = "development"
)

type EnvConfig struct {
	AppEnv string
}

type IEnvService interface {
	Init() map[string]string
	GetConfig() EnvConfig
}

type EnvService struct {
	envConfig EnvConfig
}

func (es *EnvService) Init() map[string]string {
	defaults := make(map[string]string)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = DefaultAppEnv
		defaults["APP_ENV"] = DefaultAppEnv
	}

	es.envConfig = EnvConfig{
		AppEnv: appEnv,
	}

	return defaults
}

func (es *EnvService) GetConfig() EnvConfig {
	return es.envConfig
}

func CreateEnvService() *EnvService {
	return &EnvService{}
}
