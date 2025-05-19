package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-gonic/gin"
)

func AdminListDirectories(c *gin.Context) {
	var dirs []models.Directory
	config.DB.Preload("Category").Preload("User").Find(&dirs)

	c.HTML(http.StatusOK, "admin-directories.html", gin.H{
		"Directories": dirs,
	})
}

func AdminApproveDirectory(c *gin.Context) {
	id := c.Param("id")
	config.DB.Model(&models.Directory{}).Where("id = ?", id).Update("is_approved", true)
	c.Redirect(http.StatusFound, "/admin/directories")
}

func AdminFeatureDirectory(c *gin.Context) {
	id := c.Param("id")
	area := c.Query("area") // "homepage" or "category"

	config.DB.Model(&models.Directory{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_featured": true,
			"featured_on": area,
		})

	c.Redirect(http.StatusFound, "/admin/directories")
}

func AdminDeleteDirectory(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Directory{}, id)
	c.Redirect(http.StatusFound, "/admin/directories")
}
