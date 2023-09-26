package main

import (
	"dante-gin/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	HOST = "0.0.0.0:8900"
)

func main() {
	router := gin.Default()

	router.GET("/hello/:id", handler.GetHello)
	router.POST("/hello/:id", handler.PostHello)

	// 多文件上传
	router.POST("/upload", handler.UploadMultiFile)

	// HTML 渲染
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "你好，Gin 框架！",
		})
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"title": "Gin 框架主页！",
		})
	})
	router.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "Gin 爱好者！",
		})
	})

	err := router.Run(HOST)
	if err != nil {
		return
	}
}
