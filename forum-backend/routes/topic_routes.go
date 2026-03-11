package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
)

func TopicRoutes(r *gin.RouterGroup, topicHandler *handlers.TopicHandler) {
	topicGroup := r.Group("/topics")
	{
		topicGroup.POST("/", topicHandler.CreateTopic)
		// topicGroup.GET("/", topicHandler.GetTopics)
		topicGroup.GET("/:id", topicHandler.GetTopicByID)
		topicGroup.PUT("/:id", topicHandler.UpdateTopic)
		topicGroup.DELETE("/:id", topicHandler.DeleteTopic)
	}
}