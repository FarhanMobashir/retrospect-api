# Retrospect API -> Golang + Gin + MOngoDB

Retrospect API is a RESTful API built using Go (Golang) and the Gin framework. It provides endpoints for user registration, authentication, and managing memories.

## Features

- **User Management**: Register new users and authenticate existing users.
- **Memory Management**: Create, read, update, and delete memories associated with authenticated users.
- **Middleware**: Implements authentication middleware to protect routes requiring user authentication.

## Prerequisites

Before running the API, ensure you have the following installed:

- Go (Golang)
- MongoDB

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/retrospect-api.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables:

   - Create a `.env` file based on `.env.example` and configure MongoDB connection settings.

4. Build and run the API:

   ```bash
   go run main.go
   ```

5. The API will start running on `http://localhost:8080`.

## API Endpoints

- **POST /register**: Register a new user.
- **POST /login**: Authenticate and log in a user.
- **Protected Routes** (require authentication):
  - **GET /memories**: Retrieve all memories.
  - **GET /memories/:id**: Retrieve a specific memory by ID.
  - **POST /memories**: Create a new memory.
  - **PUT /memories/:id**: Update an existing memory by ID.
  - **DELETE /memories/:id**: Delete a memory by ID.

## Middleware

The API uses middleware for authentication. Only authenticated users can access protected routes.

## Contributing

Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
