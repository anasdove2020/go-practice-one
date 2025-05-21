package main

import (
	"log"
	"net/http"
	"strconv"

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
		userGroup.GET("/:id", getUserByID)
		userGroup.POST("", createUser)
	}

	router.Run(":8080")
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = users[len(users)-1].ID + 1
	users = append(users, newUser)
}
