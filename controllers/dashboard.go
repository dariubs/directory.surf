package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserDashboard(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	var dirs []models.Directory
	config.DB.Where("user_id = ?", userID).Preload("Category").Find(&dirs)

	c.HTML(http.StatusOK, "dashboard/index.html", gin.H{
		"Directories": dirs,
	})
}
