package routes

import (
	"go-todo-api/controllers"
	"go-todo-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTodoRoutes(r *gin.Engine, db *gorm.DB) {
	todoController := controllers.TodoController{DB: db}
	auth := r.Group("/todos")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/", todoController.GetTodos)
	auth.POST("/", todoController.CreateTodo)
	auth.PUT("/:id", todoController.UpdateTodo)
	auth.DELETE("/:id", todoController.DeleteTodo)
}