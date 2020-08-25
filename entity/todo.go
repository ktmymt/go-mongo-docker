package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo entity has title and description
type Todo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}
