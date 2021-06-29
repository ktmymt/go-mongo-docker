package repository

import (
	"context"
	"go-mongo-docker/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetOwnProjects(string, string, string) ([]*entity.Project, error)
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
func (ur *userRepository) GetOwnProjects(userId string, username string, email string) ([]*entity.Project, error) {
	// register user
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

	return nil, nil
}
