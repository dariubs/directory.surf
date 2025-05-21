package services

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	GithubOAuth *oauth2.Config
	GoogleOAuth *oauth2.Config
)

func InitOAuth() {
	GithubOAuth = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
		RedirectURL:  os.Getenv("APP_DOMAIN") + "/auth/github/callback",
		Scopes:       []string{"user:email"},
	}

	GoogleOAuth = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("APP_DOMAIN") + "/auth/google/callback",
		Scopes:       []string{"email", "profile"},
	}
}
