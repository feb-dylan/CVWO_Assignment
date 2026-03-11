package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
)

func CommentRoutes(r *gin.RouterGroup ,commentHandler *handlers.CommentHandler){
	r.POST("/posts/:id/comments", commentHandler.CreateComment)
	r.GET("/posts/:id/comments", commentHandler.GetCommentsByPost)
	r.POST("/comments/:id/replies", commentHandler.CreateReply)
	r.GET("/comments/:id/replies", commentHandler.GetReplies)
	comments := r.Group("/comments")
	{
		comments.GET("/:id", commentHandler.GetCommentByID)
		comments.PUT("/:id", commentHandler.UpdateComment)
		comments.DELETE("/:id", commentHandler.DeleteComment)
	}
}