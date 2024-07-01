# User Controller in Go: Handling Registration and Authentication

This markdown file explains the implementation of a user controller in Go, focusing on user registration (`/register`) and login (`/login`) functionalities.

## 1. RegisterUser Function

The `RegisterUser` function handles user registration by accepting JSON data containing a username and password, hashing the password for security, and storing the user in a MongoDB database. Here's a breakdown of its implementation:

```go
func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
        return
    }
    user.Password = string(hashedPassword)

    // Generate a new ObjectID for the user
    user.ID = primitive.NewObjectID()

    // Insert the user into the MongoDB collection
    _, err = userCollection.InsertOne(context.Background(), user)
    if err != nil {
        // Handle duplicate username error
        if mongo.IsDuplicateKeyError(err) {
            c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Respond with the created user object
    c.JSON(http.StatusCreated, user)
}
```

### Explanation:

- **Error Handling**: Uses `c.BindJSON` to parse JSON data from the request body into a `models.User` struct. If parsing fails, it returns a `400 Bad Request` error.
- **Password Hashing**: Utilizes `bcrypt.GenerateFromPassword` to securely hash the user's password before storing it in the database.

- **Database Interaction**: Inserts the hashed user data into a MongoDB collection (`userCollection`) using `InsertOne`. Handles potential errors, including duplicate username constraints (`mongo.IsDuplicateKeyError`).

- **HTTP Response**: Returns a `201 Created` status code along with the created user data if registration is successful.

## 2. LoginUser Function

The `LoginUser` function handles user authentication by verifying credentials, comparing the hashed password from the database, and generating a JWT token for authenticated sessions:

```go
func LoginUser(c *gin.Context) {
    var input models.User
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    err := userCollection.FindOne(context.Background(), bson.M{"username": input.Username}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Generate JWT token for authenticated user
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Respond with JWT token
    c.JSON(http.StatusOK, gin.H{"token": token})
}
```

### Explanation:

- **Credentials Validation**: Parses JSON data from the request body into a `models.User` struct. Returns a `400 Bad Request` error if parsing fails.

- **Database Query**: Retrieves the user document from the MongoDB collection (`userCollection`) based on the provided username (`bson.M{"username": input.Username}`).

- **Password Verification**: Compares the hashed password stored in the database (`user.Password`) with the provided password (`input.Password`) using `utils.CheckPasswordHash`. Returns a `401 Unauthorized` error if passwords do not match.

- **JWT Token Generation**: Generates a JWT token (`token`) using `utils.GenerateJWT` for authenticated user sessions.

- **HTTP Response**: Returns a `200 OK` status code along with the JWT token if authentication is successful.
