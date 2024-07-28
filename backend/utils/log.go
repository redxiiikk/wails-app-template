package utils

import (
	"github.com/redxiiikk/wails-app-template/backend/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	//goland:noinspection GoBoolExpressions
	if config.CurrentEnv != config.ProdEnv {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}
}
