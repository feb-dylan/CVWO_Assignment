package routes

import (
	"github.com/gin-gonic/gin"
	"forum/handlers"
	"gorm.io/gorm"
)

func TopicRoutes(r *gin.Engine, db *gorm.DB) {
	topicGroup := r.Group("/topics")
	{
		topicGroup.POST("/", handlers.CreateTopic(db))
		topicGroup.GET("/", handlers.GetTopics(db))
		topicGroup.GET("/:id", handlers.GetTopicByID(db))
		topicGroup.PUT("/:id", handlers.UpdateTopic(db))
		topicGroup.DELETE("/:id", handlers.DeleteTopic(db))
	}
}