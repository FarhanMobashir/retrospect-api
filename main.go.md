```go
package main

import (
	"log"
	"retrospect-api/controllers"
	"retrospect-api/middlewares"
	"retrospect-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router with default middleware: logger and recovery (crash-free) middleware
	r := gin.Default()

	// Define public routes accessible without authentication
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes that require authentication
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware()) // Apply custom authentication middleware

	{
		log.Println("Registering memory routes")
		routes.MemoryRoutes(protected) // Register memory routes under the protected group
	}

	// Start the Gin server and listen on port 8080
	r.Run(":8080")
}
```

### Explanation:

1. **Package and Imports**:

   - `package main`: Declares that this file belongs to the main package, making it an executable program.
   - Imports necessary packages:
     - `log`: Standard Go logging package.
     - `retrospect-api/controllers`: Custom package containing user and memory controller functions.
     - `retrospect-api/middlewares`: Custom package containing authentication middleware.
     - `retrospect-api/routes`: Custom package containing routes configuration.
     - `github.com/gin-gonic/gin`: Gin web framework for Go, used for building web applications.

2. **Main Function** (`func main()`):
   - Initializes a new Gin router using `gin.Default()`, which includes default middleware:
     - Logger middleware: Logs HTTP requests.
     - Recovery middleware: Recovers from any panics and returns a 500 Internal Server Error.
3. **Public Routes**:

   - Defines two public routes accessible without authentication:
     - `POST /register`: Endpoint to register a new user, handled by `controllers.RegisterUser`.
     - `POST /login`: Endpoint to authenticate and log in a user, handled by `controllers.LoginUser`.

4. **Protected Routes**:

   - Creates a `protected` group of routes under the root path (`"/"`).
   - Applies custom authentication middleware (`middlewares.AuthMiddleware()`) to the `protected` group. This middleware ensures that only authenticated requests can access routes within this group.

5. **Memory Routes**:

   - Inside the `protected` group, logs a message indicating the registration of memory routes.
   - Calls `routes.MemoryRoutes(protected)`, which presumably registers routes related to memory management under the `protected` group. This could include CRUD operations for memories, requiring authentication.

6. **Server Startup**:
   - Calls `r.Run(":8080")` to start the Gin server and listen on port `8080`.

### Summary:

This code sets up a basic RESTful API using the Gin framework in Go. It defines public routes for user registration and login, and protected routes that require authentication for memory management. Middleware is used to enforce authentication on protected routes, enhancing security. The server starts listening on port `8080` once initialized.
