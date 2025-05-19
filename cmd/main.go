package main

import (
	"os"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/middleware"
	"github.com/directorysurf/directory.surf/routes"
	"github.com/directorysurf/directory.surf/seed"
	"github.com/directorysurf/directory.surf/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	services.InitStripe()
	services.InitR2()
	services.InitEmail()

	if os.Getenv("SEED") == "true" {
		seed.Run()
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	r.Use(middleware.Session("super-secret-key"))

	routes.AuthRoutes(r)
	routes.AdminRoutes(r)
	routes.SubmitRoutes(r)
	routes.PaymentRoutes(r)
	routes.UploadRoutes(r)
	routes.DashboardRoutes(r)
	routes.PublicRoutes(r)

	r.Run(":8080")
}
