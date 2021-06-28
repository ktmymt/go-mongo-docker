package repository

import (
	"go-mongo-docker/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetOwnProjects(string) ([]*entity.Project, error)
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
func (ur *userRepository) GetOwnProjects(userId string) ([]*entity.Project, error) {
	return nil, nil
}
