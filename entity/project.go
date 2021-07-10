package entity

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project entity has following data.
// Id, Name, Description, Todos, Color, and UpdateDate?
type Project struct {
	Id          primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	UserIds     []primitive.ObjectID `json:"userIds"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Todos       []Todo               `json:"todos"`
	Color       string               `json:"color"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

func (project *Project) Validation(errors Errors, errorMessage ErrorMessage) Errors {
	minProjectNameLength := 2
	maxProjectNameLength := 13
	maxProjectDescriptionLength := 128

	if len(project.Name) > maxProjectNameLength || len(project.Name) < minProjectNameLength {
		errorMessage.Name = "projectName"
		errorMessage.Message = "Project name must be between " + strconv.Itoa(minProjectNameLength) + "-" + strconv.Itoa(maxProjectNameLength) + " characters"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if len(project.Description) > maxProjectDescriptionLength {
		errorMessage.Name = "projectDescription"
		errorMessage.Message = "Project description must be less than " + strconv.Itoa(maxProjectDescriptionLength) + " characters"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if project.Color == "" {
		errorMessage.Name = "projectColor"
		errorMessage.Message = "Please select your project color"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	return errors
}
