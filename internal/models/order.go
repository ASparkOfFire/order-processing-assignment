package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CustomerID uint      `gorm:"not null"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []Product `gorm:"many2many:order_products;"`
	TotalPrice float64   `gorm:"-"`
}

// TableName sets the table name for the Order model
func (Order) TableName() string {
	return "orders"
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()

	return nil
}
