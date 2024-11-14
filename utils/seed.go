package utils

import (
	"context"
	"order-processing/internal/api/schema"
	"order-processing/internal/models"
	"order-processing/internal/services"
)

func SeedData(service *services.OrderProcessingService) error {
	customers := []models.Customer{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
		{Name: "Bob Johnson", Email: "bob@example.com"},
	}

	products := []schema.CreateProductRequestSchema{
		{Name: "Baby Food", Price: 10.5},
		{Name: "Toothbrush", Price: 2.75},
		{Name: "Soap", Price: 30.5},
	}

	for _, customer := range customers {
		_, err := service.CreateCustomer(context.Background(), customer)
		if err != nil {
			return err
		}
	}
	for _, product := range products {
		_, err := service.CreateProduct(context.Background(), product)
		if err != nil {
			return err
		}
	}
	return nil
}
