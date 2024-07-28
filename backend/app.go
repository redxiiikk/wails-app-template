package backend

import (
	"github.com/redxiiikk/wails-app-template/backend/api"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/dig"
)

type App struct {
	diContainer *dig.Container
}

func NewApp() (*App, error) {
	fxApp, err := NewDIContainer()
	if err != nil {
		return nil, err
	}

	return &App{
		diContainer: fxApp,
	}, nil
}

func (a *App) Run(invokeFunc func(bind ...interface{})) {
	utils.Logger.Info("[App] Run...")
	err := a.diContainer.Invoke(func(echo *api.EchoApi) {
		invokeFunc(echo)
	})

	if err != nil {
		return
	}
}
