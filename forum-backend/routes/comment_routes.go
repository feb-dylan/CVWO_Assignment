package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
	"gorm.io/gorm"
)

func CommentRoutes(r *gin.Engine , db *gorm.DB){
	commentGroup := r.Group("/comments")
	{
		commentGroup.POST("/", handlers.CreateComment(db))
		commentGroup.GET("/", handlers.GetComments(db))
		commentGroup.GET("/:id", handlers.GetCommentByID(db))
		commentGroup.PUT("/:id", handlers.UpdateComment(db))
		commentGroup.DELETE("/:id", handlers.DeleteComment(db))
	}
}