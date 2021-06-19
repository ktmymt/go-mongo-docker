package entity

import "time"

// Project entity has following data.
// Id, Name, Description, Todos, Color, and UpdateDate?
type Project struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
