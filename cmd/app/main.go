package main

import (
	"log"
	"strconv"

	"github.com/banggibima/go-fiber-restful-api/internal/app"
	"github.com/banggibima/go-fiber-restful-api/internal/config"
	"github.com/banggibima/go-fiber-restful-api/internal/database"
	"github.com/banggibima/go-fiber-restful-api/internal/handlers"
	"github.com/banggibima/go-fiber-restful-api/internal/repositories"
	"github.com/banggibima/go-fiber-restful-api/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	fiberApp := fiber.New()
	myApp := app.NewApp(userHandler)
	myApp.SetupRoutes(fiberApp)

	err = fiberApp.Listen(":" + strconv.Itoa(cfg.Server.Port))
	if err != nil {
		log.Fatalf("error starting the application: %v", err)
	}
}
