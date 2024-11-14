package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Customer{},
		&Order{},
		&Product{},
	)
}
