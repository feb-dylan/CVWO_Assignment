package models

import "time"

type Topic struct {
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	UserID uint `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	User User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}