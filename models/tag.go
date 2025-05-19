package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name        string       `gorm:"uniqueIndex;not null"`
	Directories []*Directory `gorm:"many2many:directory_tags;"`
}
