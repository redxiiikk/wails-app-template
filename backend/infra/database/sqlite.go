package database

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed all:sql
var sqlScriptDir embed.FS

//goland:noinspection GoNameStartsWithPackageName
type DatabaseClient struct {
	connection *gorm.DB
}

func NewDatabaseClient(config *config.ApplicationConfig) (*DatabaseClient, error) {
	if config.IsGenerateFrontendModel() {
		utils.Logger.Warn("[Database] skip create database connection for generate frontend model")
		return &DatabaseClient{}, nil
	}

	utils.Logger.Info("[Database] create new database connection")
	databaseFilePath := filepath.Join(config.DataDir, "sqlite3.db")

	utils.Logger.Info("[Database] register sqlite", zap.String("databaseFilePath", databaseFilePath))
	db, err := gorm.Open(sqlite.Open(databaseFilePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &DatabaseClient{
		connection: db,
	}, err
}

func (database *DatabaseClient) HealthCheck() (string, error) {
	tx := database.connection.Exec("SELECT 1")
	if tx.Error != nil {
		return "DOWN", tx.Error
	}

	return "UP", nil
}

func hashCode(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
