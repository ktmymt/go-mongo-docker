package services

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/repository"
)

type UserService interface {
	GetOwnProjects() ([]*entity.Project, error)
}

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (us *userService) GetOwnProjects() ([]*entity.Project, error) {
	return us.UserRepo.GetOwnProjects()
}
