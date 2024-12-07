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
	db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username) WHERE username != ''`)
	db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email) WHERE email != ''`)

	return db, nil
}
