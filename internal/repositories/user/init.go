package user

import (
	"context"

	"github.com/wendao2000/couply/internal/config"
	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err *errs.Error)
	GetUserByUsername(ctx context.Context, username string) (user *model.User, err *errs.Error)

	InsertUser(ctx context.Context, req *InsertUserRequest) (user *model.User, err *errs.Error)
}

type repository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRepository(cfg *config.Config, db *gorm.DB) Repository {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}
