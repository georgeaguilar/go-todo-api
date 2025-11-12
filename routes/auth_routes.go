package routes

import (
	"go-todo-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.AuthController{DB: db}

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}
