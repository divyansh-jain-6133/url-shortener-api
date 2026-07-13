package routes

import (
	"url-shortener-api/controllers"
	"url-shortener-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")

	{
		auth := api.Group("/auth")

		{
			auth.POST("/signup", controllers.Signup)
			auth.POST("/login", controllers.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())

		{
			protected.GET("/profile", controllers.Profile)
			protected.POST("/shorten", controllers.CreateShortURL)
			
		}
		router.GET("/:code", controllers.RedirectURL)

	}
}
