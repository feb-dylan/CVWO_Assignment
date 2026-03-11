package handlers

import (
	"net/http"
	"strconv"
	"forum/dto"
	"forum/services"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	var req dto.CreateCommentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.PostID = uint(postID)

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	comment, err := h.commentService.CreateComment(req, userID.(uint))
	if err != nil {
		switch err.Error() {
		case "post not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "parent comment not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) GetCommentsByPost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	comments, err := h.commentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	comment, err := h.commentService.GetCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	var req dto.UpdateCommentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	comment, err := h.commentService.UpdateComment(uint(id), req, userID.(uint))
	if err != nil {
		switch err.Error() {
		case "comment not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to update this comment"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err = h.commentService.DeleteComment(uint(id), userID.(uint))
	if err != nil {
		switch err.Error() {
		case "comment not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to delete this comment"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}

func (h *CommentHandler) CreateReply(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	var req dto.CreateCommentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parentID := uint(commentID)
	req.ParentID = &parentID

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}


	parentComment, err := h.commentService.GetCommentByID(uint(commentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "parent comment not found"})
		return
	}
	req.PostID = parentComment.PostID

	comment, err := h.commentService.CreateComment(req, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) GetReplies(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	replies, err := h.commentService.GetRepliesByCommentID(uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, replies)
}