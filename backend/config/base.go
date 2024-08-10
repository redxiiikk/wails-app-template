package config

import (
	"errors"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var CurrentEnv Env = DevEnv

type Env string

const (
	DevEnv  Env = "dev"
	ProdEnv Env = "prod"
)

type ApplicationConfig struct {
	Env Env
}

func NewApplicationConfig(appName string) (ApplicationConfig, error) {
	applicationConfig, err := parseApplicationConfig()
	if err != nil {
		return applicationConfig, err
	}

	applicationConfig.Env = CurrentEnv

	return applicationConfig, nil
}

func parseApplicationConfig() (ApplicationConfig, error) {
	result := ApplicationConfig{}

	applicationConfigPath, err := applicationConfigFilePath()
	if err != nil {
		return result, err
	}

	content, err := os.ReadFile(applicationConfigPath)
	if err != nil {
		return ApplicationConfig{}, errors.New("can't open user config dir")
	}

	err = yaml.Unmarshal(content, &result)
	if err != nil {
		return ApplicationConfig{}, err
	}

	return result, nil
}

func applicationConfigFilePath() (string, error) {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		return "", errors.New("can't open user config dir")
	}

	applicationConfigPath := ""

	switch CurrentEnv {
	case DevEnv:
		applicationConfigPath = path.Join(userConfigDir, "wails-app-template-dev", "config.yaml")
	case ProdEnv:
		applicationConfigPath = path.Join(userConfigDir, "wails-app-template", "config.yaml")
	default:
		return "", errors.New("didn't known env: env=" + string(CurrentEnv))
	}

	return applicationConfigPath, nil
}
