package main

import (
	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/middleware"
	"github.com/directorysurf/directory.surf/models"
	"github.com/directorysurf/directory.surf/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.InitDB()
	models.AutoMigrate(db)

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	r.Use(middleware.Session("super-secret-key"))

	routes.AuthRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home/index.html", gin.H{})
	})

	r.Run(":8080")
}
