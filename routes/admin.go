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

	admin.GET("/directories", controllers.AdminListDirectories)
	admin.GET("/directories/:id/approve", controllers.AdminApproveDirectory)
	admin.GET("/directories/:id/feature", controllers.AdminFeatureDirectory)
	admin.GET("/directories/:id/delete", controllers.AdminDeleteDirectory)
	// Add future admin routes here

}
