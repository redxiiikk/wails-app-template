package database

import (
	"embed"
	"github.com/glebarez/sqlite"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"path/filepath"
)

//go:embed schema.sql
var schema embed.FS

//goland:noinspection GoNameStartsWithPackageName
type DatabaseClient struct {
	client *gorm.DB
}

func NewDatabaseClient(config *config.ApplicationConfig) (*DatabaseClient, error) {
	databaseFilePath := filepath.Join(config.DataDir, "sqlite3.db")

	utils.Logger.Info("[Database] register sqlite", zap.String("databaseFilePath", databaseFilePath))
	db, err := gorm.Open(sqlite.Open(databaseFilePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = initDatabase(db)
	if err != nil {
		return nil, err
	}

	return &DatabaseClient{
		client: db,
	}, err
}

func initDatabase(db *gorm.DB) error {
	bytes, err := schema.ReadFile("schema.sql")
	if err != nil {
		return err
	}

	err = db.Connection(func(tx *gorm.DB) error {
		tx.Exec(string(bytes))
		return tx.Error
	})
	if err != nil {
		return err
	}
	return nil
}

func (database *DatabaseClient) HealthCheck() (string, error) {
	tx := database.client.Exec("SELECT 1")
	if tx.Error != nil {
		return "DOWN", tx.Error
	}

	return "UP", nil
}
