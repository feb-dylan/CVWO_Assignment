package routes

import (
	"forum/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}
}