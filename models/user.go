package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents the user model in the database
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"` // e.g., "admin" or "user"
}