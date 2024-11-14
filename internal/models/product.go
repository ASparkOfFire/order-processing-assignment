package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null"`
	Price float64 `gorm:"type:decimal(10,2);not null"`
}

// TableName sets the table name for the Product model
func (Product) TableName() string {
	return "products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return nil
}
