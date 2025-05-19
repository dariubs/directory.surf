package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/directorysurf/directory.surf/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StartPayment(c *gin.Context) {
	slug := c.Param("slug")
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var dir models.Directory
	if err := config.DB.Where("slug = ?", slug).First(&dir).Error; err != nil || dir.UserID != userID.(uint) {
		c.String(http.StatusNotFound, "Directory not found or unauthorized")
		return
	}

	// You can customize pricing logic based on plan later
	price := int64(2500) // $25 USD
	user := models.User{}
	config.DB.First(&user, dir.UserID)

	payURL, err := services.CreateCheckoutSession(slug, price, user.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "Payment error")
		return
	}

	c.Redirect(http.StatusFound, payURL)
}
