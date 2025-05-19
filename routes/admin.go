package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/directorysurf/directory.surf/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.RequireAdmin())

	admin.GET("/dashboard", controllers.AdminDashboard)

	// Add future admin routes here
}
