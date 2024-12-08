package profile

import (
	"context"

	"github.com/wendao2000/couply/internal/config"
	model "github.com/wendao2000/couply/internal/models"
	repository "github.com/wendao2000/couply/internal/repositories"
	profRepo "github.com/wendao2000/couply/internal/repositories/profile"

	"github.com/wendao2000/couply/pkg/errs"
)

type Service interface {
	CreateBasicProfile(ctx context.Context, req *model.BasicProfile) (prof *model.Profile, err *errs.Error)
}

type service struct {
	profRepo profRepo.Repository
}

func NewService(cfg *config.Config, repo *repository.Repository) Service {
	return &service{
		profRepo: repo.Profile,
	}
}
