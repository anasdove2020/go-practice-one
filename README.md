# INITIAL SETUP

Create go.mod

```
go mod init golangpracticeone
```

## Mengunduh dan menginstal package gin (web framework Go) ke dalam project-mu.

```
go get -u github.com/gin-gonic/gin
```

### Purpose

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

## mengunduh dan memperbarui package GORM


```
go get -u gorm.io/gorm
```

### Purpose

```
Apa itu gorm.io/gorm?
gorm.io/gorm adalah ORM (Object Relational Mapping) paling populer untuk bahasa Go (Golang). Package ini mempermudah kamu untuk berinteraksi dengan database tanpa perlu menulis banyak SQL secara manual.

✅ Fitur Utama GORM:
Mapping struct ke tabel database (ORM)

CRUD otomatis (Create, Read, Update, Delete)

Relasi antar tabel (One-to-One, One-to-Many, Many-to-Many)

Migrations (mengelola skema database)

Hooks (BeforeCreate, AfterUpdate, dll.)

Transactions

Preloading & eager loading

🧑‍💻 Contoh Penggunaan Dasar
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID    uint
    Name  string
    Email string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto-migrate
    db.AutoMigrate(&User{})

    // Create
    db.Create(&User{Name: "Alice", Email: "alice@example.com"})

    // Read
    var user User
    db.First(&user, 1)
    db.First(&user, "email = ?", "alice@example.com")

    // Update
    db.Model(&user).Update("Name", "Bob")

    // Delete
    db.Delete(&user)
}
Kalau kamu ingin koneksi ke database seperti MySQL atau PostgreSQL, kamu juga perlu install driver-nya, misalnya:

go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/postgres
```

## mengunduh dan memperbarui driver SQLite untuk GORM.

```
go get -u gorm.io/driver/sqlite
```

### Purpose

```
Apa itu gorm.io/driver/sqlite?
gorm.io/driver/sqlite adalah driver database SQLite yang digunakan oleh GORM. GORM sendiri tidak bisa konek ke database tanpa driver, dan driver ini memungkinkan GORM berkomunikasi dengan file database .sqlite atau .db.

🔧 Cara Pakainya
Contoh penggunaan GORM dengan SQLite:

package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Product struct {
    ID    uint
    Code  string
    Price uint
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto migrate
    db.AutoMigrate(&Product{})

    // Insert
    db.Create(&Product{Code: "P001", Price: 1000})

    // Query
    var product Product
    db.First(&product, 1)
}
🔍 Kapan Harus Pakai Driver SQLite?
Gunakan gorm.io/driver/sqlite kalau kamu:

Ingin membuat aplikasi ringan tanpa server database (SQLite = file saja).

Cocok untuk development, testing, CLI tools, atau aplikasi desktop ringan.

Tidak perlu fitur database skala besar seperti concurrency tinggi, sharding, dsb.

Kalau kamu pakai database lain, pakai driver yang sesuai:

MySQL: gorm.io/driver/mysql

PostgreSQL: gorm.io/driver/postgres

SQL Server: gorm.io/driver/sqlserver
```

## mengunduh dan memperbarui library validator versi 10 dari Go Playground.

```
go get -u github.com/go-playground/validator/v10
```

### Purpose

```
 Apa itu github.com/go-playground/validator/v10?
validator/v10 adalah library validasi struct dan field paling populer di Go, biasa digunakan untuk memvalidasi input dari user (misalnya dalam API request).

✅ Fitur Utama:
Validasi berdasarkan tag struct (validate:"required,email").

Validasi nested struct.

Custom validator (bisa bikin aturan sendiri).

Support untuk banyak tipe data (string, int, slice, dll).

Cepat dan ringan.

🧑‍💻 Contoh Penggunaan Dasar

package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"gte=18,lte=65"`
}

func main() {
    validate := validator.New()

    user := User{
        Name:  "",
        Email: "invalid-email",
        Age:   17,
    }

    err := validate.Struct(user)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            fmt.Printf("Field '%s' failed validation: %s\n", err.Field(), err.Tag())
        }
    } else {
        fmt.Println("Validation passed")
    }
}
📌 Contoh Tag Validasi:
Tag	Fungsi
required	Tidak boleh kosong
email	Harus format email
gte=10	Greater than or equal to 10
lte=100	Less than or equal to 100
len=5	Harus panjang 5
min=3,max=10	Minimal 3 karakter, maksimal 10
uuid	Format UUID valid

🧩 Integrasi dengan Framework
Validator ini juga sering dipakai dengan framework seperti:

Gin

Echo

Fiber

Contoh pada Gin:

type LoginRequest struct {
    Email string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
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



