package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	// r.GET("/todo", func(cxt *gin.Context) {
	// 	cxt.JSON(http.StatusOK, gin.H{
	// 		"todo": {
	// 			"title": "this is a title",
	// 			"description": "this is a description"
	// 		},
	// 	})
	// })

	r.Run()
}
