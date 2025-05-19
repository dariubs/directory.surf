package controllers

import (
	"net/http"

	"github.com/directorysurf/directory.surf/services"
	"github.com/gin-gonic/gin"
)

func GetPresignedURL(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing filename"})
		return
	}

	url, err := services.GetPresignedUploadURL(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get presigned URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"upload_url": url})
}
