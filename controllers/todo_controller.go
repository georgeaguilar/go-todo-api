package controllers

import (
	"net/http"

	"go-todo-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func (t *TodoController) GetTodos(c *gin.Context) {
	var todos []models.Todo
	t.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func (t *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valor por defecto
	if !todo.Completed {
		todo.Completed = false
	}

	t.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func (t *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := t.DB.First(&todo, id).Error; err != nil {
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
	t.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	t.DB.Delete(&models.Todo{}, id)
	c.Status(http.StatusNoContent)
}
