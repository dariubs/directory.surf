package models

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Category{},
		&Tag{},
		&Directory{},
	)
}
