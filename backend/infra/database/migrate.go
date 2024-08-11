package database

import (
	"errors"
	"github.com/redxiiikk/wails-app-template/backend/entity"
	"github.com/redxiiikk/wails-app-template/backend/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type MigrateHistory struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Automatically managed by GORM for update time

	Key  string `gorm:"index"`
	Hash string
}

func startMigrate(db *gorm.DB) error {
	err := createMigrateTable(db)
	if err != nil {
		return err
	}
	err = migrate(db)
	if err != nil {
		return err
	}
	return nil
}

func createMigrateTable(db *gorm.DB) error {
	utils.Logger.Info("[Database] create migrate table")
	return db.AutoMigrate(&MigrateHistory{})
}

func migrate(db *gorm.DB) error {
	utils.Logger.Info("[Database] init database")
	scripts, err := sqlScriptDir.ReadDir("sql")
	if err != nil {
		return err
	}

	for _, script := range scripts {
		bytes, err := sqlScriptDir.ReadFile("sql/" + script.Name())
		if err != nil {
			return err
		}

		sql := string(bytes)
		sqlHashCod := hashCode(sql)

		existed, err := hasMigrateHistory(db, script.Name(), sqlHashCod)
		if err != nil {
			utils.Logger.Error("[Database] script was changed", zap.String("script", script.Name()), zap.Error(err))
			return err
		}
		if existed {
			utils.Logger.Info("[Database] skip script", zap.String("script", script.Name()))
			continue
		}

		utils.Logger.Info("[Database] run script", zap.String("script", script.Name()))
		err = db.Connection(func(tx *gorm.DB) error {
			exec := tx.Exec(sql)
			if exec.Error != nil {
				return exec.Error
			}
			exec = tx.Create(&MigrateHistory{
				Key:       script.Name(),
				Hash:      sqlHashCod,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
			if exec.Error != nil {
				return exec.Error
			}
			return nil
		})
		utils.Logger.Info("[Database] script executed", zap.String("script", script.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func hasMigrateHistory(db *gorm.DB, key, hashCode string) (bool, error) {
	var count int64
	db.Model(&MigrateHistory{}).Where(&MigrateHistory{Key: key}).Count(&count)

	if count > 0 {
		migrateHistory := MigrateHistory{}
		db.Model(&MigrateHistory{}).Where(&MigrateHistory{Key: key}).Take(&migrateHistory)

		if migrateHistory.Hash != hashCode {
			utils.Logger.Info("[Database] script changed", zap.String("script", key))
			return false, errors.New("script changed: " + key + ", hash: " + hashCode)
		}
	}

	return count != 0, nil
}

func (database *SqliteClient) QueryMigrateHistory() ([]entity.MigrateHistory, error) {
	var migrateHistory []MigrateHistory
	tx := database.client.Model(&MigrateHistory{}).Find(&migrateHistory)
	if tx.Error != nil {
		return nil, tx.Error
	}

	migrateHistoryEntity := make([]entity.MigrateHistory, len(migrateHistory))
	for i, history := range migrateHistory {
		migrateHistoryEntity[i] = entity.MigrateHistory{
			ID:        history.ID,
			Key:       history.Key,
			Hash:      history.Hash,
			CreatedAt: history.CreatedAt,
			UpdatedAt: history.UpdatedAt,
		}
	}

	return migrateHistoryEntity, nil
}
