package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var userList []Users

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/user", addUser)

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func addUser(c *gin.Context) {
	var newUser Users

	err := c.BindJSON(&newUser)
	if err != nil {
		fmt.Println("Something is wrong : ", err)
		c.JSON(http.StatusBadRequest, "Bad Data provided")
		return
	}

	userList = append(userList, newUser)
	c.JSON(http.StatusCreated, "Success")
}
