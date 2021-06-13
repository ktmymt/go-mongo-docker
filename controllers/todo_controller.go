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
}

type todoController struct {
	ts services.TodoService
}

// NewTodoController returns Todo Controller
func NewTodoController(ts services.TodoService) TodoController {
	return &todoController{ts: ts}
}

func (ctl *todoController) GetTodos(c *gin.Context) {
	todos, err := ctl.ts.GetTodos()
	AvoidPanic(err)

	HTTPRes(c, http.StatusOK, "Get all todos", todos)
}

func (ctl *todoController) PostTodo(c *gin.Context) {
	newTodo := entity.Todo{}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if _, err := ctl.ts.CreateTodo(&newTodo); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Todo saved", &newTodo)
}

func (ctl *todoController) UpdateTodo(c *gin.Context) {
	updTodo := entity.Todo{}

	if err := c.ShouldBindJSON(&updTodo); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
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

func areParamsValid(params entity.Todo) bool {
	paramsValid := true

	if params.Id <= 0 || params.Title == "" || params.Status == "" || params.Schedule < 0 {
		paramsValid = false
	}

	return paramsValid
}
