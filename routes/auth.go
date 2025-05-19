package routes

import (
	"github.com/directorysurf/directory.surf/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/signup", controllers.ShowSignup)
	r.POST("/signup", controllers.Signup)

	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)

	r.GET("/logout", controllers.Logout)
}
