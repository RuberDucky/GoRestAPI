package routes

import (
	"my-gin-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthController()

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
	}
} 