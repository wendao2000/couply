package handler

import (
	"github.com/wendao2000/couply/internal/config"
	service "github.com/wendao2000/couply/internal/services"
	"github.com/wendao2000/couply/internal/services/profile"
	"github.com/wendao2000/couply/internal/services/user"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	profileService profile.Service
	userService    user.Service
}

func NewHandler(cfg *config.Config, svc *service.Service) (*Handler, error) {
	return &Handler{
		profileService: svc.Profile,
		userService:    svc.User,
	}, nil
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/signup", h.SignUp)
	r.Post("/signin", h.SignIn)
}
