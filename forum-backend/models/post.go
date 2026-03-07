package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text;not null"`
	TopicID   uint      `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint      `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Topic	 Topic     `gorm:"foreignKey:TopicID"`
	User	 User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}