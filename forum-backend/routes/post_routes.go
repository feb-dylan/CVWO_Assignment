package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
)

func PostRoutes(r *gin.RouterGroup, postHandler *handlers.PostHandler) {
	r.GET("/topics/:id/posts", postHandler.GetPostsByTopic)	
	posts := r.Group("/posts")
	{
		posts.POST("/", postHandler.CreatePost)
		posts.GET("/", postHandler.GetPosts)
		posts.GET("/:id", postHandler.GetPostByID)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
	}
}