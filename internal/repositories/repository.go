package repository

import (
	"github.com/wendao2000/couply/internal/config"
	"github.com/wendao2000/couply/internal/repositories/profile"
	"github.com/wendao2000/couply/internal/repositories/user"

	"gorm.io/gorm"
)

type Repository struct {
	Profile profile.Repository
	User    user.Repository
}

func InitRepositories(cfg *config.Config, db *gorm.DB) (*Repository, error) {
	return &Repository{
		Profile: profile.NewRepository(cfg, db),
		User:    user.NewRepository(cfg, db),
	}, nil
}
