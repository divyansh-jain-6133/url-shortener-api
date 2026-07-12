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

func Login(c *gin.Context) {

	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := services.LoginUser(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
	})
}

func Profile(c *gin.Context) {

	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "Protected Route",
		"user_id": userID,
	})
}
