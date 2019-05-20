package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

func GetHello(c *gin.Context) {

	c.String(http.StatusOK, "Hello Gin GET, id = %s\n", c.Param("id"))
}

func PostHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin POST, id = %s\n", c.Param("id"))
}
