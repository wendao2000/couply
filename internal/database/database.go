package database

import (
	"github.com/wendao2000/couply/internal/config"
	model "github.com/wendao2000/couply/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{})

	return db, nil
}
