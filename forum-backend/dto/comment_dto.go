package dto
type CreateCommentDTO struct {
	Content  string `json:"content" binding:"required"`
	ParentID *uint  `json:"parent_id,omitempty"`
	PostID uint `json:"post_id"`
}
type UpdateCommentDTO struct {
	Content *string `json:"content" binding:"omitempty"`
}
type CommentResponseDTO struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	PostID    uint   `json:"post_id"`
	UserID    uint   `json:"user_id"`
	Username string `json:"user_name"`
	ParentID  *uint  `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Replies   []CommentResponseDTO `json:"replies,omitempty"`
}