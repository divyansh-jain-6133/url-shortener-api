package controllers

import (
	"net/http"

	"url-shortener-api/models"
	"url-shortener-api/services"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var req models.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := services.RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered Successfully",
	})
}
