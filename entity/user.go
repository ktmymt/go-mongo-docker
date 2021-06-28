package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Projects []Project          `json:"projects"`
}
