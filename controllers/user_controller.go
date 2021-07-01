package controllers

import (
	"go-mongo-docker/services"
	"net/http"

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
	username := ctx.DefaultQuery("username", "")
	email := ctx.DefaultQuery("email", "")

	ownProjects, err := ctl.us.GetOwnProjects(username, email)
	AvoidPanic(err)

	HTTPRes(ctx, http.StatusOK, "get own projects", ownProjects)
}
