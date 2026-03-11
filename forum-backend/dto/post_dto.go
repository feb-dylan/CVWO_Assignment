package dto

type CreatePostDTO struct {
	Title   string `json:"title" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required"`
	TopicID uint   `json:"topic_id" binding:"required"`
}

type UpdatePostDTO struct {
	Title   *string `json:"title" binding:"omitempty,min=3,max=255"`
	Content *string `json:"content" binding:"omitempty"`
}

type PostResponseDTO struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TopicID   uint   `json:"topic_id"`
	UserID    uint   `json:"user_id"`
	Username string `json:"user_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}