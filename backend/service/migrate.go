package service

import (
	"github.com/redxiiikk/wails-app-template/backend/entity"
	"github.com/redxiiikk/wails-app-template/backend/infra/database"
)

type MigrateService struct {
	databaseClient *database.DatabaseClient
}

func NewMigrateService(databaseClient *database.DatabaseClient) *MigrateService {
	return &MigrateService{
		databaseClient: databaseClient,
	}
}

func (service *MigrateService) QueryMigrateHistory() []entity.MigrateHistory {
	history, err := service.databaseClient.QueryMigrateHistory()
	if err != nil {
		return nil
	}

	return history
}
