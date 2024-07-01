package controllers

import (
	"context"
	"net/http"
	"retrospect-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var memoryCollection *mongo.Collection

func init() {
	// setup mongo client and collection
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOption)

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	memoryCollection = client.Database("goapi").Collection("memories")
}

func CreateMemory(c *gin.Context) {

	var memory models.Memory
	if err := c.BindJSON(&memory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memory.UserID = c.MustGet("userID").(primitive.ObjectID)

	_, err := memoryCollection.InsertOne(context.Background(), memory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, memory)
}

func GetMemories(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)

	cursor, err := memoryCollection.Find(context.Background(), bson.M{"user_id": userID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var memories []models.Memory

	if err := cursor.All(context.Background(), &memories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, memories)
}

func GetSingleMemory(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	memoryID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(memoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var memory models.Memory
	err = memoryCollection.FindOne(context.Background(), bson.M{"_id": objID, "user_id": userID}).Decode(&memory)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "memory not found"})
		return
	}

	c.JSON(http.StatusOK, memory)
}

func UpdateMemory(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	memoryID := c.Param("id")

	var memory models.Memory

	if err := c.BindJSON(&memory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(memoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	filter := bson.M{"_id": objID, "user_id": userID}

	update := bson.M{
		"$set": bson.M{
			"title":       memory.Title, // Update specific fields here
			"description": memory.Body,  // Update specific fields here
			// Add more fields as necessary
		},
	}

	// Perform the update operation
	result, err := memoryCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "memory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Memory updated successfully"})
}

func DeleteMemory(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	memoryID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(memoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	_, err = memoryCollection.DeleteOne(context.Background(), bson.M{"_id": objID, "user_id": userID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "memory deleted"})

}
