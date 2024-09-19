package models

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Body      string
	UserID    uint
	PostId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
