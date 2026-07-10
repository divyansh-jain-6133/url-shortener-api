package routes

import (
	"url-shortener-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")

	{
		auth := api.Group("/auth")

		{
			auth.POST("/signup", controllers.Signup)
		}
	}
}
