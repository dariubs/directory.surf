package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/directorysurf/directory.surf/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ---- GITHUB ----

func GitHubLogin(c *gin.Context) {
	url := services.GithubOAuth.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GitHubCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := services.GithubOAuth.Exchange(c, code)
	if err != nil {
		c.String(http.StatusInternalServerError, "OAuth error")
		return
	}

	client := services.GithubOAuth.Client(c, token)
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch user email")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var emails []struct {
		Email   string `json:"email"`
		Primary bool   `json:"primary"`
	}
	json.Unmarshal(body, &emails)

	var userEmail string
	for _, e := range emails {
		if e.Primary {
			userEmail = e.Email
			break
		}
	}

	if userEmail == "" {
		c.String(http.StatusUnauthorized, "No email found")
		return
	}

	loginOrCreateUser(c, userEmail)
}

// ---- GOOGLE ----

func GoogleLogin(c *gin.Context) {
	url := services.GoogleOAuth.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := services.GoogleOAuth.Exchange(c, code)
	if err != nil {
		c.String(http.StatusInternalServerError, "OAuth error")
		return
	}

	client := services.GoogleOAuth.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch user info")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var profile struct {
		Email string `json:"email"`
	}
	json.Unmarshal(body, &profile)

	if profile.Email == "" {
		c.String(http.StatusUnauthorized, "No email found")
		return
	}

	loginOrCreateUser(c, profile.Email)
}

// ---- Shared Logic ----

func loginOrCreateUser(c *gin.Context, email string) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		user = models.User{Email: email}
		config.DB.Create(&user)
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}
