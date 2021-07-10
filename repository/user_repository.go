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
	GetOwnProjects(*entity.User) ([]*entity.Project, error)
	CreateNewUser(*entity.User) (*entity.User, error)
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

/**
 * @summary: gets projects by user id
 * @return : projects, error
 */
func (ur *userRepository) GetOwnProjects(user *entity.User) ([]*entity.Project, error) {

	// get projects
	projectCollection := ur.db.Database("taski").Collection("projects")
	projectFilter := options.Find()
	projectFilter.SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	projectFindResult, err := projectCollection.Find(context.Background(), bson.D{}, projectFilter)
	avoidPanic(err)

	var projects []*entity.Project
	for projectFindResult.Next(context.Background()) {
		var project *entity.Project
		err := projectFindResult.Decode(&project)
		avoidPanic(err)
		// projects = append(projects, project)

		for _, eachUser := range project.User {
			if eachUser.Email == user.Email {
				projects = append(projects, project)
			}
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

func (ur *userRepository) CreateNewUser(user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := ur.db.Database("taski").Collection("users")

	insert := bson.D{
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		// {Key: "image", Value: user.Image},
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

	var newUser *entity.User
	decodedUpdateResult := updateResult.Decode(&newUser)
	avoidPanic(decodedUpdateResult)

	return newUser, nil
}
