### Memory Model

```go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Memory represents the structure of a memory document in MongoDB.
type Memory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Title  string             `bson:"title"`
	Body   string             `bson:"body"`
}
```

- **Explanation**:
  - **`Memory` struct**: Represents a memory document stored in MongoDB.
    - `ID`: Unique identifier for the memory, automatically generated (`omitempty` ensures it's omitted if empty).
    - `UserID`: ObjectId referencing the user who owns the memory.
    - `Title`: String title of the memory.
    - `Body`: String content or description of the memory.

### User Model

```go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents the structure of a user document in MongoDB.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
```

- **Explanation**:
  - **`User` struct**: Represents a user document stored in MongoDB.
    - `ID`: Unique identifier for the user, automatically generated (`omitempty` ensures it's omitted if empty).
    - `Username`: String representing the username of the user.
    - `Password`: String representing the hashed password of the user.

### Summary

These structs define the schema of documents stored in MongoDB for your application. They use the `go.mongodb.org/mongo-driver/bson/primitive` package to handle MongoDB's `ObjectID` type. This setup allows you to map MongoDB documents directly to Go structs, facilitating seamless interaction with MongoDB through Go code.
