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
	collection := p.db.Database("projects-db").Collection("projects")

	filter := options.Find()
	filter.SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	cur, err := collection.Find(context.Background(), bson.D{}, filter)
	avoidPanic(err)

	var results []*entity.Project

	for cur.Next(context.Background()) {
		var projcet *entity.Project
		err := cur.Decode(&projcet)
		avoidPanic(err)
		results = append(results, projcet)
	}

	return results, nil
}

// CreateProject() registers a project in db.
func (p *projectRepository) CreateProject(project *entity.Project) (*entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("projects-db").Collection("projects")

	insert := bson.D{
		{Key: "id", Value: primitive.NewObjectID()},
		{Key: "name", Value: project.Name},
		{Key: "description", Value: project.Description},
		{Key: "todos", Value: project.Todos},
		{Key: "color", Value: project.Color},
		{Key: "updatedAt", Value: time.Now()},
	}
	_, err := collection.InsertOne(ctx, insert)
	avoidPanic(err)

	return project, nil
}

// UpdateProject() updates data of a project.
func (p *projectRepository) UpdateProject(project *entity.Project, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := p.db.Database("projects-db").Collection("projects")

	filter := bson.M{"id": convertToInt(id)}
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

	collection := p.db.Database("projects-db").Collection("projects")

	filter := bson.M{"id": convertToInt(id)}
	result, err := collection.DeleteOne(ctx, filter)
	avoidPanic(err)

	return result, nil
}
