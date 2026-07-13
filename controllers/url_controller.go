package controllers

import (
	"net/http"

	"url-shortener-api/models"
	"url-shortener-api/services"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {

	var req models.ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	userID := uint(userIDValue.(float64))

	url, err := services.CreateShortURL(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Short URL Created",
		"short_code": url.ShortURL,
		"long_url":   url.LongURL,
	})
}

func RedirectURL(c *gin.Context) {

	code := c.Param("code")

	url, err := services.GetURLByCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Short URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
