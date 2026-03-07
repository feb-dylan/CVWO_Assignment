package routes

import (
	"forum/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", handlers.RegisterUser(db))
		userGroup.POST("/login", handlers.LoginUser(db))
	}
}