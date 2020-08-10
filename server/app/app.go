package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	r = gin.Default()
)

// Run app
func Run() {
	// config := configs.GetConfig()

	r.GET("/", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"message": "Hello world from go server",
		})
	})

	r.Run()
}
