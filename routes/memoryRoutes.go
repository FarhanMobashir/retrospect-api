package routes

import (
	"retrospect-api/controllers"

	"github.com/gin-gonic/gin"
)

func MemoryRoutes(rg *gin.RouterGroup) {
	memoryRoutes := rg.Group("/memories")

	{
		memoryRoutes.POST("/", controllers.CreateMemory)
		memoryRoutes.GET("/", controllers.GetMemories)
		memoryRoutes.GET("/:id", controllers.GetSingleMemory) // Add this line
		memoryRoutes.PUT("/:id", controllers.UpdateMemory)
		memoryRoutes.DELETE("/:id", controllers.DeleteMemory)
	}
}
