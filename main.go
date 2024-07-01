package main

import (
	"log"
	"retrospect-api/controllers"
	"retrospect-api/middlewares"
	"retrospect-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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
