package controllers

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetOwnProjects(*gin.Context)
	PostUser(*gin.Context)
	UpdateProjectMembers(*gin.Context)
}

type userController struct {
	us services.UserService
}

//
func NewUserController(us services.UserService) UserController {
	return &userController{
		us: us,
	}
}

func (ctl *userController) GetOwnProjects(ctx *gin.Context) {
	userId := ctx.Param("id")

	projectUser, err := ctl.us.GetOwnProjects(userId)
	AvoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "get own projects", projectUser)
}

func (ctl *userController) PostUser(ctx *gin.Context) {
	user := entity.User{}
	errors := entity.Errors{}
	errorMessage := entity.ErrorMessage{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := user.Validation(errors, errorMessage); len(err.Errors) > 0 {
		HTTPRes(ctx, http.StatusBadRequest, "Validation Error", err.Errors)
		return
	}

	newUser, err := ctl.us.CreateNewUser(&user)
	if err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(ctx, http.StatusOK, "create new user", newUser)
}

func (ctl *userController) UpdateProjectMembers(ctx *gin.Context) {
	projectId := ctx.Query("projectId")
	userId := ctx.Query("userId")

	newUser, err := ctl.us.UpdateProjectMembers(projectId, userId)
	AvoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "assign new user", newUser)
}
