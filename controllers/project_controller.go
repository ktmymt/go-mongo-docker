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
	UpdateProject(*gin.Context)
	DeleteProject(*gin.Context)
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
	AvoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "get project test -> ok", projects)
}

func (ctl *projectController) PostProject(ctx *gin.Context) {
	newProject := entity.Project{}

	if err := ctx.ShouldBindJSON(&newProject); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctl.ps.ValidateLength(&newProject); len(err) > 0 {
		HTTPRes(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if _, err := ctl.ps.CreateProject(&newProject); err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(ctx, http.StatusOK, "Project saved", &newProject)
}

func (ctl *projectController) UpdateProject(ctx *gin.Context) {
	updProject := entity.Project{}

	if err := ctx.ShouldBindJSON(&updProject); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctl.ps.ValidateLength(&updProject); len(err) > 0 {
		HTTPRes(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if _, err := ctl.ps.UpdateProject(&updProject, ctx.Param("id")); err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(ctx, http.StatusOK, "Project updated", &updProject)
}

func (ctl *projectController) DeleteProject(ctx *gin.Context) {
	delProject := entity.Project{}

	if err := ctx.ShouldBindJSON(&delProject); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	if _, err := ctl.ps.DeleteProject(&delProject, ctx.Param("id")); err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(ctx, http.StatusOK, "Project deleted", &delProject)

}
