package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type CreatePostTable struct{}

func (m CreatePostTable) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Post{})
}

func (m CreatePostTable) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Post{})
}
