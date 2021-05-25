package controllers

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
type ProjectController interface {
	GetProjects(*gin.Context)
	PostProject(*gin.Context)
}

//
type projectController struct {
	ps services.ProjectService
}

//
func NewProjectController(ps services.ProjectService) ProjectController {
	return &projectController{
		ps: ps,
	}
}

func (ctl *projectController) GetProjects(ctx *gin.Context) {
	projects, err := ctl.ps.GetProjects()
	avoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "get project test -> ok", projects)
}

func (ctl *projectController) PostProject(ctx *gin.Context) {
	newProject := entity.Project{}

	if err := ctx.ShouldBindJSON(&newProject); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if _, err := ctl.ps.CreateProject(&newProject); err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(ctx, http.StatusOK, "Project saved", &newProject)
}

// avoidPanic() catches an error and terminates the program.
func avoidPanic(err error) {
	if err != nil {
		panic(err)
	}
}
