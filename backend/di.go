package backend

import (
	"github.com/redxiiikk/wails-app-template/backend/api"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/redxiiikk/wails-app-template/backend/infra/database"
	"github.com/redxiiikk/wails-app-template/backend/service"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/dig"
)

func NewDIContainer(appName string) (*dig.Container, error) {
	container := dig.New()

	registerFuncs := []func(*dig.Container) error{
		registerConfig(appName),
		registerInfra,
		registerApi,
		registerService,
	}

	for _, fun := range registerFuncs {
		err := fun(container)
		if err != nil {
			return nil, err
		}
	}

	return container, nil
}

func registerConfig(appName string) func(container *dig.Container) error {
	return func(container *dig.Container) error {
		utils.Logger.Info("[DI] register config")
		return container.Provide(
			config.NewApplicationConfig(appName),
		)
	}
}

func registerApi(container *dig.Container) error {
	utils.Logger.Info("[DI] register api")
	err := container.Provide(api.NewEchoApi)
	if err != nil {
		return err
	}

	err = container.Provide(api.NewHealthCheckApi)
	if err != nil {
		return err
	}

	err = container.Provide(api.NewMigrateHistoryApi)
	if err != nil {
		return err
	}

	return nil
}

func registerService(container *dig.Container) error {
	utils.Logger.Info("[DI] register service")
	err := container.Provide(service.NewMigrateService)
	if err != nil {
		return err
	}

	return nil
}

func registerInfra(container *dig.Container) error {
	utils.Logger.Info("[DI] register infra")
	err := container.Provide(database.NewDatabaseClient)
	if err != nil {
		return err
	}

	return nil
}
