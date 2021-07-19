package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Project  []Project          `json:"project"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Image    string             `json:"image"`
}

func (user *User) Validation(errors Errors, errorMessage ErrorMessage) Errors {

	if user.Username == "" {
		errorMessage.Name = "Username"
		errorMessage.Message = "Please input your username"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if user.Email == "" {
		errorMessage.Name = "Email"
		errorMessage.Message = "Please input your email address"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if user.Image == "" {
		errorMessage.Name = "Image"
		errorMessage.Message = "Please input your image URL"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	return errors
}
