package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-todo-api/models"
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

	r.GET("/todos", func(c *gin.Context) {
		var todos []models.Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if c.Request.Body != nil && !todo.Completed {
			todo.Completed = false
		}

		db.Create(&todo)
		c.JSON(http.StatusCreated, todo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo models.Todo
		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		var input models.Todo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo.Title = input.Title
		todo.Description = input.Description
		todo.Completed = input.Completed
		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Todo{}, id)
		c.Status(http.StatusNoContent)
	})

	r.Run(":8080")
}
