package entity

// Project entity has ID, Name, Description, and Todos
type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Todos       []Todo `json:"todos"`
}
