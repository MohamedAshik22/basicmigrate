package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FullName  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Age       int       `gorm:"not null"`
	Dob       time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
