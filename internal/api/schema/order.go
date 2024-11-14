package schema

import (
	"time"
)

type Order struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Products   []Product `json:"products"`
	CustomerID uint      `json:"customer_id"`
	TotalPrice float64   `json:"total_price"`
}

type CreateOrderRequestSchema struct {
	CustomerID uint      `json:"customer_id" validate:"required"`
	Products   []Product `json:"products" validate:"required,min=1,dive"`
}
