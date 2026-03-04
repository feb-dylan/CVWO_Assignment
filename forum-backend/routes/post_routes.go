package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
	"gorm.io/gorm"
)
func PostRoutes(r *gin.Engine, db *gorm.DB) {
	postGroup := r.Group("/posts")
	{
		postGroup.POST("/", handlers.CreatePost(db))
		postGroup.GET("/", handlers.GetPosts(db))
		postGroup.GET("/:id", handlers.GetPostByID(db))
		postGroup.PUT("/:id", handlers.UpdatePost(db))
		postGroup.DELETE("/:id", handlers.DeletePost(db))
	}
}