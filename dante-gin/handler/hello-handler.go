package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dante-gin/model"
)

func GetHello(c *gin.Context) {
	name := c.Param("id")
	if name == "x" {
		name = "无名"
	}
	//c.String(http.StatusOK, "Hello Gin GET, id = %s\n", c.Param("id"))
	c.JSON(http.StatusOK, model.User{Name: name, Address: "太平桥东里17号楼", Method: "GET"})
}

func PostHello(c *gin.Context) {
	name := c.Param("id")
	if name == "x" {
		name = "无名"
	}

	//c.String(http.StatusOK, "Hello Gin POST, id = %s\n", c.Param("id"))

	//c.JSON(http.StatusOK, model.User{Name: name, Address: "太平桥东里17号楼", Method: "POST"})

	users := []model.User{buildUser(name), buildUser(name + "_" + name)}
	c.JSON(http.StatusOK, gin.H {
		"users": users,
		"title": "用户列表",
	})
}

func buildUser(name string) model.User {
	return model.User{Name: name, Address: "太平桥东里17号楼", Method: "POST"}
}
