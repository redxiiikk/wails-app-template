package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Env string

const (
	DevEnv  Env = "dev"
	ProdEnv Env = "prod"
)

var CurrentEnv = DevEnv

type ApplicationConfig struct {
	AppName string `yaml:"-"`
	Env     Env    `yaml:"-"`
	DataDir string `yaml:"-"`
}

func NewApplicationConfig(appName string) func() (*ApplicationConfig, error) {
	return func() (*ApplicationConfig, error) {
		dataDir, err := applicationDataDir(appName)
		if err != nil {
			return &ApplicationConfig{}, err
		}

		applicationConfig, err := parseApplicationConfig(dataDir)
		if err != nil {
			return &applicationConfig, err
		}

		applicationConfig.AppName = appName
		applicationConfig.Env = CurrentEnv
		applicationConfig.DataDir = dataDir

		return &applicationConfig, nil
	}
}

func applicationDataDir(appName string) (string, error) {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		return "", errors.New("can't open user config dir")
	}

	var dataDir string

	switch CurrentEnv {
	case DevEnv:
		dataDir = path.Join(userConfigDir, appName+"-dev")
	case ProdEnv:
		dataDir = path.Join(userConfigDir, appName)
	default:
		return "", errors.New("didn't known env: env=" + string(CurrentEnv))
	}

	f, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return "", err
		}
	} else {
		if !f.IsDir() {
			return "", errors.New("dataDir is not directory")
		}
	}

	return dataDir, nil
}

func parseApplicationConfig(dataDir string) (ApplicationConfig, error) {
	result := ApplicationConfig{}

	applicationConfigPath := filepath.Join(dataDir, "config.yaml")

	_, err := os.Stat(applicationConfigPath)
	if os.IsNotExist(err) {
		return ApplicationConfig{}, nil
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
