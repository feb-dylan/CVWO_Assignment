package handlers

import (
	"net/http"
	"strconv"
	"forum/dto"
	"forum/services"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	topicService *services.TopicService
}

func NewTopicHandler(topicService *services.TopicService) *TopicHandler {
	return &TopicHandler{topicService: topicService}
}

func (h *TopicHandler) CreateTopic(c *gin.Context) {
	var topicDTO dto.CreateTopicDTO
	if err := c.ShouldBindJSON(&topicDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	topic, err := h.topicService.CreateTopic(topicDTO, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

func (h *TopicHandler) GetTopics(c *gin.Context) {
	topics, err := h.topicService.GetAllTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) GetTopicByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic id"})
		return
	}
	topic, err := h.topicService.GetTopicByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
		return
	}
	c.JSON(http.StatusOK, topic)
}
func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid topic id"})
		return
	}

	var req dto.UpdateTopicDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	topic, err := h.topicService.UpdateTopic(uint(id), req, userID.(uint))
	if err != nil {
		switch err.Error() {
		case "topic not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to update this topic"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid topic id"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = h.topicService.DeleteTopic(uint(id), userID.(uint))
	if err != nil {
		switch err.Error() {
		case "topic not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to delete this topic"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "topic deleted successfully"})
}