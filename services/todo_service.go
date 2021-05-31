package services

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

// TodoService interface
type TodoService interface {
	GetTodos() ([]*entity.Todo, error)
	CreateTodo(todo *entity.Todo) (*entity.Todo, error)
	UpdateTodo(todo *entity.Todo) (*mongo.UpdateResult, error)
}

type todoService struct {
	Repo repository.Repository
}

// NewTodoService return service
func NewTodoService(repo repository.Repository) TodoService {
	return &todoService{
		Repo: repo,
	}
}

func (ts *todoService) GetTodos() ([]*entity.Todo, error) {
	return ts.Repo.GetTodos()
}

func (ts *todoService) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	return ts.Repo.CreateTodo(todo)
}

func (ts *todoService) UpdateTodo(todo *entity.Todo) (*mongo.UpdateResult, error) {
	return ts.Repo.UpdateTodo(todo)
}
