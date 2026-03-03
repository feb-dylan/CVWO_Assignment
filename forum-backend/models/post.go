package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null"`
	Content   string    `gorm:"type:text;not null"`
	TopicID   uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Topic	 Topic     `gorm:"foreignKey:TopicID"`
	User	 User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}