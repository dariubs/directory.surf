package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/config"
	"github.com/directorysurf/directory.surf/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{Email: email, Password: string(hashedPassword)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"Error": "Email already exists"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Invalid email or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
