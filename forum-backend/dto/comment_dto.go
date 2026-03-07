package dto
import "time"
type CreateCommentDTO struct {
	Content  string `json:"content" binding:"required"`
	PostID   uint   `json:"post_id" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
	ParentID *uint  `json:"parent_id,omitempty"`
}
type UpdateCommentDTO struct {
	Content *string `json:"content" binding:"omitempty"`
}
type CommentResponseDTO struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	PostID    uint   `json:"post_id"`
	UserID    uint   `json:"user_id"`
	ParentID  *uint  `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}