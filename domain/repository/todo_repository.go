package repository

import (
	"context"
	"go-mongo-docker/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Repository functions
type Repository interface {
	CreateTodo(*entity.Todo) (*entity.Todo, error)
}

// TodoRepository structure
type TodoRepository struct {
	db *mongo.Client
}

// NewTodoRepository returns Todo repository
func NewTodoRepository(db *mongo.Client) Repository {
	return &TodoRepository{
		db: db,
	}
}

// CreateTodo saves todo to db
func (t *TodoRepository) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("todos-db").Collection("todos")
	_, err := collection.InsertOne(ctx, *todo)

	if err != nil {
		panic(err)
	}

	return todo, nil
}
