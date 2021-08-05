package entity

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo entity has title and description
/**
 * about "schedule", it shows when to do the task.
 * For instance,
 * 1 -> today
 * 2 -> tomorrow
 * 3 -> the day after tomorrow
 * etc...
 */
type Todo struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ProjectId primitive.ObjectID `json:"projectId"`
	UserId    primitive.ObjectID `json:"userId"`
	Title     string             `json:"title"`
	IsDone    bool               `json:"isDone"`
	Status    string             `json:"status"`
	Schedule  int                `json:"schedule"`
}

func (todo *Todo) Validation(errors Errors, errorMessage ErrorMessage) Errors {
	maxTodoTitleLength := 40

	if len(todo.Title) > maxTodoTitleLength {
		errorMessage.Name = "todoTitile"
		errorMessage.Message = "Todo's title must be less than " + strconv.Itoa(maxTodoTitleLength) + "characters."
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if len(todo.Status) == 0 {
		errorMessage.Name = "todoStatus"
		errorMessage.Message = "Please select the Todo status"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	if todo.Schedule < 0 {
		errorMessage.Name = "todoSchedule"
		errorMessage.Message = "The value of the Todo schedule must be more than 0"
		errors.Errors = append(errors.Errors, errorMessage)
	}

	return errors
}
