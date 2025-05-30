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
	hasNewsletter := c.Query("has_newsletter") == "on"
	hasBlog := c.Query("has_blog") == "on"

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

	if hasNewsletter {
		query = query.Where("has_newsletter = ?", true)
	}

	if hasBlog {
		query = query.Where("has_blog = ?", true)
	}

	query.Order("is_featured DESC, created_at DESC").Find(&directories)

	var categories []models.Category
	var tags []models.Tag
	config.DB.Find(&categories)
	config.DB.Find(&tags)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Directories":     directories,
		"Categories":      categories,
		"Tags":            tags,
		"CurrentCategory": categorySlug,
		"CurrentTag":      tagFilter,
		"SearchQuery":     search,
		"HasNewsletter":   hasNewsletter,
		"HasBlog":         hasBlog,
	})
}

func PublicDirectoryDetail(c *gin.Context) {
	slug := c.Param("slug")

	var dir models.Directory
	err := config.DB.Preload("Category").
		Preload("Tags").
		Preload("Alternates").
		Where("slug = ? AND is_approved = ?", slug, true).
		First(&dir).Error

	if err != nil {
		c.String(http.StatusNotFound, "Directory not found")
		return
	}

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Directory":   dir,
		"Title":       dir.Name + " | Directory Surf",
		"Description": dir.Description,
		"Canonical":   "https://directory.surf/directory/" + dir.Slug,
	})
}
