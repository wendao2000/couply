package profile

import (
	"context"

	"github.com/wendao2000/couply/internal/config"
	repository "github.com/wendao2000/couply/internal/repositories"

	"github.com/wendao2000/couply/pkg/errs"
)

type Service interface {
	CreateBasicProfile(ctx context.Context, req *BasicProfile) (err *errs.Error)
}

type service struct {
}

func NewService(cfg *config.Config, repo *repository.Repository) Service {
	return &service{}
}
