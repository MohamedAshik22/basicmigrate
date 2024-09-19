package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type CreateCommentTable struct{}

func (m CreateCommentTable) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Comment{})
}

func (m CreateCommentTable) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Comment{})
}

func (m CreateCommentTable) MigrationID() string {
	return "0003_create_comments"
}
