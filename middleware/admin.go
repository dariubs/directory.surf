package middleware

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil || !user.IsAdmin {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
