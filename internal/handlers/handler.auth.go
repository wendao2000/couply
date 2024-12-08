package handler

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/wendao2000/couply/internal/models"

	"github.com/wendao2000/couply/pkg/errs"
	"github.com/wendao2000/couply/pkg/response"
)

// SignUp handles user registration
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req model.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.FromError(errs.ParamError.WithMessage(err.Error())))
		return
	}

	// Call service to create user
	user, token, err := h.userService.CreateUser(r.Context(), req.Auth)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.HttpCode())
		json.NewEncoder(w).Encode(response.FromError(err))
		return
	}

	// TODO: move to MQ
	go func() {
		prof, err := h.profileService.CreateBasicProfile(r.Context(), &model.BasicProfile{
			UserID:      user.ID,
			Name:        req.Name,
			DateOfBirth: req.DateOfBirth,
		})
		if err != nil {
			// TODO: push back to MQ
			log.Println("Failed while create basic profile, err:", err)
			return
		}

		// TOOD: push notification to user
		log.Printf("Successfully created profile, prof: %+v", prof)
	}()

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

// SignIn handles user authentication
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req model.Auth
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.FromError(errs.ParamError.WithMessage(err.Error())))
		return
	}

	// Call service to authenticate user
	_, token, err := h.userService.Authenticate(r.Context(), &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.HttpCode())
		json.NewEncoder(w).Encode(response.FromError(err))
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
