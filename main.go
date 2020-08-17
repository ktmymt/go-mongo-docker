package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-mongo-docker/configs"
	"go-mongo-docker/controllers"
	"go-mongo-docker/domain/repository"
	"go-mongo-docker/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	config := configs.GetConfig()

	options := options.Client().ApplyURI(config.MongoDB.URI)

	mongodb, err := mongo.Connect(context.Background(), options)

	if err != nil {
		panic(err)
	}

	todoRepo := repository.NewTodoRepository(mongodb)

	todoService := services.TodoService(todoRepo)

	todoCtl := controllers.NewTodoController(todoService)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	r.POST("/todo", todoCtl.PostTodo)

	r.Run()
}
