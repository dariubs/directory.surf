package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func SubmitRoutes(r *gin.Engine) {
	r.GET("/submit", controllers.ShowSubmitForm)
	r.POST("/submit", controllers.SubmitDirectory)
}
