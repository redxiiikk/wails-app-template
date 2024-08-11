package api

import (
	"github.com/redxiiikk/wails-app-template/backend/entity"
	"github.com/redxiiikk/wails-app-template/backend/service"
)

type MigrateHistoryApi struct {
	migrateService *service.MigrateService
}

func NewMigrateHistoryApi(migrateService *service.MigrateService) *MigrateHistoryApi {
	return &MigrateHistoryApi{
		migrateService: migrateService,
	}
}

func (api *MigrateHistoryApi) MigrateHistory() []entity.MigrateHistory {
	return api.migrateService.QueryMigrateHistory()
}
