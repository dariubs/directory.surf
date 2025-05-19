package models

import (
	"time"

	"gorm.io/gorm"
)

type Directory struct {
	gorm.Model

	Name           string `gorm:"not null"`
	Slug           string `gorm:"uniqueIndex;not null"`
	URL            string `gorm:"not null"`
	Description    string `gorm:"type:text"`
	LogoURL        string
	ScreenshotURLs []string `gorm:"type:text[]"`
	PricingModel   string
	PriceStart     float64

	HasNewsletter bool
	HasBlog       bool

	CategoryID uint
	Category   Category

	Tags       []*Tag       `gorm:"many2many:directory_tags;"`
	Alternates []*Directory `gorm:"many2many:directory_alternatives;joinForeignKey:DirectoryID;joinReferences:AlternateID"`

	UserID uint // owner
	User   User

	IsApproved bool
	IsFeatured bool
	FeaturedOn string // "homepage", "category", etc.

	SubmittedAt time.Time
}
