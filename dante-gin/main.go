package main

import (
	"dante-gin/handler"
	"github.com/gin-gonic/gin"
)

const (
	HOST = "0.0.0.0:8900"
)

func main() {
	router := gin.Default()

	router.GET("/hello/:id", handler.GetHello)
	router.POST("/hello/:id", handler.PostHello)

	router.Run(HOST)
}
