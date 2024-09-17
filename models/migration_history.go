package models

import (
	"time"
)

type MigrationHistory struct {
	ID          uint      `gorm:"primaryKey"`
	MigrationID string    `gorm:"unique;not null"`
	AppliedAt   time.Time `gorm:"not null"`
}
