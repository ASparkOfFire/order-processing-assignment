package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}

// TableName sets the table name for the Customer model
func (Customer) TableName() string {
	return "customers"
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return nil
}
