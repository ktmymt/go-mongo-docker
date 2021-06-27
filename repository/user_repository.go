package repository

import (
	"go-mongo-docker/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetOwnProject(*entity.User) ([]*entity.Project, error)
}

type userRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

/**
 * @summary: stores a user info on database
 * @return : projects retrieved by the user id
 */
func (ur *userRepository) GetOwnProject(user *entity.User) ([]*entity.Project, error) {
	return nil, nil
}
