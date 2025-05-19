package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-gonic/gin"
)

func Sitemap(c *gin.Context) {
	var directories []models.Directory
	config.DB.Where("is_approved = ?", true).Find(&directories)

	base := "https://directory.surf"

	var urls []string
	urls = append(urls, base)
	for _, d := range directories {
		urls = append(urls, base+"/directory/"+d.Slug)
	}

	// Basic XML response
	sitemap := `<?xml version="1.0" encoding="UTF-8"?>` +
		`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`

	for _, u := range urls {
		sitemap += `<url><loc>` + u + `</loc></url>`
	}

	sitemap += `</urlset>`

	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, sitemap)
}
