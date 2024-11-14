package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"order-processing/internal/api"
	"order-processing/internal/api/handlers"
	"order-processing/internal/models"
	"order-processing/internal/repository"
	"order-processing/internal/services"
	"order-processing/utils"
)

func main() {
	app := fiber.New()
	apiGroup := app.Group("/api")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.AppConfig.Database.DatabaseHost,
		utils.AppConfig.Database.DatabasePort,
		utils.AppConfig.Database.DatabaseUser,
		utils.AppConfig.Database.DatabasePass,
		utils.AppConfig.Database.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	if err := models.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := repository.NewPostgresOrderProcessor(db)
	service := services.NewOrderProcessingService(repo)
	handler := handlers.NewOrderProcessingHandler(service)

	// Seed data
	if err := utils.SeedData(service); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	apiGroup.Get("/", api.MakeHandler(handler.HandleRoot))

	apiGroup.Get("/customers", api.MakeHandler(handler.HandleListCustomers))
	apiGroup.Get("/products", api.MakeHandler(handler.HandleListProducts))
	apiGroup.Get("/customers/:id", api.MakeHandler(handler.HandleGetCustomer))
	apiGroup.Post("/orders", api.MakeHandler(handler.HandleCreateOrder))
	apiGroup.Get("/orders/:id", api.MakeHandler(handler.HandleGetOrder))

	if err := app.Listen(fmt.Sprintf("%s:%d", utils.AppConfig.Server.HTTPListenAddress, utils.AppConfig.Server.HTTPPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Infof("Server Listening at %s:%d", utils.AppConfig.Server.HTTPListenAddress, utils.AppConfig.Server.HTTPPort)
}
