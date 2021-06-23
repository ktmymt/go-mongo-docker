package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Title     string             `json:"title"`
	IsDone    bool               `json:"isDone"`
	Status    string             `json:"status"`
	Schedule  int                `json:"schedule"`
}
