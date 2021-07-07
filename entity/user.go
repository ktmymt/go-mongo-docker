package entity

import (
	"image"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Image    image.Image        `json:"image"`
	Projects []Project          `json:"projects"`
}
