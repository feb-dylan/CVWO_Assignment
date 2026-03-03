package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primaryKey"`
	Content   string     `gorm:"type:text;not null"`
	PostID    uint       `gorm:"not null"`
	UserID    uint       `gorm:"not null"`
	ParentID  *uint      `gorm:"default:null"`
	Post	  Post       `gorm:"foreignKey:PostID"`
	User	  User       `gorm:"foreignKey:UserID"`
	Parent	*Comment   `gorm:"foreignKey:ParentID"`
	Replies	[]Comment  `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}