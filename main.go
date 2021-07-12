package main

import (
	"context"
	"go-mongo-docker/configs"
	"go-mongo-docker/controllers"
	"go-mongo-docker/repository"
	"go-mongo-docker/services"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("FRONT_URL")}
	r.Use(cors.New(corsConfig))

	// Setup DB
	config := configs.GetConfig()
	options := options.Client().ApplyURI(config.MongoDB.URI)
	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		panic(err)
	}

	// Setup USER controller
	userRepo := repository.NewUserRepository(client)
	userServ := services.NewUserService(userRepo)
	userCont := controllers.NewUserController(userServ)

	// Setup PROJECT controller
	projectRepo := repository.NewProjectRepository(client)
	projectserv := services.ProjectService(projectRepo)
	projcetCont := controllers.NewProjectController(projectserv)

	// Setup TODO controller
	todoRepo := repository.NewTodoRepository(client)
	todoService := services.TodoService(todoRepo)
	todoCtl := controllers.NewTodoController(todoService)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	// Setup routers for "User"
	r.GET("/api/userProjects/:id", userCont.GetOwnProjects)
	r.POST("/api/user", userCont.PostUser)
	// /api/updMembers?projectId=&userId=
	r.PUT("/api/updMembers", userCont.UpdateProjectMembers)

	// Setup routers for "Project"
	r.GET("/api/projects", projcetCont.GetProjects)
	r.POST("/api/project", projcetCont.PostProject)
	r.PUT("/api/updProject/:id", projcetCont.UpdateProject)
	r.DELETE("/api/delProject/:id", projcetCont.DeleteProject)

	// Setup routers for "TODO"
	r.GET("/api/todos/:id", todoCtl.GetTodos)
	r.POST("/api/todo", todoCtl.PostTodo)
	r.PUT("/api/updTodo/:id", todoCtl.UpdateTodo)
	r.DELETE("/api/delTodo/:id", todoCtl.DeleteTodo)

	r.Run()
}
