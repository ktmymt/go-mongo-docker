package entity

import (
	"strconv"
	"time"
)

// Project entity has following data.
// Id, Name, Description, Todos, Color, and UpdateDate?
type Project struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Todos       []Todo    `json:"todos"`
	Color       string    `json:"color"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (project *Project) ValidateLength(errors Errors, errorMessage ErrorMessage) Errors {
	maxProjectNameLength := 14
	maxProjectDescriptionLength := 128

	if len(project.Name) > maxProjectNameLength {
		errorMessage.Name = "projectName"
		errorMessage.Message = "The project name must be less than " + strconv.Itoa(maxProjectNameLength) + " characters"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if len(project.Description) > maxProjectDescriptionLength {
		errorMessage.Name = "ProjectDescription"
		errorMessage.Message = "The project description must be less than " + strconv.Itoa(maxProjectDescriptionLength) + " characters"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	return errors
}
