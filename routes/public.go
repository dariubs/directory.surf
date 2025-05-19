package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine) {
	r.GET("/", controllers.PublicDirectoryIndex)
}
