### Routing in Go with Gin Framework

#### User Routes

```go
package routes

import (
	"retrospect-api/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes sets up routes for user authentication and registration.
func UserRoutes(r *gin.Engine) {
	// Register endpoint for user registration
	r.POST("/register", controllers.RegisterUser)

	// Login endpoint for user authentication
	r.POST("/login", controllers.LoginUser)
}
```

- **Explanation**:
  - **`UserRoutes` Function**: This function takes a `gin.Engine` (`r`) as an argument and defines two routes:
    - `POST /register`: Calls the `controllers.RegisterUser` function to handle user registration.
    - `POST /login`: Calls the `controllers.LoginUser` function to handle user login.

#### Memory Routes

```go
package routes

import (
	"retrospect-api/controllers"

	"github.com/gin-gonic/gin"
)

// MemoryRoutes sets up routes for memory CRUD operations.
func MemoryRoutes(rg *gin.RouterGroup) {
	// Group routes under /memories path
	memoryRoutes := rg.Group("/memories")

	// Define CRUD endpoints for memories
	{
		// POST /memories/ - Create a new memory
		memoryRoutes.POST("/", controllers.CreateMemory)

		// GET /memories/ - Retrieve all memories
		memoryRoutes.GET("/", controllers.GetMemories)

		// GET /memories/:id - Retrieve a single memory by ID
		memoryRoutes.GET("/:id", controllers.GetSingleMemory)

		// PUT /memories/:id - Update a memory by ID
		memoryRoutes.PUT("/:id", controllers.UpdateMemory)

		// DELETE /memories/:id - Delete a memory by ID
		memoryRoutes.DELETE("/:id", controllers.DeleteMemory)
	}
}
```

- **Explanation**:
  - **`MemoryRoutes` Function**: This function sets up routes for memory management within a specified `gin.RouterGroup` (`rg`).
    - **`rg.Group("/memories")`**: Creates a sub-group of routes under the `/memories` path.
    - **CRUD Endpoints**:
      - `POST /memories/`: Creates a new memory using `controllers.CreateMemory`.
      - `GET /memories/`: Retrieves all memories using `controllers.GetMemories`.
      - `GET /memories/:id`: Retrieves a specific memory by ID using `controllers.GetSingleMemory`.
      - `PUT /memories/:id`: Updates a memory by ID using `controllers.UpdateMemory`.
      - `DELETE /memories/:id`: Deletes a memory by ID using `controllers.DeleteMemory`.

### Summary

- **User Routes** (`UserRoutes`):

  - Handles user registration and authentication.
  - Public endpoints accessible without authentication.

- **Memory Routes** (`MemoryRoutes`):
  - Manages CRUD operations for memories.
  - Routes are grouped under `/memories` and require authentication to access.

These routes organize the functionality of your application into logical groups (`/register`, `/login` for users; `/memories` for memories) and utilize controller functions (`controllers.*`) to handle business logic for each endpoint. The use of `gin.RouterGroup` allows for modular and structured route definitions, enhancing maintainability and readability of your codebase.
