package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Memory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Title  string             `bson:"title"`
	Body   string             `bson:"body"`
}
