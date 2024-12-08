package main

import (
	"log"
	"net/http"

	"github.com/wendao2000/couply/internal/config"
	"github.com/wendao2000/couply/internal/database"
	handler "github.com/wendao2000/couply/internal/handlers"
	repository "github.com/wendao2000/couply/internal/repositories"
	service "github.com/wendao2000/couply/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// initialize .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed while initializing env file, err:", err)
	}

	// Initialize config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Failed while initializing config, err:", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed while initializing config, err:", err)
	}

	// Initialize repository
	repo, err := repository.InitRepositories(cfg, db)
	if err != nil {
		log.Fatal("Failed while initializing repositories, err:", err)
	}

	// Initialize service
	svc, err := service.InitServices(cfg, repo)
	if err != nil {
		log.Fatal("Failed while initializing services, err:", err)
	}

	// Initialize handler
	hdlr, err := handler.NewHandler(cfg, svc)
	if err != nil {
		log.Fatal("Failed while initializing handler, err:", err)
	}

	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	hdlr.RegisterRoutes(r)

	log.Println("Server is running on port :2025")
	err = http.ListenAndServe(":2025", r)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
