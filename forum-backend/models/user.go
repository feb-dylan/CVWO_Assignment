package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"size:100;unique;not null;index"`
	Password string `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}