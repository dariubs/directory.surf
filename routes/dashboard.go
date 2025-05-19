package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/directorysurf/directory.surf/middleware"
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.Engine) {
	dashboard := r.Group("/dashboard")
	dashboard.Use(middleware.RequireAuth())
	dashboard.GET("/", controllers.UserDashboard)
}
