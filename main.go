package main

import (
	"embed"
	"github.com/redxiiikk/wails-app-template/backend"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/zap"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app, err := backend.NewApp()
	if err != nil {
		utils.Logger.Error("start application failed", zap.String("errorMessage", err.Error()))
	}
	app.Run(wailsRun(app))
}

func wailsRun(app *backend.App) func(bind ...interface{}) {
	return func(bind ...interface{}) {
		utils.Logger.Info("[App] Wails Start Run...")
		err := wails.Run(&options.App{
			Title:      app.Name,
			Fullscreen: true,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			Bind: bind,
		})

		if err != nil {
			println("Error:", err.Error())
		}
	}

}
