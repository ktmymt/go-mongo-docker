package repository

import (
	"context"
	"go-mongo-docker/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Repository functions
type Repository interface {
	GetTodos() ([]*entity.Todo, error)
	CreateTodo(*entity.Todo) (*entity.Todo, error)
	UpdateTodo(*entity.Todo) (*entity.Todo, error)
}

// TodoRepository structure has db
type TodoRepository struct {
	db *mongo.Client
}

// NewTodoRepository returns Todo repository
func NewTodoRepository(db *mongo.Client) Repository {
	return &TodoRepository{
		db: db,
	}
}

// GetTodos returns all todos
func (t *TodoRepository) GetTodos() ([]*entity.Todo, error) {
	collection := t.db.Database("todos-db").Collection("todos")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	var results []*entity.Todo

	for cur.Next(context.Background()) {
		// create a value into which the single document can be decoded
		var elem *entity.Todo
		err := cur.Decode(&elem)

		if err != nil {
			panic(err)
		}

		results = append(results, elem)
	}

	return results, nil
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

// UpdateTodo modify todo data
func (t *TodoRepository) UpdateTodo(todo *entity.Todo) (*entity.Todo, error) {
	return todo, nil
}
