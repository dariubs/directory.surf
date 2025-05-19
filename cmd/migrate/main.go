package main

import (
	"fmt"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
)

func main() {
	fmt.Println("🔧 Running database migrations...")

	// Load env + init DB
	config.LoadEnv()
	db := config.InitDB()

	// Run migrations
	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Tag{},
		&models.Directory{},
	)
	if err != nil {
		panic("❌ Migration failed: " + err.Error())
	}

	fmt.Println("✅ Migrations complete.")
}
