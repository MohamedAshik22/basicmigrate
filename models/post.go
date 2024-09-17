package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Body      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
