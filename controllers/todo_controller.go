package controllers

import (
	"github.com/gin-gonic/gin"
	"go-mongo-docker/entity"
	"go-mongo-docker/services"
	"net/http"
)

// TodoOutput structure
type TodoOutput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodoInput structure
type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodoController interface
type TodoController interface {
	GetTodos(*gin.Context)
	PostTodo(*gin.Context)
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

	if err != nil {
		panic(err)
	}

	HTTPRes(c, http.StatusOK, "Get all todos", todos)
}

func (ctl *todoController) PostTodo(c *gin.Context) {
	var todoInput TodoInput
	if err := c.ShouldBindJSON(&todoInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	t := ctl.inputToTodo(todoInput)

	if _, err := ctl.ts.CreateTodo(&t); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	todoOutput := ctl.mapToTodoOutput(&t)
	HTTPRes(c, http.StatusOK, "Todo saved", todoOutput)
}

func (ctl *todoController) inputToTodo(input TodoInput) entity.Todo {
	return entity.Todo{
		Title:       input.Title,
		Description: input.Description,
	}
}

func (ctl *todoController) mapToTodoOutput(t *entity.Todo) *TodoOutput {
	return &TodoOutput{
		Title:       t.Title,
		Description: t.Description,
	}
}
