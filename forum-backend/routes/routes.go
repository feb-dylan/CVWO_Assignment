package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// UserRoutes(r, db)
	TopicRoutes(r, db)
	PostRoutes(r, db)
	CommentRoutes(r, db)
}