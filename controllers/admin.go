package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-dashboard.html", nil)
}
