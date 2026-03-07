package dto
import "time"
type CreateTopicDTO struct {
	Title       string `json:"title" binding:"required,min=3,max=255"`
	Description string `json:"description" binding:"omitempty,max=1000"`
	UserID      uint   `json:"user_id" binding:"required"`
}

type UpdateTopicDTO struct {
	Title       *string `json:"title" binding:"omitempty,min=3,max=255"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
}

type TopicResponseDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}