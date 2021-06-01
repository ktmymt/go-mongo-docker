package entity

// Project entity has following data.
// Id, Name, Description, Todos, Color, and UpdateDate?
type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Todos       []Todo `json:"todos"`
	Color       string `json:"color"`
	// UpdatedAt   ???   `json:"updatedAt"`
}
