package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primaryKey"`
	Content   string     `gorm:"type:text;not null"`
	PostID    uint       `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint       `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ParentID  *uint      `gorm:"default:null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
	Post	  Post       `gorm:"foreignKey:PostID"`
	User	  User       `gorm:"foreignKey:UserID"`
	Parent	*Comment   `gorm:"foreignKey:ParentID"`
	Replies	[]Comment  `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}