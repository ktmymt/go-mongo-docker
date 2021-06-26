package controllers

import (
	"go-mongo-docker/entity"
	"go-mongo-docker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TodoController interface
type TodoController interface {
	GetTodos(*gin.Context)
	PostTodo(*gin.Context)
	UpdateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}

type todoController struct {
	ts services.TodoService
}

// NewTodoController returns Todo Controller
func NewTodoController(ts services.TodoService) TodoController {
	return &todoController{ts: ts}
}

func (ctl *todoController) GetTodos(c *gin.Context) {
	todos, err := ctl.ts.GetTodos(c.Param("id"))
	AvoidPanic(err)

	HTTPRes(c, http.StatusOK, "get todos", todos)
}

func (ctl *todoController) PostTodo(c *gin.Context) {
	newTodo := entity.Todo{}
	errors := entity.Errors{}
	errorMessage := entity.ErrorMessage{}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := newTodo.Validation(errors, errorMessage); len(err.Errors) > 0 {
		HTTPRes(c, http.StatusBadRequest, "Validation Error", err.Errors)
		return
	}

	newCreatedTodo, err := ctl.ts.CreateTodo(&newTodo)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Todo saved", newCreatedTodo)
}

func (ctl *todoController) UpdateTodo(c *gin.Context) {
	updTodo := entity.Todo{}
	errors := entity.Errors{}
	errorMessage := entity.ErrorMessage{}

	if err := c.ShouldBindJSON(&updTodo); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := updTodo.Validation(errors, errorMessage); len(err.Errors) > 0 {
		HTTPRes(c, http.StatusBadRequest, "Validation Error", err.Errors)
		return
	}

	if err := c.ShouldBind(&updTodo); err == nil {
		if !areParamsValid(updTodo) {
			HTTPRes(c, http.StatusBadRequest, "Param invalid", nil)
			return
		}
	}

	result, err := ctl.ts.UpdateTodo(&updTodo, c.Param("id"))

	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	} else if result.ModifiedCount == 0 {
		HTTPRes(c, http.StatusBadRequest, "Update error: Zero Todo modified", nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Todo saved", &updTodo)
}

func (ctl *todoController) DeleteTodo(c *gin.Context) {
	delTodo := entity.Todo{}

	if err := c.ShouldBindJSON(&delTodo); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if _, err := ctl.ts.DeleteTodo(&delTodo, c.Param("id")); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Todo saved", &delTodo)
}

func areParamsValid(params entity.Todo) bool {
	paramsValid := true

	if params.Title == "" || params.Status == "" || params.Schedule < 0 {
		paramsValid = false
	}

	return paramsValid
}
