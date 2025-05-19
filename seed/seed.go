package seed

import (
	"fmt"
	"os"
	"strings"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
)

func Run() {
	db := config.DB

	// --- Default Categories
	categories := []string{
		"SaaS Directories",
		"Startup Platforms",
		"AI Directories",
		"Launch Sites",
		"Indie Hacker Tools",
	}

	for _, name := range categories {
		slug := slugify(name)
		var cat models.Category
		if err := db.Where("slug = ?", slug).First(&cat).Error; err != nil {
			db.Create(&models.Category{
				Name: name,
				Slug: slug,
			})
		}
	}

	fmt.Println("✅ Seeded categories")

	email := os.Getenv("ADMIN_EMAIl")
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err == nil {
		user.IsAdmin = true
		db.Save(&user)
		fmt.Println("✅ Promoted user to admin")
	}
}

func slugify(input string) string {
	// Very basic slug, improve if needed
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}
