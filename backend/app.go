package backend

import (
	"github.com/redxiiikk/wails-app-template/backend/api"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/dig"
)

type App struct {
	Name        string
	diContainer *dig.Container
}

func NewApp() (*App, error) {
	appName := "wails-app-template"

	fxApp, err := NewDIContainer(appName)
	if err != nil {
		return nil, err
	}

	return &App{
		Name:        appName,
		diContainer: fxApp,
	}, nil
}

func (a *App) Run(invokeFunc func(bind ...interface{})) {
	utils.Logger.Info("[App] Run...")
	err := a.diContainer.Invoke(func(echo *api.EchoApi, healthCheckApi *api.HealthCheckApi) {
		utils.Logger.Info("[App] Invoke...")
		invokeFunc(echo, healthCheckApi)
	})

	if err != nil {
		return
	}
}
