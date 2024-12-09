package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/delivery/http/route"
	"github.com/lunadiotic/golang-travel-booking/internal/repository"
	"github.com/lunadiotic/golang-travel-booking/internal/usecase"
	"github.com/lunadiotic/golang-travel-booking/pkg/config"
	"github.com/lunadiotic/golang-travel-booking/pkg/database"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Setup Database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Setup Repository
	userRepo := repository.NewUserRepository(db)

	// Setup UseCase dengan repository
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Setup Router
	router := route.NewRouter(userUseCase)

	// Setup Gin
	r := gin.Default()

	// Setup routes
	router.SetupRoutes(r)

	// Run server
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Cannot run server:", err)
	}
}
