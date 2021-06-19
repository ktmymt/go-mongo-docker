package services

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

// Projects ervice interface
type ProjectService interface {
	GetProjects() ([]*entity.Project, error)
	CreateProject(project *entity.Project) (*entity.Project, error)
	UpdateProject(project *entity.Project, id string) (*mongo.UpdateResult, error)
	DeleteProject(project *entity.Project, id string) (*mongo.DeleteResult, error)
}

// Project service structure
type projectService struct {
	ProjectRepo repository.ProjectRepository
}

// NewProjectService return service
func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{
		ProjectRepo: projectRepo,
	}
}

func (ps *projectService) GetProjects() ([]*entity.Project, error) {
	return ps.ProjectRepo.GetProjects()
}

func (ps *projectService) CreateProject(project *entity.Project) (*entity.Project, error) {
	return ps.ProjectRepo.CreateProject(project)
}

func (ps *projectService) UpdateProject(project *entity.Project, id string) (*mongo.UpdateResult, error) {
	return ps.ProjectRepo.UpdateProject(project, id)
}

func (ps *projectService) DeleteProject(project *entity.Project, id string) (*mongo.DeleteResult, error) {
	return ps.ProjectRepo.DeleteProject(project, id)
}
