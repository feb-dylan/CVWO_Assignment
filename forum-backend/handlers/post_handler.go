package handlers

import (
	"forum/dto"
	"forum/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreatePostDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post := models.Post{
			Title:     input.Title,
			Content:   input.Content,
			TopicID:   input.TopicID,
			UserID:    input.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, dto.PostResponseDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			TopicID:   post.TopicID,
			UserID:    post.UserID,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func GetPosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []models.Post
		if err := db.Find(&posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response []dto.PostResponseDTO
		for _, post := range posts {
			response = append(response, dto.PostResponseDTO{
				ID:        post.ID,
				Title:     post.Title,
				Content:   post.Content,
				TopicID:   post.TopicID,
				UserID:    post.UserID,
				CreatedAt: post.CreatedAt.Format(time.RFC3339),
				UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetPostByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post models.Post
		if err := db.First(&post, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		c.JSON(http.StatusOK, dto.PostResponseDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			TopicID:   post.TopicID,
			UserID:    post.UserID,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func UpdatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post models.Post
		if err := db.First(&post, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		var input dto.UpdatePostDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.Title != nil {
			post.Title = *input.Title
		}
		if input.Content != nil {
			post.Content = *input.Content
		}
		post.UpdatedAt = time.Now()

		if err := db.Save(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.PostResponseDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			TopicID:   post.TopicID,
			UserID:    post.UserID,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
			UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func DeletePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Post{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	}
}