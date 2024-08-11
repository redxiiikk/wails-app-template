package database

import (
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabaseClient(t *testing.T) {
	appConfig := config.ApplicationConfig{
		Env:     config.DevEnv,
		DataDir: t.TempDir(),
	}

	client, err := NewDatabaseClient(&appConfig)
	assert.NoError(t, err, "NewDatabaseClient() failed: ", err)

	status, err := client.HealthCheck()
	assert.NoError(t, err, "HealthCheck() failed: ", err)
	assert.Equal(t, "UP", status, "HealthCheck() failed: ", status)
}
