package service

import (
	"github.com/wendao2000/couply/internal/config"
	repository "github.com/wendao2000/couply/internal/repositories"
	"github.com/wendao2000/couply/internal/services/profile"
	"github.com/wendao2000/couply/internal/services/user"
)

type Service struct {
	Profile profile.Service
	User    user.Service
}

func InitServices(cfg *config.Config, repo *repository.Repository) (*Service, error) {
	return &Service{
		Profile: profile.NewService(cfg, repo),
		User:    user.NewService(cfg, repo),
	}, nil
}
