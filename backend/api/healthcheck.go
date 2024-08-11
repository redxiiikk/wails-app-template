package api

import (
	"github.com/redxiiikk/wails-app-template/backend/infra/database"
	"github.com/redxiiikk/wails-app-template/backend/utils"
)

type HealthCheckApi struct {
	databaseClient *database.SqliteClient
}

type HealthCheckResponse struct {
	Items []HealthCheckItem `json:"items"`
}

type HealthCheckItem struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func NewHealthCheckApi(databaseClient *database.SqliteClient) *HealthCheckApi {
	utils.Logger.Info("[API] create new health check api instance")
	return &HealthCheckApi{
		databaseClient: databaseClient,
	}
}

func (api *HealthCheckApi) HealthCheck() HealthCheckResponse {
	databaseStatus, err := api.databaseClient.HealthCheck()

	return HealthCheckResponse{
		Items: []HealthCheckItem{
			convertToHealthCheckItem("database", databaseStatus, err),
		},
	}
}

func convertToHealthCheckItem(name, status string, err error) HealthCheckItem {
	databaseHealthCheck := HealthCheckItem{
		Name:   name,
		Status: status,
	}
	if err != nil {
		databaseHealthCheck.ErrorMessage = err.Error()
	}
	return databaseHealthCheck
}
