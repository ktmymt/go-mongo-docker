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
	GetTodos(string) ([]*entity.Todo, error)
	CreateTodo(*entity.Todo) (*entity.Todo, error)
	UpdateTodo(*entity.Todo) (*mongo.UpdateResult, error)
	DeleteTodo(*entity.Todo) (*mongo.DeleteResult, error)
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

func (t *TodoRepository) GetTodos(id string) ([]*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")
	todoFindResult, err := collection.Find(ctx, bson.M{"projectId": convertToObjectId(id)})
	avoidPanic(err)

	var todos []*entity.Todo
	for todoFindResult.Next(context.Background()) {
		var todo *entity.Todo
		err := todoFindResult.Decode(&todo)
		avoidPanic(err)
		todos = append(todos, todo)
	}

	return todos, nil
}

// CreateTodo saves todo to db
func (t *TodoRepository) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	insert := bson.D{
		{Key: "projectId", Value: todo.ProjectId},
		{Key: "userId", Value: todo.UserId},
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
	avoidPanic(err)

	updateResult := collection.FindOne(ctx, bson.M{"id": autoIncrementedId})

	var newTodo *entity.Todo
	decodedUpdateResult := updateResult.Decode(&newTodo)
	avoidPanic(decodedUpdateResult)

	updateNewUpdatedAtInProject(t, todo, ctx)

	return newTodo, nil
}

// UpdateTodo modify todo data
func (t *TodoRepository) UpdateTodo(todo *entity.Todo) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	filter := bson.M{"_id": todo.Id}
	update := bson.M{
		"$set": bson.M{
			"title":    todo.Title,
			"isDone":   todo.IsDone,
			"status":   todo.Status,
			"schedule": todo.Schedule,
		}}

	result, err := collection.UpdateOne(ctx, filter, update)
	avoidPanic(err)

	updateNewUpdatedAtInProject(t, todo, ctx)

	return result, nil
}

func (t *TodoRepository) DeleteTodo(todo *entity.Todo) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := t.db.Database("taski").Collection("todos")

	filter := bson.M{"_id": todo.Id}
	result, err := collection.DeleteOne(ctx, filter)
	avoidPanic(err)

	updateNewUpdatedAtInProject(t, todo, ctx)

	return result, nil
}

// update UpdatedAt in Project
func updateNewUpdatedAtInProject(t *TodoRepository, todo *entity.Todo, ctx context.Context) {
	projectCollection := t.db.Database("taski").Collection("projects")
	projectFilter := bson.M{"_id": todo.ProjectId}
	projectUpdate := bson.M{"$set": bson.M{"updatedAt": time.Now()}}

	_, projectUpdateErr := projectCollection.UpdateOne(ctx, projectFilter, projectUpdate)
	avoidPanic(projectUpdateErr)
}
