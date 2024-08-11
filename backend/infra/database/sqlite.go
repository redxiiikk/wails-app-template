package database

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"github.com/glebarez/sqlite"
	"github.com/redxiiikk/wails-app-template/backend/config"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path/filepath"
)

//go:embed all:sql
var sqlScriptDir embed.FS

//goland:noinspection GoNameStartsWithPackageName
type SqliteClient struct {
	client *gorm.DB
}

func NewDatabaseClient(config *config.ApplicationConfig) (*SqliteClient, error) {
	utils.Logger.Info("[Database] create new database client")
	databaseFilePath := filepath.Join(config.DataDir, "sqlite3.db")

	utils.Logger.Info("[Database] register sqlite", zap.String("databaseFilePath", databaseFilePath))
	db, err := gorm.Open(sqlite.Open(databaseFilePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	err = startMigrate(db)
	if err != nil {
		return nil, err
	}

	return &SqliteClient{
		client: db,
	}, err
}

func (database *SqliteClient) HealthCheck() (string, error) {
	tx := database.client.Exec("SELECT 1")
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
