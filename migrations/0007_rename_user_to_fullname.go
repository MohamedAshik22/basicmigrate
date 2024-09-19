package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type RenameUserNameToFullName struct{}

func (m RenameUserNameToFullName) Up(db *gorm.DB) error {
	return db.Migrator().RenameColumn(&models.User{}, "Name", "fullname")
}

func (m RenameUserNameToFullName) Down(db *gorm.DB) error {
	return db.Migrator().RenameColumn(&models.User{}, "fullname", "Name")
}

func (m RenameUserNameToFullName) MigrationID() string {
	return "0007_rename_user_name_to_full_name"
}
