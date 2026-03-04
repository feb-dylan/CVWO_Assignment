package handlers

import (
	"forum/dto"
	"forum/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTopic(db *gorm.DB) gin.HandlerFunc	{
	return func(c *gin.Context)	{
	var input dto.CreateTopicDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topic := models.Topic{
		Title: input.Title,
		Description: input.Description,
		UserID: input.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated , dto.TopicResponseDTO{
		ID: topic.ID,
		Title: topic.Title,
		Description: topic.Description,
		UserID: topic.UserID,
		CreatedAt: topic.CreatedAt.Format(time.RFC3339),
		UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
	})
	}
}
func GetTopics(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var topics []models.Topic
		if err := db.Find(&topics).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var response []dto.TopicResponseDTO
		for _, topic := range topics {
			response = append(response, dto.TopicResponseDTO{
				ID: topic.ID,
				Title: topic.Title,
				Description: topic.Description,
				UserID: topic.UserID,
				CreatedAt: topic.CreatedAt.Format(time.RFC3339),
				UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
			})
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetTopicByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var topic models.Topic
		if err := db.First(&topic, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
			return
		}
		c.JSON(http.StatusOK, dto.TopicResponseDTO{
			ID: topic.ID,
			Title: topic.Title,
			Description: topic.Description,
			UserID: topic.UserID,
			CreatedAt: topic.CreatedAt.Format(time.RFC3339),
			UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func UpdateTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var topic models.Topic
		if err := db.First(&topic, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
			return
		}
		var input dto.UpdateTopicDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.Title != nil {
			topic.Title = *input.Title
		}
		if input.Description != nil {
			topic.Description = *input.Description
		}
		topic.UpdatedAt = time.Now()
		if err := db.Save(&topic).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.TopicResponseDTO{
			ID: topic.ID,
			Title: topic.Title,
			Description: topic.Description,
			UserID: topic.UserID,
			CreatedAt: topic.CreatedAt.Format(time.RFC3339),
			UpdatedAt: topic.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func DeleteTopic(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Topic{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Topic deleted"})
	}
}