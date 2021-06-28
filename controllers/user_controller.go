package controllers

import (
	"go-mongo-docker/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetOwnProjects(*gin.Context)
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

	// ownProjects, err := ctl.us.GetOwnProjects()
	// AvoidPanic(err)

	// HTTPRes(ctx, http.StatusOK, "get own projects", ownProjects)

}
