package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Projects []Project          `json:"projects"`
}

func (user *User) Validation(errors Errors, errorMessage ErrorMessage) Errors {

	if len(user.Username) == 0 {
		errorMessage.Name = "Username"
		errorMessage.Message = "Username is required"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if len(user.Email) == 0 {
		errorMessage.Name = "Email"
		errorMessage.Message = "Email is required"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	return errors
}
