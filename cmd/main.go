package main

import (
	"url-shortener-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "URL Shortener API is Running 🚀",
		})
	})

	router.Run(":8080")
}
