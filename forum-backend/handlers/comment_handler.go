package handlers

import (
	"forum/dto"
	"forum/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateCommentDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment := models.Comment{
			Content:   input.Content,
			PostID:    input.PostID,
			UserID:    input.UserID,
			ParentID:  input.ParentID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, dto.CommentResponseDTO{
			ID:        comment.ID,
			Content:   comment.Content,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			ParentID:  comment.ParentID,
			CreatedAt: comment.CreatedAt.Format(time.RFC3339),
			UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func GetComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comments []models.Comment
		if err := db.Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response []dto.CommentResponseDTO
		for _, comment := range comments {
			response = append(response, dto.CommentResponseDTO{
				ID:        comment.ID,
				Content:   comment.Content,
				PostID:    comment.PostID,
				UserID:    comment.UserID,
				ParentID:  comment.ParentID,
				CreatedAt: comment.CreatedAt.Format(time.RFC3339),
				UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, response)
	}
}

func GetCommentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var comment models.Comment
		if err := db.First(&comment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}

		c.JSON(http.StatusOK, dto.CommentResponseDTO{
			ID:        comment.ID,
			Content:   comment.Content,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			ParentID:  comment.ParentID,
			CreatedAt: comment.CreatedAt.Format(time.RFC3339),
			UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func UpdateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var comment models.Comment
		if err := db.First(&comment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}

		var input dto.UpdateCommentDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.Content != nil {
			comment.Content = *input.Content
		}
		comment.UpdatedAt = time.Now()

		if err := db.Save(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.CommentResponseDTO{
			ID:        comment.ID,
			Content:   comment.Content,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			ParentID:  comment.ParentID,
			CreatedAt: comment.CreatedAt.Format(time.RFC3339),
			UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
		})
	}
}

func DeleteComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Comment{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
	}
}