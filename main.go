package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Anas", Email: "anas@gmail.com"},
	{ID: 2, Name: "Dafa", Email: "dafa@gmail.com"},
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		log.Println("[LOG]", method, path)
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Logger())

	// Router dasar
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Group route /users
	userGroup := router.Group("/users")
	{
		userGroup.GET("", getAllUsers)
	}

	router.Run(":8080")
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
