package backend

import (
	"context"
	"github.com/redxiiikk/wails-app-template/backend/api"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/redxiiikk/wails-app-template/backend/infra/database"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type AppEventName string

const (
	ApplicationStarted AppEventName = "app-ready"
)

type AppEvent struct {
	appEventName AppEventName
	parameters   interface{}
}

type AppEventConsumer struct {
	ConsumerName string
	AppEventName AppEventName
	Callback     func(event AppEvent)
}

type App struct {
	Name        string
	context     context.Context
	diContainer *dig.Container

	isReady              bool
	isApplicationStarted bool

	channel   chan AppEvent
	consumers []AppEventConsumer
}

func NewApp() (*App, error) {
	appName := "wails-app-template"

	diContainer, err := NewDIContainer(appName)
	if err != nil {
		return nil, err
	}

	app := &App{
		Name:        appName,
		diContainer: diContainer,

		isReady:              false,
		isApplicationStarted: false,

		channel:   make(chan AppEvent, 10),
		consumers: []AppEventConsumer{},
	}

	app.Subscribe("init application", ApplicationStarted, func(event AppEvent) {
		initApplication(app, diContainer)
	})

	app.StartEventLoop()

	return app, nil
}

func initApplication(app *App, diContainer *dig.Container) {
	app.isReady = false
	_ = diContainer.Invoke(func(config config.ApplicationConfig, databaseClient *database.DatabaseClient) {
		if config.IsGenerateFrontendModel() {
			utils.Logger.Info("[App] Generate frontend model, so skip database migrate...")
			return
		}

		utils.Logger.Info("[App] Initialize database client...")
		err := databaseClient.StartMigrate()
		if err != nil {
			utils.Logger.Error("Failed to initialize database client", zap.Error(err))
			return
		}
		utils.Logger.Info("[App] Database client initialized successfully")
	})
	app.isReady = true
}

func (a *App) Run(invokeFunc func(bind ...interface{})) {
	utils.Logger.Info("[App] Run...")
	err := a.diContainer.Invoke(func(
		echo *api.EchoApi,
		healthCheckApi *api.HealthCheckApi,
		migrateHistoryApi *api.MigrateHistoryApi,
	) {
		utils.Logger.Info("[App] Invoke...")
		invokeFunc(echo, healthCheckApi, migrateHistoryApi)
	})

	if err != nil {
		return
	}
}

func (a *App) SetContext(ctx context.Context) {
	a.context = ctx
}

func (a *App) PublishEvent(eventName AppEventName, parameters interface{}) {
	utils.Logger.Info("[App] Publish event", zap.String("eventName", string(eventName)))
	a.channel <- AppEvent{
		appEventName: eventName,
		parameters:   parameters,
	}
}

func (a *App) Subscribe(name string, eventName AppEventName, callback func(appEvent AppEvent)) {
	utils.Logger.Info("[App] Subscribe to event", zap.String("eventName", string(eventName)))

	a.consumers = append(a.consumers, AppEventConsumer{
		ConsumerName: name,
		AppEventName: eventName,
		Callback:     callback,
	})
}

func (a *App) StartEventLoop() {
	utils.Logger.Info("[App] Start event loop")
	go func() {
		for {
			event := <-a.channel
			utils.Logger.Info("[App] Received event", zap.String("eventName", string(event.appEventName)))
			for _, consumer := range a.consumers {
				if consumer.AppEventName == event.appEventName {
					utils.Logger.Info(
						"[App] Notify consumer",
						zap.String("evetName", string(event.appEventName)),
						zap.String("consumerName", consumer.ConsumerName),
					)
					consumer.Callback(event)
				}
			}
		}
	}()
}
