package schema

import "time"

type Product struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
}

type CreateProductRequestSchema struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
