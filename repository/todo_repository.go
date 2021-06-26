package repository

import (
	"context"
	"go-mongo-docker/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository functions
type Repository interface {
	CreateTodo(*entity.Todo) (*entity.Todo, error)
	UpdateTodo(*entity.Todo, string) (*mongo.UpdateResult, error)
	DeleteTodo(*entity.Todo, string) (*mongo.DeleteResult, error)
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

// CreateTodo saves todo to db
func (t *TodoRepository) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	insert := bson.D{
		{Key: "projectId", Value: todo.ProjectId},
		{Key: "title", Value: todo.Title},
		{Key: "isDone", Value: todo.IsDone},
		{Key: "status", Value: todo.Status},
		{Key: "schedule", Value: todo.Schedule},
	}

	incompleteInsertion, err := collection.InsertOne(ctx, insert)
	avoidPanic(err)

	autoIncrementedId := incompleteInsertion.InsertedID.(primitive.ObjectID).Hex()
	filter := bson.M{"_id": convertToObjectId(autoIncrementedId)}
	update := bson.M{
		"$set": bson.M{
			"id": autoIncrementedId,
		}}

	_, err = collection.UpdateOne(ctx, filter, update)

	return todo, nil
}

// UpdateTodo modify todo data
func (t *TodoRepository) UpdateTodo(todo *entity.Todo, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	filter := bson.M{"_id": convertToObjectId(id)}
	update := bson.M{
		"$set": bson.M{
			"title":    todo.Title,
			"isDone":   todo.IsDone,
			"status":   todo.Status,
			"schedule": todo.Schedule,
		}}

	result, err := collection.UpdateOne(ctx, filter, update)
	avoidPanic(err)

	return result, nil
}

func (t *TodoRepository) DeleteTodo(todo *entity.Todo, id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	filter := bson.M{"_id": convertToObjectId(id)}
	result, err := collection.DeleteOne(ctx, filter)
	avoidPanic(err)

	return result, nil
}
