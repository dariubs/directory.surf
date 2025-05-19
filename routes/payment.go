package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	r.GET("/pay/:slug", controllers.StartPayment)
}
