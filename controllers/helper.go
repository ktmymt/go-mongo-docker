package controllers

import (
	"github.com/gin-gonic/gin"
)

// Response object as HTTP response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTPRes normalize HTTP Response format
func HTTPRes(c *gin.Context, httpCode int, message string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	})
	return
}
