package controllers

import (
	"go-mongo-docker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetOwnProjects(*gin.Context)
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
	email := ctx.DefaultQuery("email", "")

	ownProjects, err := ctl.us.GetOwnProjects(email)
	AvoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "get own projects", ownProjects)
}

func (ctl *userController) UpdateProjectMembers(ctx *gin.Context) {
	// email := ctx.DefaultQuery("email", "")

	// newUser, err := ctl.us.UpdateProjectMembers(email)
	// AvoidPanic(err)

}
