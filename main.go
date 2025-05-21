package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Model sederhana
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Simulasi database in-memory
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// Middleware logging sederhana
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log path dan method
		method := c.Request.Method
		path := c.Request.URL.Path
		println("[LOG]", method, path)
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Logger())

	// Route dasar
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Group route /users
	userGroup := router.Group("/users")
	{
		userGroup.GET("", getAllUsers)
		userGroup.GET("/:id", getUserByID)
		userGroup.POST("", createUser)
		userGroup.PUT("/:id", updateUser)
		userGroup.DELETE("/:id", deleteUser)
	}

	router.Run(":8080")
}

// Handler get semua user
func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Handler get user berdasarkan ID
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

// Handler create user baru (body JSON)
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Simpel: assign ID bertambah 1 dari terakhir
	newUser.ID = users[len(users)-1].ID + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// Handler update user berdasarkan ID
func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			c.JSON(http.StatusOK, users[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// Handler delete user berdasarkan ID
func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}
