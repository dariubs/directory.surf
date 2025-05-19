package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine) {
	r.GET("/", controllers.PublicDirectoryIndex)
	r.GET("/sitemap.xml", controllers.Sitemap)
	r.StaticFile("/robots.txt", "./static/robots.txt")
}
