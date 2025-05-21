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

	r.GET("/auth/github", controllers.GitHubLogin)
	r.GET("/auth/github/callback", controllers.GitHubCallback)

	r.GET("/auth/google", controllers.GoogleLogin)
	r.GET("/auth/google/callback", controllers.GoogleCallback)

	r.GET("/logout", controllers.Logout)
}
