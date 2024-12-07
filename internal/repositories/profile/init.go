package profile

import (
	"github.com/wendao2000/couply/internal/config"
)

type Repository interface {
}

type repository struct {
}

func NewRepository(cfg *config.Config) Repository {
	return &repository{}
}
