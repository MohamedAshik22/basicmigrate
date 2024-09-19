package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type CreateUserTable struct{}

func (m CreateUserTable) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}

func (m CreateUserTable) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.User{})
}

func (m CreateUserTable) MigrationID() string {
	return "0001_create_users"
}
