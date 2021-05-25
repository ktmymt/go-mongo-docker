package services

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/repository"
)

// Projects ervice interface
type ProjectService interface {
	GetProjects() ([]*entity.Project, error)
	CreateProject(project *entity.Project) (*entity.Project, error)
	// UpdateProject(project *entity.Project) (*entity.Project, error)
	// DeleteProject(project *entity.Project) (*entity.Project, error)
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

// func (ps *projectService) UpdateProject(project *entity.Project) (*entity.Project, error){
// 	return ps.ProjectRepo.UpdateProject(project)
// }

// func (ps *projectService) DeleteProject(project *entity.Project) (*entity.Project, error){
// 	retur ps.ProjectRepo.DeleteProject(project)
// }
