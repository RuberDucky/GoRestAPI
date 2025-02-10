package main

import (
	"my-gin-project/config"
	"my-gin-project/db"
	"my-gin-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.GinMode)

	// Initialize database
	db.InitDB(config.GetDBURL())

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupUserRoutes(r)
	routes.SetupAuthRoutes(r)

	// Basic health check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})

	// Run server with configured port
	r.Run(":" + config.AppConfig.Server.Port)
}
