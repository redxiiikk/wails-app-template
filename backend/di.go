package backend

import (
	"github.com/redxiiikk/wails-app-template/backend/api"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"go.uber.org/dig"
)

func NewDIContainer() (*dig.Container, error) {
	container := dig.New()

	registerFuncs := []func(*dig.Container) error{
		registerConfig,
		registerApi,
		registerInfra,
	}

	for _, fun := range registerFuncs {
		err := fun(container)
		if err != nil {
			return nil, err
		}
	}

	return container, nil
}

func registerConfig(container *dig.Container) error {
	return container.Provide(
		config.NewApplicationConfig,
	)
}

func registerApi(container *dig.Container) error {
	err := container.Provide(api.NewEchoApi)
	if err != nil {
		return err
	}

	return nil
}

func registerInfra(_ *dig.Container) error {
	return nil
}
