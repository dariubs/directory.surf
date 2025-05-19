package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string      `gorm:"uniqueIndex;not null"`
	Slug        string      `gorm:"uniqueIndex;not null"`
	Directories []Directory `gorm:"foreignKey:CategoryID"`
}
