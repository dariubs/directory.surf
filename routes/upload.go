package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func UploadRoutes(r *gin.Engine) {
	r.GET("/api/upload-url", controllers.GetPresignedURL)
}
