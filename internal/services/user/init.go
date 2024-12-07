package user

import (
	"context"

	"github.com/wendao2000/couply/internal/config"
	model "github.com/wendao2000/couply/internal/models"
	repository "github.com/wendao2000/couply/internal/repositories"
	"github.com/wendao2000/couply/internal/repositories/user"

	"github.com/wendao2000/couply/pkg/errs"
)

type Service interface {
	CreateUser(ctx context.Context, req *model.Auth) (user *model.User, token *model.UserToken, err *errs.Error)
	Authenticate(ctx context.Context, req *model.Auth) (user *model.User, token *model.UserToken, err *errs.Error)
}

type service struct {
	cfg *config.Config

	userRepo user.Repository
}

func NewService(cfg *config.Config, repo *repository.Repository) Service {
	return &service{
		cfg:      cfg,
		userRepo: repo.User,
	}
}
