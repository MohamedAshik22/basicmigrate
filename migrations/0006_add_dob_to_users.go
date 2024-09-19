package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type AddDobToUsers struct{}

func (m AddDobToUsers) Up(db *gorm.DB) error {
	return db.Migrator().AddColumn(models.User{}, "Dob")
}

func (m AddDobToUsers) Down(db *gorm.DB) error {
	return db.Migrator().DropColumn(models.User{}, "Dob")
}

func (m AddDobToUsers) MigrationID() string {
	return "0006_add_dob_to_users"
}
