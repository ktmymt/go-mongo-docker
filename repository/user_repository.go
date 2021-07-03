package repository

import (
	"context"
	"go-mongo-docker/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetOwnProjects(string, string) ([]*entity.Project, error)
}

type userRepository struct {
	db *mongo.Client
}

// NewProjectRepository returns "ProjectRepository"
func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

func createNewUser(ur *userRepository, username string, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := ur.db.Database("taski").Collection("users")
	insert := bson.D{
		{Key: "username", Value: username},
		{Key: "email", Value: email},
	}

	incompleteInsertion, err := collection.InsertOne(ctx, insert)
	avoidPanic(err)

	autoIncrementedId := incompleteInsertion.InsertedID.(primitive.ObjectID)
	filter := bson.M{"_id": autoIncrementedId}
	update := bson.M{
		"$set": bson.M{
			"id": autoIncrementedId,
		}}

	_, err = collection.UpdateOne(ctx, filter, update)
	avoidPanic(err)
}

/**
 * @summary: gets projects by user id
 * @return : projects, error
 */
func (ur *userRepository) GetOwnProjects(username string, email string) ([]*entity.Project, error) {

	// get projects
	projectCollection := ur.db.Database("taski").Collection("projects")
	projectFilter := options.Find()
	projectFilter.SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	projectFindResult, err := projectCollection.Find(context.Background(), bson.D{}, projectFilter)
	avoidPanic(err)

	var projects []*entity.Project
	for projectFindResult.Next(context.Background()) {
		var projcet *entity.Project
		err := projectFindResult.Decode(&projcet)
		avoidPanic(err)
		if projcet.UserEmail == email {
			projects = append(projects, projcet)
		}
	}

	// get todos
	todoCollection := ur.db.Database("taski").Collection("todos")
	todoFindResult, err := todoCollection.Find(context.Background(), bson.D{})
	avoidPanic(err)

	var todos []*entity.Todo
	for todoFindResult.Next(context.Background()) {
		var todo *entity.Todo
		err := todoFindResult.Decode(&todo)
		avoidPanic(err)
		todos = append(todos, todo)
	}

	for _, project := range projects {
		for _, todo := range todos {
			if project.Id == todo.ProjectId {
				project.Todos = append(project.Todos, *todo)
			}
		}
	}

	return projects, nil
}
