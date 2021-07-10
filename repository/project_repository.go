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

// Project repository functions
type ProjectRepository interface {
	GetProjects() ([]*entity.Project, error)
	CreateProject(*entity.Project) (*entity.Project, error)
	UpdateProject(*entity.Project, string) (*mongo.UpdateResult, error)
	DeleteProject(*entity.Project, string) (*mongo.DeleteResult, error)
}

// Project repository structure has db
type projectRepository struct {
	db *mongo.Client
}

// NewProjectRepository returns "ProjectRepository"
func NewProjectRepository(db *mongo.Client) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

// GetProjects() returns all projects.
func (p *projectRepository) GetProjects() ([]*entity.Project, error) {
	projectCollection := p.db.Database("taski").Collection("projects")

	filter := options.Find()
	filter.SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	projectFindResult, err := projectCollection.Find(context.Background(), bson.D{}, filter)
	avoidPanic(err)

	todoCollection := p.db.Database("taski").Collection("todos")
	todoFindResult, err := todoCollection.Find(context.Background(), bson.D{})
	avoidPanic(err)

	var projects []*entity.Project
	var todos []*entity.Todo

	for projectFindResult.Next(context.Background()) {
		var projcet *entity.Project
		err := projectFindResult.Decode(&projcet)
		avoidPanic(err)
		projects = append(projects, projcet)
	}

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

// CreateProject() registers a project in db.
func (p *projectRepository) CreateProject(project *entity.Project) (*entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("taski").Collection("projects")

	insert := bson.D{
		{Key: "userIds", Value: project.UserIds},
		{Key: "name", Value: project.Name},
		{Key: "description", Value: project.Description},
		{Key: "todos", Value: project.Todos},
		{Key: "color", Value: project.Color},
		{Key: "updatedAt", Value: time.Now()},
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

	var newProject *entity.Project
	decodedUpdateResult := updateResult.Decode(&newProject)
	avoidPanic(decodedUpdateResult)

	return newProject, nil
}

// UpdateProject() updates data of a project.
func (p *projectRepository) UpdateProject(project *entity.Project, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("taski").Collection("projects")

	filter := bson.M{"_id": convertToObjectId(id)}
	update := bson.M{
		"$set": bson.M{
			"name":        project.Name,
			"description": project.Description,
			"todos":       project.Todos,
			"color":       project.Color,
			"updatedAt":   time.Now(),
		}}

	result, err := collection.UpdateOne(ctx, filter, update)
	avoidPanic(err)

	return result, nil
}

// DeleteProject() deletes data of a project.
func (p *projectRepository) DeleteProject(project *entity.Project, id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	projectCollection := p.db.Database("taski").Collection("projects")
	todoCollection := p.db.Database("taski").Collection("todos")

	projectFilter := bson.M{"_id": convertToObjectId(id)}
	result, err := projectCollection.DeleteOne(ctx, projectFilter)
	avoidPanic(err)

	todoFilter := bson.M{"projectId": convertToObjectId(id)}
	_, err = todoCollection.DeleteMany(ctx, todoFilter)
	avoidPanic(err)

	return result, nil
}
