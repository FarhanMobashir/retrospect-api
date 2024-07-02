package main

import (
	"log"
	"os"
	"retrospect-api/controllers"
	"retrospect-api/middlewares"
	"retrospect-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	frontendUrl := os.Getenv("FRONTEND_URL")

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontendUrl} // Replace with your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// public routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	{
		log.Println("Registering memory routes")
		routes.MemoryRoutes(protected)
	}

	r.Run(":8080")
}
