package migrations

import (
	"basicmigrate/models"

	"gorm.io/gorm"
)

type AddAgeToUsers struct{}

func (m AddAgeToUsers) Up(db *gorm.DB) error {
	return db.Migrator().AddColumn(models.User{}, "Age")
}

func (m AddAgeToUsers) Down(db *gorm.DB) error {
	return db.Migrator().DropColumn(models.User{}, "Age")
}

func (m AddAgeToUsers) MigrationID() string {
	return "0005_add_age_to_users"
}
