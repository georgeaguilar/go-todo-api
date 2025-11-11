package routes

import (
	"go-todo-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTodoRoutes(r *gin.Engine, db *gorm.DB) {
	todoController := controllers.TodoController{DB: db}

	r.GET("/todos", todoController.GetTodos)
	r.POST("/todos", todoController.CreateTodo)
	r.PUT("/todos/:id", todoController.UpdateTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)
}
