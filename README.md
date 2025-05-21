# INITIAL SETUP

Create go.mod

```
go mod init golangpracticeone
```

## Mengunduh dan menginstal package gin (web framework Go) ke dalam project-mu.

```
🌐 Fungsi Umum gin di Aplikasi Go
Gin membantu kamu membangun aplikasi web / REST API dengan lebih mudah, cepat, dan terstruktur. Secara fungsional, gin menyediakan:

🔧 1. Routing
Mengatur permintaan HTTP ke handler yang tepat.

Contoh:

r.GET("/users", getUsers)   // Handle GET /users
r.POST("/users", createUser) // Handle POST /users
📦 2. Parameter dan Query Parsing
Membaca parameter dari URL, query string, dan body JSON.

Contoh:

id := c.Param("id")         // /users/:id
name := c.Query("name")     // /users?name=John
📤 3. JSON Response Helper
Mempermudah pengiriman data sebagai respons JSON.

c.JSON(200, gin.H{"message": "success"})
🧱 4. Middleware Support
Bisa pakai middleware seperti logging, autentikasi, CORS, dll.

r.Use(gin.Logger())
r.Use(gin.Recovery())
Bisa bikin custom middleware juga.

🚀 5. Performance Cepat
Gin menggunakan httprouter di belakang layar — sangat cepat dan efisien.

🔒 6. Built-in Validasi & Binding
Bisa langsung bind JSON ke struct dan validasi data.

type Login struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}

var login Login
if err := c.ShouldBindJSON(&login); err != nil {
  c.JSON(400, gin.H{"error": err.Error()})
}
```



Update go.mod & Create go.sum

```
go get -u github.com/gin-gonic/gin
```
    


# API DOCUMENTATION

GET http://localhost:8080/users

GET http://localhost:8080/users/1

POST http://localhost:8080/users

```
{
    "name": "Anas",
    "email": "anas@gmail.com"
}
```



