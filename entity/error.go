package entity

type Errors struct {
	Errors []ErrorMessage
}

type ErrorMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
