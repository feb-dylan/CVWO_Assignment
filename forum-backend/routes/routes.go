package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
)

func RegisterRoutes(
	r *gin.Engine,
	userHandler *handlers.UserHandler,
	topicHandler *handlers.TopicHandler,
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
	authMiddleware gin.HandlerFunc,
) {
	AuthRoutes(r, userHandler)
	r.GET("/topics", topicHandler.GetTopics)
	protected := r.Group("/")
	protected.Use(authMiddleware)
	{
		TopicRoutes(protected, topicHandler) 
		PostRoutes(protected, postHandler)
		CommentRoutes(protected, commentHandler)
	}
}