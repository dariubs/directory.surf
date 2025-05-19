package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/directorysurf/directory.surf/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowSubmitForm(c *gin.Context) {
	var categories []models.Category
	var directories []models.Directory

	config.DB.Find(&categories)
	config.DB.Find(&directories)

	c.HTML(http.StatusOK, "submit-form.html", gin.H{
		"Categories":       categories,
		"Alternatives":     directories,
		"TurnstileSiteKey": os.Getenv("TURNSTILE_SITE_KEY"),
	})
}

func SubmitDirectory(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Extract form data
	name := c.PostForm("name")
	url := c.PostForm("url")
	description := c.PostForm("description")
	categoryID := c.PostForm("category_id")
	tags := c.PostForm("tags") // comma-separated
	pricingModel := c.PostForm("pricing_model")
	priceStart := c.PostForm("price_start")
	hasNewsletter := c.PostForm("has_newsletter") == "on"
	hasBlog := c.PostForm("has_blog") == "on"
	alternatives := c.PostFormArray("alternatives")

	// Placeholder for image uploads (next step)
	logoURL := ""
	screenshotURLs := []string{}

	directory := models.Directory{
		Name:           name,
		Slug:           strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		URL:            url,
		Description:    description,
		LogoURL:        logoURL,
		ScreenshotURLs: screenshotURLs,
		PricingModel:   pricingModel,
		PriceStart:     parsePrice(priceStart),
		HasNewsletter:  hasNewsletter,
		HasBlog:        hasBlog,
		CategoryID:     parseUint(categoryID),
		UserID:         userID.(uint),
		IsApproved:     false,
		SubmittedAt:    time.Now(),
	}

	// Save main directory
	if err := config.DB.Create(&directory).Error; err != nil {
		c.HTML(http.StatusBadRequest, "submit/form.html", gin.H{"Error": "Submission failed"})
		return
	}

	// Handle tags
	tagList := strings.Split(tags, ",")
	for _, t := range tagList {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		var tag models.Tag
		if err := config.DB.Where("name = ?", t).First(&tag).Error; err != nil {
			tag = models.Tag{Name: t}
			config.DB.Create(&tag)
		}
		config.DB.Model(&directory).Association("Tags").Append(&tag)
	}

	// Handle alternatives
	for _, altID := range alternatives {
		var alt models.Directory
		if err := config.DB.First(&alt, altID).Error; err == nil {
			config.DB.Model(&directory).Association("Alternates").Append(&alt)
		}
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		// fallback: don't send email if user not found
		fmt.Println("Error loading user:", err)
		return
	}

	go services.SendEmail(user.Email, "Submission Received", fmt.Sprintf(`
  <p>Your submission "<strong>%s</strong>" has been received and is pending review.</p>
  <p>Once approved, it will appear on the site.</p>
`, directory.Name))

	// Notify admin
	go services.SendEmail(os.Getenv("ADMIN_EMAIL"), "New Directory Submission", fmt.Sprintf(`
  <p><strong>%s</strong> was just submitted by %s</p>
  <p><a href="https://directory.surf/admin/directories">Review in Admin Panel</a></p>
`, directory.Name, user.Email))

	// Redirect to payment (next step)
	c.Redirect(http.StatusFound, "/pay/"+directory.Slug)
}

func parseUint(s string) uint {
	var u uint
	fmt.Sscanf(s, "%d", &u)
	return u
}

func parsePrice(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}
