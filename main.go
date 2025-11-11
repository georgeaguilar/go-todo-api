package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-todo-api/models"
	"go-todo-api/routes"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "password"
	dbName := "todo_db"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ No se pudo conectar a la base de datos")
	}

	db.AutoMigrate(&models.Todo{})

	r := gin.Default()

	routes.RegisterTodoRoutes(r, db)

	r.Run(":8080")
}
