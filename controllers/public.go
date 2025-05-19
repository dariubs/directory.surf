package controllers

import (
	"net/http"
	"strings"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-gonic/gin"
)

func PublicDirectoryIndex(c *gin.Context) {
	categorySlug := c.Query("category")
	tagFilter := c.Query("tag")
	search := c.Query("q")

	var directories []models.Directory
	query := config.DB.Preload("Category").Preload("Tags").
		Where("is_approved = ?", true)

	if search != "" {
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")
	}

	if categorySlug != "" {
		var cat models.Category
		if err := config.DB.Where("slug = ?", categorySlug).First(&cat).Error; err == nil {
			query = query.Where("category_id = ?", cat.ID)
		}
	}

	if tagFilter != "" {
		query = query.Joins("JOIN directory_tags ON directory_tags.directory_id = directories.id").
			Joins("JOIN tags ON tags.id = directory_tags.tag_id").
			Where("tags.name = ?", tagFilter)
	}

	query.Order("is_featured DESC, created_at DESC").Find(&directories)

	var categories []models.Category
	var tags []models.Tag
	config.DB.Find(&categories)
	config.DB.Find(&tags)

	c.HTML(http.StatusOK, "public/index.html", gin.H{
		"Directories":     directories,
		"Categories":      categories,
		"Tags":            tags,
		"CurrentCategory": categorySlug,
		"CurrentTag":      tagFilter,
		"SearchQuery":     search,

		"Title":       "Directory Surf — Discover the Best Startup Directories",
		"Description": "A curated directory of web directories for startups, SaaS, AI, and more.",
		"Canonical":   "https://directory.surf/",
	})
}
