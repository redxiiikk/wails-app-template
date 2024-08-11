package database

import (
	"github.com/glebarez/sqlite"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path/filepath"
	"testing"
)

func TestNewDatabaseClient(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(db *gorm.DB)
		wantErr bool
	}{
		{
			name: "should migrate successfully when database don't exist",
		},
		{
			name: "should migrate successfully when has one migrate history",
			setup: func(db *gorm.DB) {
				err := db.AutoMigrate(&MigrateHistory{})
				if !assert.NoError(t, err, "createMigrateTable() failed: ", err) {
					return
				}

				db.Save(&MigrateHistory{
					Key:  "000000.schema.sql",
					Hash: "576fd5580e5081f6035996bd05e7548f2139527f86e8947ec7bb699b28daef77",
				})
			},
		},
		{
			name:    "should throw err when migrate hash is different",
			wantErr: true,
			setup: func(db *gorm.DB) {
				err := db.AutoMigrate(&MigrateHistory{})
				if !assert.NoError(t, err, "createMigrateTable() failed: ", err) {
					return
				}

				db.Save(&MigrateHistory{
					Key:  "000000.schema.sql",
					Hash: "XXXXX",
				})
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			appConfig := &config.ApplicationConfig{
				Env:     config.DevEnv,
				DataDir: tt.TempDir(),
			}

			if tc.setup != nil {
				databaseFilePath := filepath.Join(appConfig.DataDir, "sqlite3.db")
				db, err := gorm.Open(sqlite.Open(databaseFilePath), &gorm.Config{
					Logger: logger.Default.LogMode(logger.Info),
				})

				assert.NoError(tt, err, "gorm.Open() failed: ", err)

				tc.setup(db)
			}

			sqliteClient, err := NewDatabaseClient(appConfig)

			if tc.wantErr {
				assert.Error(tt, err, "NewDatabaseClient() should failed")
				return
			}

			assert.NoError(tt, err, "NewDatabaseClient() failed: ", err)
			assert.NotNil(tt, sqliteClient, "NewDatabaseClient() sqlite is nil: ", sqliteClient)

			status, err := sqliteClient.HealthCheck()
			assert.NoError(t, err, "HealthCheck() failed: ", err)
			assert.Equal(t, "UP", status, "HealthCheck() failed: ", status)
		})
	}
}
