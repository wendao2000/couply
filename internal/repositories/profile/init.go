package profile

import (
	"context"

	"github.com/wendao2000/couply/internal/config"
	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"

	"gorm.io/gorm"
)

type Repository interface {
	InsertProfile(ctx context.Context, req *InsertProfileRequest) (prof *model.Profile, err *errs.Error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(cfg *config.Config, db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
